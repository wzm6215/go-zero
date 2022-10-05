import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/layout/admin/adminList',
  name: 'Admin',
  component: () => import('@/view/person/person.vue')
},
{
  path: '/layout/bookstore/books',
  name: 'Books',
  component: () => import('@/view/bookstore/book/list.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
