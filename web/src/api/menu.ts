import axios from 'axios';

export interface MenuItem {
  id: number;
  pid: number;
  path: string;
  name: string;
  children: MenuItem[];
}

export interface MenuParams extends Partial<MenuItem> {
  current: number;
  pageSize: number;
}

export function ApiGetList() {
  return axios.get<MenuItem[]>('/api/admin/menu');
}

export function ApiGetDetail(id: number) {
  return axios.get<string>(`/api/admin/menu/menus/${id}`);
}

export function ApiCreate(content: string) {
  return axios.post('/api/admin/menu/menus', { content });
}

export function ApiUpdate(id: number, content: string) {
  return axios.post(`/api/admin/menu/menus/${id}`, { content });
}

export function ApiDelete(id: number) {
  return axios.post(`/api/admin/menu/menus/${id}`);
}
