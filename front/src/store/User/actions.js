import user from '@/api/user'

export default {
  setRegisterUser: async ({ commit, dispatch }, payload) => {
    try {
      const data = await user.registerUSer(payload)

      return data
    } catch (error) {
      console.log('ERROR')
      console.error(error.response.data)
      return error.response.data
    }
  }
}