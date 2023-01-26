import { createStore } from 'vuex'

//Modules
import Auth from './Auth'

export default createStore({
  namespaced: true,
  modules: {
    Auth
  }
})
