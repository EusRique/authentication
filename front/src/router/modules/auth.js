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
      }
    ]
  }
]