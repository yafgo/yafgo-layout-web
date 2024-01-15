import request from '@/axios'

export const getMenuListApi = () => {
  return request.get({ url: '/mock/menu/list' })
}

export const apiGetMenuList = () => {
  return request.get({ url: '/api/admin/menu' })
}

export const apiGetDetail = (id: number) => {
  return request.get<string>({ url: `/api/admin/menu/menus/${id}` })
}

export const apiSave = (data: any) => {
  const id = data.id ? '/' + data.id : ''
  return request.post({ url: `/api/admin/menu/menus${id}`, data: data })
}

export const apiDelete = (id: number) => {
  return request.post({ url: `/api/admin/menu/menus/${id}` })
}
