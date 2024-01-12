export interface UserLoginType {
  username: string
  password: string
}

export interface UserType {
  id: number;
  name?: string;
  avatar?: string;
  phone: string;
  username: string
  password: string
  role: string
  roleId: string
  token: string
}
