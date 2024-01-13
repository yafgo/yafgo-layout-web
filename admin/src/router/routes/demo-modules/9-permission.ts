import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/authorization',
  redirect: '/authorization/user',
  name: 'Authorization',
  component: Layout,
  meta: {
    title: t('router.authorization'),
    icon: 'eos-icons:role-binding',
    alwaysShow: true,
    order: 0
  },
  children: [
    {
      path: 'department',
      component: () => import('@/views/Authorization/Department/Department.vue'),
      name: 'Department',
      meta: {
        title: t('router.department')
      }
    },
    {
      path: 'user',
      component: () => import('@/views/Authorization/User/User.vue'),
      name: 'User',
      meta: {
        title: t('router.user')
      }
    },
    {
      path: 'menu',
      component: () => import('@/views/Authorization/Menu/Menu.vue'),
      name: 'Menu',
      meta: {
        title: t('router.menuManagement')
      }
    },
    {
      path: 'role',
      component: () => import('@/views/Authorization/Role/Role.vue'),
      name: 'Role',
      meta: {
        title: t('router.role')
      }
    }
  ]
}

export default RouteItem
