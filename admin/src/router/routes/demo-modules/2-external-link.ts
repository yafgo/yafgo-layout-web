import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw[] = [
  {
    path: '/external-link',
    name: 'ExternalLink',
    component: Layout,
    meta: {
      order: 0
    },
    children: [
      {
        path: 'https://element-plus-admin-doc.cn/',
        name: 'DocumentLink',
        meta: {
          title: t('router.document'),
          icon: 'clarity:document-solid'
        }
      }
    ]
  },
  {
    path: '/guide',
    name: 'Guide',
    component: Layout,
    meta: {
      order: 0
    },
    children: [
      {
        path: 'index',
        component: () => import('@/views/Guide/Guide.vue'),
        name: 'GuideDemo',
        meta: {
          title: t('router.guide'),
          icon: 'cib:telegram-plane'
        }
      }
    ]
  }
]

export default RouteItem
