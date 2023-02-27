import user from '@/api/user'
import { useToast } from 'vue-toastification'

const toast = useToast()

export default {
  setRegisterUser: async ({ commit, dispatch }, payload) => {
    try {
      await user.registerUSer(payload)
      toast.success('Usuário registrado com sucesso!!!');
      return

    } catch (error) {
      console.log('ERROR')
      console.error(error)
      if (Array.isArray(error.response.data.message)) {
        error.response.data.message.forEach(message => {
          toast.error(message);
        });
      } else {
        toast.error(error.response.data.message);
      }
      return
    }
  }
}