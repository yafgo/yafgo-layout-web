import type { RouteRecordNormalized } from 'vue-router'

function formatModules(_modules: any, result: RouteRecordNormalized[]) {
  Object.keys(_modules).forEach((key) => {
    const defaultModule = _modules[key].default
    if (!defaultModule) return
    const moduleList = Array.isArray(defaultModule) ? [...defaultModule] : [defaultModule]
    result.push(...moduleList)
  })
  result.sort((a, b) => (a.meta?.order || 0) - (b.meta?.order || 0))
  return result
}

const modulesDemo = import.meta.glob('./demo-modules/*.ts', { eager: true })
const modules = import.meta.glob('./modules/*.ts', { eager: true })

export const demoRoutes: RouteRecordNormalized[] = formatModules(modulesDemo, [])
export const appRoutes: RouteRecordNormalized[] = formatModules(modules, [])
