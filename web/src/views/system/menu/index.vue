<template>
  <div class="container">
    <Breadcrumb :items="['menu.system', 'menu.system.menu']" />
    <a-card class="general-card">
      <a-row style="padding: 16px 0 10px 0; align-items: center">
        <a-col :span="12">
          <a-space>
            <a-button type="primary">
              <template #icon> <icon-plus /> </template>
              {{ $t('searchTable.operation.create') }}
            </a-button>
          </a-space>
        </a-col>
        <a-col
          :span="12"
          style="display: flex; align-items: center; justify-content: end"
        >
          <a-tooltip :content="$t('searchTable.actions.refresh')">
            <div class="action-icon" @click="search">
              <icon-refresh size="18" />
            </div>
          </a-tooltip>
          <a-dropdown @select="handleSelectDensity">
            <a-tooltip :content="$t('searchTable.actions.density')">
              <div class="action-icon"><icon-line-height size="18" /></div>
            </a-tooltip>
            <template #content>
              <a-doption
                v-for="item in densityList"
                :key="item.value"
                :value="item.value"
                :class="{ active: item.value === size }"
              >
                <span>{{ item.name }}</span>
              </a-doption>
            </template>
          </a-dropdown>
        </a-col>
      </a-row>
      <a-table
        row-key="id"
        :loading="loading"
        :columns="(cloneColumns as TableColumnData[])"
        :data="renderData"
        :bordered="{ cell: true }"
        :size="size"
        :default-expand-all-rows="true"
      >
        <template #name="{ record }">
          {{ $t(record?.meta?.locale) }}
        </template>
        <template #actions>
          <a-button type="text" size="small">
            {{ $t('searchTable.columns.operations.view') }}
          </a-button>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
  import { ApiGetList, MenuItem, MenuParams } from '@/api/menu';
  import useLoading from '@/hooks/loading';
  import type { TableColumnData } from '@arco-design/web-vue/es/table/interface';
  import cloneDeep from 'lodash/cloneDeep';
  import { computed, ref, watch } from 'vue';
  import { useI18n } from 'vue-i18n';

  type SizeProps = 'mini' | 'small' | 'medium' | 'large';
  type Column = TableColumnData & { checked?: true };

  const { loading, setLoading } = useLoading(true);
  const { t } = useI18n();
  const renderData = ref<MenuItem[]>([]);
  const formModel = ref();
  const cloneColumns = ref<Column[]>([]);
  const showColumns = ref<Column[]>([]);

  /** 表格显示密度 */
  const size = ref<SizeProps>('medium');
  const densityList = computed(() => [
    {
      name: t('searchTable.size.mini'),
      value: 'mini',
    },
    {
      name: t('searchTable.size.small'),
      value: 'small',
    },
    {
      name: t('searchTable.size.medium'),
      value: 'medium',
    },
    {
      name: t('searchTable.size.large'),
      value: 'large',
    },
  ]);
  /** 切换表格显示密度 */
  const handleSelectDensity = (
    val: string | number | Record<string, any> | undefined,
    e: Event
  ) => {
    size.value = val as SizeProps;
  };

  const columns = computed<TableColumnData[]>(() => [
    {
      title: '序号',
      dataIndex: 'id',
    },
    {
      title: '菜单名称',
      dataIndex: 'name',
      slotName: 'name',
    },
    {
      title: '图标',
      dataIndex: 'meta.icon',
      slotName: 'icon',
    },
    {
      title: '路由名称',
      dataIndex: 'name',
    },
    {
      title: '路由地址',
      dataIndex: 'path',
    },
    /* {
      title: 'meta',
      dataIndex: 'meta',
    }, */
    {
      title: '操作',
      dataIndex: 'actions',
      slotName: 'actions',
    },
  ]);

  const fetchData = async (
    params: MenuParams = { current: 1, pageSize: 20 }
  ) => {
    setLoading(true);
    try {
      const { data } = await ApiGetList();
      renderData.value = data;
    } catch (err) {
      // you can report use errorHandler or other
    } finally {
      setLoading(false);
    }
  };

  const search = () => {
    fetchData({
      ...formModel.value,
    } as unknown as MenuParams);
  };

  fetchData();

  watch(
    () => columns.value,
    (val) => {
      cloneColumns.value = cloneDeep(val);
      cloneColumns.value.forEach((item, index) => {
        item.checked = true;
      });
      showColumns.value = cloneDeep(cloneColumns.value);
    },
    { deep: true, immediate: true }
  );
</script>

<script lang="ts">
  export default {
    name: 'MenuManagement',
  };
</script>

<style scoped lang="less">
  .container {
    padding: 0 20px 20px 20px;
  }
  :deep(.arco-table-th) {
    &:last-child {
      .arco-table-th-item-title {
        margin-left: 16px;
      }
    }
  }
  .action-icon {
    margin-left: 12px;
    cursor: pointer;
  }
  .active {
    color: #0960bd;
    background-color: #e3f4fc;
  }
</style>
