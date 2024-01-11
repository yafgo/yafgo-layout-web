import axios from 'axios';

export function ApiGetAllCfg() {
  return axios.get('/api/admin/system/cfg');
}

export function ApiGetCfgInRedis() {
  return axios.get<string>('/api/admin/system/cfg_in_redis');
}

export function ApiSetCfgInRedis(content: string) {
  return axios.post('/api/admin/system/cfg_in_redis', { content });
}
