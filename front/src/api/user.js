import axios from '@/plugins/axios'

const registerUSer = async payload => await axios.post(`/signup`, { ...payload })

export default {
  registerUSer
}