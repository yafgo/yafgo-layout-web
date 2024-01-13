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
      component: () => import('@/views/hooks/useWatermark.vue'),
      name: 'UseWatermark',
      meta: {
        title: 'useWatermark'
      }
    },
    {
      path: 'useTagsView',
      component: () => import('@/views/hooks/useTagsView.vue'),
      name: 'UseTagsView',
      meta: {
        title: 'useTagsView'
      }
    },
    {
      path: 'useValidator',
      component: () => import('@/views/hooks/useValidator.vue'),
      name: 'UseValidator',
      meta: {
        title: 'useValidator'
      }
    },
    {
      path: 'useCrudSchemas',
      component: () => import('@/views/hooks/useCrudSchemas.vue'),
      name: 'UseCrudSchemas',
      meta: {
        title: 'useCrudSchemas'
      }
    },
    {
      path: 'useClipboard',
      component: () => import('@/views/hooks/useClipboard.vue'),
      name: 'UseClipboard',
      meta: {
        title: 'useClipboard'
      }
    },
    {
      path: 'useNetwork',
      component: () => import('@/views/hooks/useNetwork.vue'),
      name: 'UseNetwork',
      meta: {
        title: 'useNetwork'
      }
    }
  ]
}

export default RouteItem
