'use strict'

import axios from 'axios'
import store from '../store/index'

const config = {
  baseURL: 'http://localhost:3000/v1/api',
  timeout: 60 * 1000
}

let _axios = axios.create(config)
_axios.interceptors.request.use(
  function(config) {
    const isLoggedIn = store.getters['getUser'] ? true : false
    if (isLoggedIn) {
      config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
    }
    return config
  }
)

export default _axios