import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faFilm } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import router from './router'

import App from './App.vue'

import '@/assets/css/tailwind.css'

Vue.config.productionTip = false

library.add([
  faFilm,
])

Vue.component('fa-icon', FontAwesomeIcon)

new Vue({
  el: '#app',
  router,
  render: (h) => h(App),
})
