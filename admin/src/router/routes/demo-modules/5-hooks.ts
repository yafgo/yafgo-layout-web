import { Layout } from '@/utils/routerHelper'

const RouteItem: AppRouteRecordRaw = {
  path: '/hooks',
  redirect: '/hooks/useWatermark',
  name: 'Hooks',
  component: Layout,
  meta: {
    title: 'hooks',
    icon: 'ic:outline-webhook',
    alwaysShow: true
  },
  children: [
    {
      path: 'useWatermark',
      component: () => import('@/views/demo/hooks/useWatermark.vue'),
      name: 'UseWatermark',
      meta: {
        title: 'useWatermark'
      }
    },
    {
      path: 'useTagsView',
      component: () => import('@/views/demo/hooks/useTagsView.vue'),
      name: 'UseTagsView',
      meta: {
        title: 'useTagsView'
      }
    },
    {
      path: 'useValidator',
      component: () => import('@/views/demo/hooks/useValidator.vue'),
      name: 'UseValidator',
      meta: {
        title: 'useValidator'
      }
    },
    {
      path: 'useCrudSchemas',
      component: () => import('@/views/demo/hooks/useCrudSchemas.vue'),
      name: 'UseCrudSchemas',
      meta: {
        title: 'useCrudSchemas'
      }
    },
    {
      path: 'useClipboard',
      component: () => import('@/views/demo/hooks/useClipboard.vue'),
      name: 'UseClipboard',
      meta: {
        title: 'useClipboard'
      }
    },
    {
      path: 'useNetwork',
      component: () => import('@/views/demo/hooks/useNetwork.vue'),
      name: 'UseNetwork',
      meta: {
        title: 'useNetwork'
      }
    }
  ]
}

export default RouteItem
