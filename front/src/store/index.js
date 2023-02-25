import { createStore } from 'vuex'

//Modules
import Auth from './Auth'
import User from './User'

export default createStore({
  namespaced: true,
  modules: {
    Auth,
    User
  }
})
