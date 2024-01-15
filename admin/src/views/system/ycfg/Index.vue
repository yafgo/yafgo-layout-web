<template>
  <el-row :gutter="8">
    <el-col :span="12">
      当前配置
      <CfgAll ref="cfgAllRef" />
    </el-col>
    <el-col :span="12">
      编辑配置
      <CfgEdit @change="handleChange" />
    </el-col>
  </el-row>
</template>

<script setup>
import { ElCol, ElRow } from 'element-plus'
import { ref } from 'vue'
import CfgAll from './components/CfgAll.vue'
import CfgEdit from './components/CfgEdit.vue'

const cfgAllRef = ref()

const handleChange = () => {
  cfgAllRef.value?.refresh()
}
</script>

<style lang="less" scoped></style>
