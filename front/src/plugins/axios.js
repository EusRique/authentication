'use strict'

import axios from 'axios'

const config = {
  baseURL: 'http://localhost:3000/v1/api',
  timeout: 60 * 1000
}

let _axios = axios.create(config)

_axios = axios.create(config)

export default _axios