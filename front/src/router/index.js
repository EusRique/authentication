import store from '@/store'
import { createRouter, createWebHistory } from 'vue-router'

//Modules
import Auth from "@/router/modules/auth.js"

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "home" */ '@/views/Home.vue'),
    meta: {
        title: 'Seja bem Vindo'
    }
  },

  ...Auth
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach(async (to, from, next) => {
  await store.dispatch('checkAuth')
  const user = store.getters['getUser']
  const isLoggedIn = Object.keys(user).length === 0 ? false : true

  const publicPages = [
    "Login",
    "SignUp"
  ]

  const authRequired = !publicPages.includes(to.name)

  if (authRequired && !isLoggedIn) {
    return next("/login")
  } else if (!authRequired && isLoggedIn) {
    return next("/")
  }

  return next();
})

export default router
