import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/error',
  redirect: '/error/404',
  name: 'Error',
  component: Layout,
  meta: {
    title: t('router.errorPage'),
    icon: 'ci:error',
    alwaysShow: true
  },
  children: [
    {
      path: '404-demo',
      component: () => import('@/views/Error/404.vue'),
      name: '404Demo',
      meta: {
        title: '404'
      }
    },
    {
      path: '403-demo',
      component: () => import('@/views/Error/403.vue'),
      name: '403Demo',
      meta: {
        title: '403'
      }
    },
    {
      path: '500-demo',
      component: () => import('@/views/Error/500.vue'),
      name: '500Demo',
      meta: {
        title: '500'
      }
    }
  ]
}

export default RouteItem
