import auth from '@/api/auth'
import AuthModel from '../Auth/model'
import VueJwtDecode from 'vue-jwt-decode'

export default {
  setLogin: async ({ commit, dispatch }, payload) => {
    try {
      const data = await auth.login(payload)
      localStorage.setItem('token', data.data.token)

      const userData = await dispatch('decodeToken', data.data.token)
      
      const user = AuthModel.mountUserData(userData)
      
      commit('setToken', data.data.token)
      commit('setUser', user)

    } catch (error) {
      console.log('ERROR')
      console.log(error)
    }
  },

  decodeToken: async ({ }, token) => {
    try {
      if (!token) return

      const decodeToken = await VueJwtDecode.decode(token)

      return decodeToken
    } catch (error) {
      console.log("ERROR")
      console.log(error)
    }
  },

  checkAuth: async ({ commit, dispatch }) => {
    try {
      const token = localStorage.getItem('token')
      if (!token) return

      let user = await dispatch("decodeToken", token)
      if (!user) return

      const userData = AuthModel.mountUserData(user)
      commit(types.SET_USER, userData)
    } catch (error) {
      
    }
  }
}  
