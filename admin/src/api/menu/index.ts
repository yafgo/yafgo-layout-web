import request from '@/axios'

export const getMenuListApi = () => {
  return request.get({ url: '/mock/menu/list' })
}

export const apiGetMenuList = () => {
  return request.get({ url: '/api/admin/menu' })
}
