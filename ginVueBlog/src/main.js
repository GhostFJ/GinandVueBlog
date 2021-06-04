import Vue from 'vue'
import Vuelidate from 'vuelidate';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import '@/plugins/element.js'
// import FameUtil from '@/utils/fame'
import axios from 'axios';
import VueAxios from 'vue-axios';
import router from './router'
import store from './store';
import dayjs from 'dayjs'
import App from './App.vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import '@/assets/css/animations.css';


Vue.config.productionTip = false

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.prototype.$dayjs = dayjs
Vue.prototype.$util = FameUtil.FUNCTIONS
Vue.prototype.$static = FameUtil.STATIC

// vuelidate
Vue.use(Vuelidate);

// axios
Vue.use(VueAxios, axios);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
