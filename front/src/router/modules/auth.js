export default [
  {
    path: "/login",
    component: () =>
      import(/* webpackChunkName: "auth" */ "@/views/Auth/Auth.vue"),
    
    children: [
      {
        path: "",
        name: "Login",
        component: () =>
          import(/* webpackChunkName: "login" */ "@/views/Auth/Login.vue"),
        meta: {
          title: "Seja bem-vindo!!!"
        }
      },
      {
        path: "/signup",
        name: "SignUp",
        component: () =>
          import (/* webpackChunkName: "signup" */ "@/views/Auth/SignUp.vue"),
        meta: {
          title: "Inscreva-se"
        }
      }
    ]
  }
]