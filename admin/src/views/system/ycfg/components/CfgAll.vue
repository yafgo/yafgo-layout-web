<template>
  <div class="container-cfg-all">
    <div class="editor-main" ref="editorRef"></div>
    <div class="btn-wrap">
      <ElButton type="primary" @click="refresh(true)">刷新</ElButton>
    </div>
  </div>
</template>

<script setup>
import { ApiGetAllCfg } from '@/api/system'
import { json } from '@codemirror/lang-json'
import { EditorState } from '@codemirror/state'
import { materialDark } from '@ddietr/codemirror-themes/material-dark'
import { basicSetup, EditorView } from 'codemirror'
import { ElButton, ElMessage } from 'element-plus'
import { defineExpose, onMounted, ref } from 'vue'

const editorRef = ref()
const editorView = ref()

const initCodeMirror = () => {
  if (editorView.value) {
    return
  }
  const startState = EditorState.create({
    doc: ``,
    extensions: [
      basicSetup,
      materialDark,
      json(),
      EditorView.editable.of(false) // 将编辑器设置为只读
    ]
  })
  if (editorRef.value) {
    editorView.value = new EditorView({
      state: startState,
      parent: editorRef.value
    })
  }
}

const getEditorContent = () => {
  return editorView.value.state.doc.toString()
}

const setEditorContent = (content) => {
  editorView.value.dispatch({
    changes: {
      from: 0,
      to: editorView.value.state.doc.length,
      insert: content
    }
  })
}

const refresh = async (showMsg = false) => {
  const res = await ApiGetAllCfg()
  if (!res || !res.success) {
    return
  }
  let cfgStr = JSON.stringify(res.data || {}, null, 2)
  setEditorContent(cfgStr)
  showMsg && ElMessage.success('刷新成功')
}

onMounted(() => {
  initCodeMirror()
  refresh()
})

defineExpose({
  refresh
})
</script>

<style lang="less" scoped>
.container-cfg-all {
  position: relative;

  .editor-main {
    width: 100%;
    height: calc(100vh - 200px);

    // [重要]撑满高度
    :deep(.cm-editor) {
      height: 100%;
    }
  }

  .btn-wrap {
    position: absolute;
    right: 16px;
    bottom: 16px;
  }
}
</style>
