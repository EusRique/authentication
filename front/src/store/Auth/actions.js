import auth from '@/api/auth'
import * as types from './mutation-types'

export default {
  setLogin: async ({ dispatch, commit }, payload) => {
    try {
      const data = await auth.login(payload)
      commit(types.SET_TOKEN, data.data.results)
    } catch (error) {
      console.log('ERROR')
      console.log(error)
    }
  }
}  
