import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Vuesax from 'vuesax'
import routes from './routes'

import 'vuesax/dist/vuesax.css'
import 'material-icons/iconfont/material-icons.css'

Vue.config.productionTip = false

Vue.use(VueRouter)
Vue.use(Vuesax)

const router = new VueRouter({
  routes
})

new Vue({
  router,
  el: '#app',
  render: h => h(App),
}).$mount('#app')
