import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/system',
  component: Layout,
  redirect: '/system/cfg',
  name: 'GroupSystem',
  meta: {
    title: t('router.system'),
    icon: 'flat-color-icons:settings'
  },
  children: [
    {
      path: 'settings',
      component: () => import('@/views/system/settings/Index.vue'),
      name: 'SysSettings',
      meta: {
        title: t('router.systemSetting'),
        icon: 'flat-color-icons:settings'
      }
    },
    {
      path: 'cfg',
      component: () => import('@/views/system/ycfg/Index.vue'),
      name: 'YCfg',
      meta: {
        title: t('router.systemYcfg'),
        icon: 'flat-color-icons:data-configuration'
      }
    }
  ]
}

export default RouteItem
