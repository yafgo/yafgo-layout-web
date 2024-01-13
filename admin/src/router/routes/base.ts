import { Layout } from '@/utils/routerHelper'

export const ROUTE_REDIRECT: AppRouteRecordRaw = {
  path: '/redirect',
  name: 'redirectWrapper',
  component: Layout,
  meta: {
    hidden: true,
    noTagsView: true
  },
  children: [
    {
      path: '/redirect/:path',
      name: 'Redirect',
      component: () => import('@/views/Redirect/Redirect.vue'),
      meta: {
        hidden: true,
        noTagsView: true
      }
    }
  ]
}

export const ROUTE_NOT_FOUND: AppRouteRecordRaw = {
  // path: '/:pathMatch(.*)*',
  path: '/404',
  name: 'notFound',
  component: () => import('@/views/Error/404.vue'),
  meta: {
    hidden: true,
    title: '404',
    noTagsView: true
  }
}
