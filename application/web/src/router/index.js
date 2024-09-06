import Vue from 'vue'
import Router from 'vue-router'
import ModelUpload from '@/views/ModelManagement/ModelUpload.vue';
import ModelList from '@/views/ModelManagement/ModelList.vue';
import ModelDetails from '@/views/ModelManagement/ModelDetails.vue';

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
import { component } from 'vue/types/umd';
import { title } from '@/settings';

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/uplink',
    children: [{
      path: 'uplink',
      name: 'Uplink',
      component: () => import('@/views/uplink/index'),
      meta: { title: '溯源信息录入', icon: 'el-icon-edit-outline' }
    }]
  },
  {
    path: '/trace',
    component: Layout,
    children: [{
      path: 'trace',
      name: 'Trace',
      component: () => import('@/views/trace/index'),
      meta: { title: '溯源查询', icon: 'el-icon-search' }
    }]
  },
  {
    path: '/model-management',
    component: Layout,
    redirect: '/model-management/list',
    name: 'ModelManagement',
    meta: { title: 'AI模型管理', icon: 'el-icon-ai' },
    children: [
      {
        path: 'upload',
        name: 'ModelUpload',
        component: ModelUpload,
        meta: { title: '上传模型', icon: 'el-icon-upload' }
      },
      {
        path: 'list',
        name: 'ModelList',
        component: ModelList,
        meta: { title: '模型列表', icon: 'el-icon-list' }
      },
      {
        path: 'details/:id',
        name: 'ModelDetails',
        component: ModelDetails,
        meta: { title: '模型详情', icon: 'el-icon-info' },
        hidden: true
      }
    ]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'http://114.55.9.118:8080',
        meta: { title: '区块链浏览器', icon: 'el-icon-discover' }
      }
    ]
  },
  {
    path: 'external-link',
    component: Layout,
    children: [
      {
        path: 'http://114.55.9.118:7880',
        meta: {title: '大模型使用', icon: 'el-icon-message'}
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true },
  
]

const createRouter = () => new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
