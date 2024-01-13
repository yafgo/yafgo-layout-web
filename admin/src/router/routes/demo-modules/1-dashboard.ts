import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/dashboard',
  redirect: '/dashboard/analysis',
  name: 'Dashboard',
  component: Layout,
  meta: {
    title: t('router.dashboard'),
    icon: 'ant-design:dashboard-filled',
    alwaysShow: true,
    order: 0
  },
  children: [
    {
      path: 'analysis',
      component: () => import('@/views/Dashboard/Analysis.vue'),
      name: 'Analysis',
      meta: {
        title: t('router.analysis'),
        noCache: true,
        affix: true
      }
    },
    {
      path: 'workplace',
      component: () => import('@/views/Dashboard/Workplace.vue'),
      name: 'Workplace',
      meta: {
        title: t('router.workplace'),
        noCache: true
      }
    }
  ]
}

export default RouteItem
