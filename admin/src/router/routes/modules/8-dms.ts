import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/dms',
  component: Layout,
  redirect: '/dms/list',
  name: 'GroupDms',
  meta: {
    title: t('router.dms'),
    icon: 'flat-color-icons:database'
  },
  children: [
    {
      path: 'list',
      component: () => import('@/views/dms/manage/index.vue'),
      name: 'DmsList',
      meta: {
        title: t('router.dmsList'),
        icon: 'flat-color-icons:database'
      }
    },
    {
      path: 'detail',
      component: () => import('@/views/dms/detail/index.vue'),
      name: 'DmsDetail',
      meta: {
        title: t('router.dmsDetail'),
        icon: 'flat-color-icons:data-sheet'
      }
    }
  ]
}

export default RouteItem
