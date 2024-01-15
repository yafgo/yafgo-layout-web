import { useI18n } from '@/hooks/web/useI18n'
import { Layout } from '@/utils/routerHelper'

const { t } = useI18n()

const RouteItem: AppRouteRecordRaw = {
  path: '/components',
  name: 'ComponentsDemo',
  component: Layout,
  meta: {
    title: t('router.component'),
    icon: 'bx:bxs-component',
    alwaysShow: true
  },
  children: [
    {
      path: 'form',
      // component: getParentLayout(),
      redirect: '/components/form/default-form',
      name: 'Form',
      meta: {
        title: t('router.form'),
        alwaysShow: true
      },
      children: [
        {
          path: 'default-form',
          component: () => import('@/views/demo/Components/Form/DefaultForm.vue'),
          name: 'DefaultForm',
          meta: {
            title: t('router.defaultForm')
          }
        },
        {
          path: 'use-form',
          component: () => import('@/views/demo/Components/Form/UseFormDemo.vue'),
          name: 'UseForm',
          meta: {
            title: 'UseForm'
          }
        }
      ]
    },
    {
      path: 'table',
      redirect: '/components/table/default-table',
      name: 'TableDemo',
      meta: {
        title: t('router.table'),
        alwaysShow: true
      },
      children: [
        {
          path: 'default-table',
          component: () => import('@/views/demo/Components/Table/DefaultTable.vue'),
          name: 'DefaultTable',
          meta: {
            title: t('router.defaultTable')
          }
        },
        {
          path: 'use-table',
          component: () => import('@/views/demo/Components/Table/UseTableDemo.vue'),
          name: 'UseTable',
          meta: {
            title: 'UseTable'
          }
        },
        {
          path: 'tree-table',
          component: () => import('@/views/demo/Components/Table/TreeTable.vue'),
          name: 'TreeTable',
          meta: {
            title: t('router.treeTable')
          }
        },
        {
          path: 'table-image-preview',
          component: () => import('@/views/demo/Components/Table/TableImagePreview.vue'),
          name: 'TableImagePreview',
          meta: {
            title: t('router.PicturePreview')
          }
        },
        {
          path: 'table-video-preview',
          component: () => import('@/views/demo/Components/Table/TableVideoPreview.vue'),
          name: 'TableVideoPreview',
          meta: {
            title: t('router.tableVideoPreview')
          }
        },
        {
          path: 'card-table',
          component: () => import('@/views/demo/Components/Table/CardTable.vue'),
          name: 'CardTable',
          meta: {
            title: t('router.cardTable')
          }
        }
      ]
    },
    {
      path: 'editor-demo',
      // component: getParentLayout(),
      redirect: '/components/editor-demo/editor',
      name: 'EditorDemo',
      meta: {
        title: t('router.editor'),
        alwaysShow: true
      },
      children: [
        {
          path: 'editor',
          component: () => import('@/views/demo/Components/Editor/Editor.vue'),
          name: 'Editor',
          meta: {
            title: t('router.richText')
          }
        },
        {
          path: 'json-editor',
          component: () => import('@/views/demo/Components/Editor/JsonEditor.vue'),
          name: 'JsonEditor',
          meta: {
            title: t('router.jsonEditor')
          }
        }
      ]
    },
    {
      path: 'search',
      component: () => import('@/views/demo/Components/Search.vue'),
      name: 'Search',
      meta: {
        title: t('router.search')
      }
    },
    {
      path: 'descriptions',
      component: () => import('@/views/demo/Components/Descriptions.vue'),
      name: 'Descriptions',
      meta: {
        title: t('router.descriptions')
      }
    },
    {
      path: 'image-viewer',
      component: () => import('@/views/demo/Components/ImageViewer.vue'),
      name: 'ImageViewer',
      meta: {
        title: t('router.imageViewer')
      }
    },
    {
      path: 'dialog',
      component: () => import('@/views/demo/Components/Dialog.vue'),
      name: 'Dialog',
      meta: {
        title: t('router.dialog')
      }
    },
    {
      path: 'icon',
      component: () => import('@/views/demo/Components/Icon.vue'),
      name: 'Icon',
      meta: {
        title: t('router.icon')
      }
    },
    {
      path: 'icon-picker',
      component: () => import('@/views/demo/Components/IconPicker.vue'),
      name: 'IconPicker',
      meta: {
        title: t('router.iconPicker')
      }
    },
    {
      path: 'echart',
      component: () => import('@/views/demo/Components/Echart.vue'),
      name: 'Echart',
      meta: {
        title: t('router.echart')
      }
    },
    {
      path: 'count-to',
      component: () => import('@/views/demo/Components/CountTo.vue'),
      name: 'CountTo',
      meta: {
        title: t('router.countTo')
      }
    },
    {
      path: 'qrcode',
      component: () => import('@/views/demo/Components/Qrcode.vue'),
      name: 'Qrcode',
      meta: {
        title: t('router.qrcode')
      }
    },
    {
      path: 'highlight',
      component: () => import('@/views/demo/Components/Highlight.vue'),
      name: 'Highlight',
      meta: {
        title: t('router.highlight')
      }
    },
    {
      path: 'infotip',
      component: () => import('@/views/demo/Components/Infotip.vue'),
      name: 'Infotip',
      meta: {
        title: t('router.infotip')
      }
    },
    {
      path: 'input-password',
      component: () => import('@/views/demo/Components/InputPassword.vue'),
      name: 'InputPassword',
      meta: {
        title: t('router.inputPassword')
      }
    },
    {
      path: 'waterfall',
      component: () => import('@/views/demo/Components/Waterfall.vue'),
      name: 'waterfall',
      meta: {
        title: t('router.waterfall')
      }
    },
    {
      path: 'image-cropping',
      component: () => import('@/views/demo/Components/ImageCropping.vue'),
      name: 'ImageCropping',
      meta: {
        title: t('router.imageCropping')
      }
    },
    {
      path: 'video-player',
      component: () => import('@/views/demo/Components/VideoPlayer.vue'),
      name: 'VideoPlayer',
      meta: {
        title: t('router.videoPlayer')
      }
    }
  ]
}

export default RouteItem
