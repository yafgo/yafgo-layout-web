import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const SYSTEM: AppRouteRecordRaw = {
  path: '/system',
  name: 'System',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.system',
    icon: 'icon-settings',
    requiresAuth: true,
    order: 39,
  },
  children: [
    {
      path: 'menu',
      name: 'Menu',
      component: () => import('@/views/system/menu/index.vue'),
      meta: {
        locale: 'menu.system.menu',
        icon: 'icon-menu',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'ycfg',
      name: 'Ycfg',
      component: () => import('@/views/system/ycfg/index.vue'),
      meta: {
        locale: 'menu.system.ycfg',
        icon: 'icon-code',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default SYSTEM;
