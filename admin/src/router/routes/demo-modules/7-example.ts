import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/example',
  redirect: '/example/example-dialog',
  name: 'Example',
  component: Layout,
  meta: {
    title: t('router.example'),
    icon: 'ep:management',
    alwaysShow: true
  },
  children: [
    {
      path: 'example-dialog',
      component: () => import('@/views/demo/Example/Dialog/ExampleDialog.vue'),
      name: 'ExampleDialog',
      meta: {
        title: t('router.exampleDialog')
      }
    },
    {
      path: 'example-page',
      component: () => import('@/views/demo/Example/Page/ExamplePage.vue'),
      name: 'ExamplePage',
      meta: {
        title: t('router.examplePage')
      }
    },
    {
      path: 'example-add',
      component: () => import('@/views/demo/Example/Page/ExampleAdd.vue'),
      name: 'ExampleAdd',
      meta: {
        title: t('router.exampleAdd'),
        noTagsView: true,
        noCache: true,
        hidden: true,
        canTo: true,
        activeMenu: '/example/example-page'
      }
    },
    {
      path: 'example-edit',
      component: () => import('@/views/demo/Example/Page/ExampleEdit.vue'),
      name: 'ExampleEdit',
      meta: {
        title: t('router.exampleEdit'),
        noTagsView: true,
        noCache: true,
        hidden: true,
        canTo: true,
        activeMenu: '/example/example-page'
      }
    },
    {
      path: 'example-detail',
      component: () => import('@/views/demo/Example/Page/ExampleDetail.vue'),
      name: 'ExampleDetail',
      meta: {
        title: t('router.exampleDetail'),
        noTagsView: true,
        noCache: true,
        hidden: true,
        canTo: true,
        activeMenu: '/example/example-page'
      }
    }
  ]
}

export default RouteItem
