import axios from '@/plugins/axios'

const login = async payload => await axios.post(`/login`, { ...payload })

export default {
  login
}