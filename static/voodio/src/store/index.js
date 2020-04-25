import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'
import { getCurrFullHost } from '@/utils/url'

import tmdb from './modules/tmdb'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  state: {
    baseURLApi: getCurrFullHost()
  },
  modules: {
    tmdb
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})