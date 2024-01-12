import { defineStore } from 'pinia'
import { store } from '../index'
import { UserLoginType, UserType } from '@/api/login/types'
import { ElMessageBox } from 'element-plus'
import { useI18n } from '@/hooks/web/useI18n'
import { apiLogin } from '@/api/login'
import { useTagsViewStore } from './tagsView'
import router from '@/router'
import { apiGetUserInfo } from '@/api/user'

interface UserState {
  userInfo?: UserType
  tokenKey: string
  token: string
  roleRouters?: string[] | AppCustomRouteRecordRaw[]
  rememberMe: boolean
  loginInfo?: UserLoginType
}

export const useUserStore = defineStore('user', {
  persist: {
    key: 'yafgo-user'
  },
  state: (): UserState => {
    return {
      userInfo: undefined,
      tokenKey: 'Authorization',
      token: '',
      roleRouters: undefined,
      // 记住我
      rememberMe: true,
      loginInfo: undefined
    }
  },
  getters: {
    getTokenKey(): string {
      return this.tokenKey
    },
    getToken(): string {
      return this.token
    },
    getUserInfo(): UserType | undefined {
      return this.userInfo
    },
    getRoleRouters(): string[] | AppCustomRouteRecordRaw[] | undefined {
      return this.roleRouters
    },
    getRememberMe(): boolean {
      return this.rememberMe
    },
    getLoginInfo(): UserLoginType | undefined {
      return this.loginInfo
    }
  },
  actions: {
    setTokenKey(tokenKey: string) {
      this.tokenKey = tokenKey
    },
    setToken(token: string) {
      this.token = token
      // setToken(token)
    },
    setUserInfo(userInfo?: UserType) {
      this.userInfo = userInfo
    },
    setRoleRouters(roleRouters: string[] | AppCustomRouteRecordRaw[]) {
      this.roleRouters = roleRouters
    },
    logoutConfirm() {
      const { t } = useI18n()
      ElMessageBox.confirm(t('common.loginOutMessage'), t('common.reminder'), {
        confirmButtonText: t('common.ok'),
        cancelButtonText: t('common.cancel'),
        type: 'warning'
      }).then(async () => {
        const res = await this.logout()
        if (res) {
          this.reset()
        }
      })
    },
    reset() {
      const tagsViewStore = useTagsViewStore()
      tagsViewStore.delAllViews()
      this.setToken('')
      this.setUserInfo(undefined)
      this.setRoleRouters([])
      router.replace('/login')
    },
    setRememberMe(rememberMe: boolean) {
      this.rememberMe = rememberMe
    },
    setLoginInfo(loginInfo: UserLoginType | undefined) {
      this.loginInfo = loginInfo
    },
    async login(formData: UserLoginType, rememberMe?: boolean) {
      const res = await apiLogin(formData)
      if (!res) {
        return res
      }
      // 记住我
      if (rememberMe) {
        this.setLoginInfo({
          username: formData.username,
          password: formData.password
        })
      } else {
        this.setLoginInfo(undefined)
      }
      this.setToken(res.data?.token || '')
      this.setRememberMe(rememberMe || false)
      this.setUserInfo(res.data)
      return res
    },
    async logout() {
      // await apiLogout()
      this.reset()
      return true
    },
    async fetchUserInfo() {
      const res = await apiGetUserInfo()
      this.setUserInfo(res.data)
    }
  }
})

export const useUserStoreWithOut = () => {
  return useUserStore(store)
}
