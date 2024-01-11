import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const SYSTEM: AppRouteRecordRaw = {
  path: '/system',
  name: 'system',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: '系统设置',
    icon: 'icon-settings',
    requiresAuth: true,
    order: 39,
  },
  children: [
    {
      path: 'ycfg',
      name: 'Ycfg',
      component: () => import('@/views/system/ycfg/index.vue'),
      meta: {
        locale: '程序配置',
        icon: 'icon-settings',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default SYSTEM;
