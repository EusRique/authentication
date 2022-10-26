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
  await store.dispatch("User/checkAuth")
  const user = store.getters["User/getUser"];
  const isLoggedIn = user ? true : false

  const publicPages = [
    "Login"
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
