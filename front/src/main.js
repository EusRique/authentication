import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Toast from './plugins/toastification'

//CSS
import "./assets/css/global.css"

createApp(App).use(store).use(router).use(Toast).mount('#app')
