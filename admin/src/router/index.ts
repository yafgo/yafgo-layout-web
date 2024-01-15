import { createRouter, createWebHashHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import type { App } from 'vue'
import { Layout } from '@/utils/routerHelper'
import { useI18n } from '@/hooks/web/useI18n'
import { NO_RESET_WHITE_LIST } from '@/constants'
import { appRoutes, demoRoutes } from './routes'
import { ROUTE_REDIRECT, ROUTE_NOT_FOUND } from './routes/base'
import createRouteGuard from './guard'

const { t } = useI18n()

export const constantRouterMap: AppRouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/dashboard/analysis',
    name: 'Root',
    component: Layout,
    meta: {
      hidden: true
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login/Login.vue'),
    meta: {
      hidden: true,
      title: t('router.login'),
      noTagsView: true
    }
  },
  ...(demoRoutes as AppRouteRecordRaw[]),
  ...(appRoutes as AppRouteRecordRaw[]),
  ROUTE_REDIRECT,
  ROUTE_NOT_FOUND
]

export const asyncRouterMap: AppRouteRecordRaw[] = [
  // ...(demoRoutes as AppRouteRecordRaw[]),
  // ...(appRoutes as AppRouteRecordRaw[])
]

const router = createRouter({
  history: createWebHashHistory(),
  strict: true,
  routes: constantRouterMap as RouteRecordRaw[],
  scrollBehavior: () => ({ left: 0, top: 0 })
})

createRouteGuard(router)

export const resetRouter = (): void => {
  router.getRoutes().forEach((route) => {
    const { name } = route
    if (name && !NO_RESET_WHITE_LIST.includes(name as string)) {
      router.hasRoute(name) && router.removeRoute(name)
    }
  })
}

export const setupRouter = (app: App<Element>) => {
  app.use(router)
}

export default router
