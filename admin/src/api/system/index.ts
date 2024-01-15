import request from '@/axios'

export const ApiGetAllCfg = (): Promise<IResponse> => {
  return request.get({ url: '/api/admin/system/cfg' })
}

export const ApiGetCfgInRedis = (): Promise<IResponse> => {
  return request.get({ url: '/api/admin/system/cfg_in_redis' })
}

export const ApiSetCfgInRedis = (content: string): Promise<IResponse> => {
  return request.post({
    url: '/api/admin/system/cfg_in_redis',
    data: { content: content }
  })
}
