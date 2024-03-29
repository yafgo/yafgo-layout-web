import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/function',
  redirect: '/function/multipleTabs',
  name: 'Function',
  component: Layout,
  meta: {
    title: t('router.function'),
    icon: 'ri:function-fill',
    alwaysShow: true
  },
  children: [
    {
      path: 'multiple-tabs',
      component: () => import('@/views/demo/Function/MultipleTabs.vue'),
      name: 'MultipleTabs',
      meta: {
        title: t('router.multipleTabs')
      }
    },
    {
      path: 'multiple-tabs-demo/:id',
      component: () => import('@/views/demo/Function/MultipleTabsDemo.vue'),
      name: 'MultipleTabsDemo',
      meta: {
        hidden: true,
        title: t('router.details'),
        canTo: true,
        activeMenu: '/function/multiple-tabs'
      }
    },
    {
      path: 'request',
      component: () => import('@/views/demo/Function/Request.vue'),
      name: 'Request',
      meta: {
        title: t('router.request')
      }
    },
    {
      path: 'test',
      component: () => import('@/views/demo/Function/Test.vue'),
      name: 'Test',
      meta: {
        title: t('router.permission'),
        permission: ['add', 'edit', 'delete']
      }
    }
  ]
}

export default RouteItem
