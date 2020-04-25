import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faFilm, faCogs } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import router from './router'
import store from './store'

import App from './App.vue'

import '@/assets/css/tailwind.css'

Vue.config.productionTip = false

library.add([
  faFilm,
  faCogs
])

Vue.component('fa-icon', FontAwesomeIcon)

new Vue({
  el: '#app',
  router,
  store,
  render: (h) => h(App),
})
