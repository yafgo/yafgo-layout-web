import request from '@/axios'
import { UserType } from '../login/types'

export const apiGetUserInfo = (): Promise<IResponse<UserType>> => {
  return request.get({ url: '/api/v1/user/info' })
}
