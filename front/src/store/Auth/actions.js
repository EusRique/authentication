import auth from '@/api/auth'
import AuthModel from '../Auth/model'
import VueJwtDecode from 'vue-jwt-decode'
import router from '@/router'
import { useToast } from 'vue-toastification'

const toast = useToast()

export default {
  setLogin: async ({ commit, dispatch }, payload) => {
    try {
      const data = await auth.login(payload)
      localStorage.setItem('token', data.data.token)

      return router.push("/")

    } catch (error) {
      console.log('ERROR')
      console.error(error.response.data)
      if (Array.isArray(error.response.data.message)) {
        error.response.data.message.forEach(message => {
          toast.error(message)
        });
      } else {
        toast.error(error.response.data.message)
      }
      return
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
      commit('setUser', userData)
    } catch (error) {
      console.log('ERROR')
      console.log(error)
    }
  }
}  
