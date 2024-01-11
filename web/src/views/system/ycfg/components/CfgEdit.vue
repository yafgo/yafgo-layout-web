<template>
  <div class="container-cfg-edit">
    <div ref="editorRef" class="editor-main"></div>
    <div class="btn-wrap">
      <a-button type="primary" @click="saveCfg">保存当前配置</a-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { StreamLanguage } from '@codemirror/language';
  import { yaml as yamlMode } from '@codemirror/legacy-modes/mode/yaml';
  import { EditorState } from '@codemirror/state';
  import { githubDark } from '@ddietr/codemirror-themes/github-dark';
  import { EditorView, basicSetup } from 'codemirror';
  import { defineEmits, onMounted, ref } from 'vue';
  import { ApiGetCfgInRedis, ApiSetCfgInRedis } from '@/api/system';
  import { Message } from '@arco-design/web-vue';

  const editorRef = ref();
  const editorView = ref();
  const yaml = StreamLanguage.define(yamlMode);

  const emit = defineEmits({
    change: () => true,
  });

  const initCodeMirror = () => {
    if (editorView.value) {
      return;
    }
    const startState = EditorState.create({
      doc: ``,
      extensions: [basicSetup, githubDark, yaml],
    });
    if (editorRef.value) {
      editorView.value = new EditorView({
        state: startState,
        parent: editorRef.value,
      });
    }
  };

  const getEditorContent = () => {
    return editorView.value.state.doc.toString();
  };

  const setEditorContent = (content: string) => {
    editorView.value.dispatch({
      changes: {
        from: 0,
        to: editorView.value.state.doc.length,
        insert: content,
      },
    });
  };

  const getCfgInRedis = async () => {
    const res = await ApiGetCfgInRedis();
    if (!res || !res.success) {
      return;
    }
    setEditorContent(res.data);
  };

  const saveCfg = async () => {
    const newContent = editorView.value.state.doc.toString();
    // console.log('编辑后', newContent);
    const res = await ApiSetCfgInRedis(newContent);
    if (!res || !res.success) {
      return;
    }
    // 成功后刷新
    emit('change');
    await getCfgInRedis();
    Message.success('保存成功');
  };

  onMounted(() => {
    initCodeMirror();
    getCfgInRedis();
  });
</script>

<style lang="less" scoped>
  .container-cfg-edit {
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
