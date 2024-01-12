import request from '@/axios'
import type { UserLoginType, UserType } from './types'

interface RoleParams {
  roleName: string
}

export const apiLogin = (data: UserLoginType): Promise<IResponse<UserType>> => {
  return request.post({ url: '/api/v1/user/login/username', data })
}

export const apiLogout = (): Promise<IResponse> => {
  return request.get({ url: '/mock/user/loginOut' })
}

export const getUserListApi = ({ params }: AxiosConfig) => {
  return request.get<{
    code: string
    data: {
      list: UserType[]
      total: number
    }
  }>({ url: '/mock/user/list', params })
}

export const getAdminRoleApi = (
  params: RoleParams
): Promise<IResponse<AppCustomRouteRecordRaw[]>> => {
  return request.get({ url: '/mock/role/list', params })
}

export const getTestRoleApi = (params: RoleParams): Promise<IResponse<string[]>> => {
  return request.get({ url: '/mock/role/list2', params })
}
