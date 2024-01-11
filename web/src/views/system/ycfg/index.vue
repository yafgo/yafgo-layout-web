<template>
  <a-row :gutter="8" class="editor-container">
    <a-col :span="12">
      当前配置
      <CfgAll ref="cfgAllRef" />
    </a-col>
    <a-col :span="12">
      编辑配置
      <CfgEdit @change="handleChange" />
    </a-col>
  </a-row>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import CfgAll from './components/CfgAll.vue';
  import CfgEdit from './components/CfgEdit.vue';

  const cfgAllRef = ref();

  const handleChange = () => {
    cfgAllRef.value?.refresh();
  };
</script>

<style lang="less" scoped>
  .editor-container {
    padding: 12px 16px;
  }
</style>
