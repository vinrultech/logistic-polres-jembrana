import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import './plugins/swal';
import './plugins/toast';
import '@babel/polyfill'
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@fortawesome/fontawesome-free/css/all.css'

import '@/assets/style.css'

import axios from 'axios'
import progress from './plugins/progress_dialog';

import moment from 'moment';

import injector from 'vue-inject';




moment.locale("id")



Vue.use(progress)
Vue.use(injector)


const host = window.location.hostname
//const port = '80'
const port = '1331'

axios.defaults.baseURL = `http://${host}:${port}/api`

/*
axios.interceptors.request.use(function (config) {
  // Do something before request is sent
  
  
  let title = "Get"
  let text = "Sedang mendapatkan data dari server"
  if (config.url === "user/login") {
    title = "Login"
    text = "Sedang login ke server"
  } else {
    const token = utils.session();
    if (config.method === "post") {
      title = "Tambah"
      text = "Sedang tambah data ke server"
    } else if (config.method === "put") {
      title = "Update"
      text = "Sedang update data ke server"
    } else if (config.method === "delete") {
      title = "Hapus"
      text = "Sedang hapus data ke server"
    }

    config.headers = {
      'Content-Type': `application/json`,
      'Authorization': `Bearer ${token}`,
    }

  }
  config.progress.show({
    title: title,
    text: text,
  })
  return config;
}, function (error) {
  return Promise.reject(error);
});

axios.interceptors.response.use(function (response) {

  response.config.progress.hide()
  if (response.config.toast !== undefined && response.config.message) {
    response.config.toast.s(response.config.message)
  }
  return response;
}, function (error) {
  error.response.config.progress.hide()
  const code = parseInt(error.response.status)
  if (code === 400 || code === 401 || code === 403) {
    if (error.response.config.swal !== undefined && error.response.config.router !== undefined) {
      utils.error(error, error.response.config.swal, error.response.config.router)
    }
  } else if (error.response.config.method === "get") {
    if (error.response.config.swal !== undefined && error.response.config.router !== undefined) {
      utils.error(error, error.response.config.swal, error.response.config.router)
    }
  }


  return Promise.reject(error);
});

*/

Vue.axios = Vue.prototype.$axios = axios
Vue.host = Vue.prototype.$host = `http://${host}:${port}/`

injector.constant('axios', axios)
injector.constant('router', router)

import Default from './layouts/default'
import Admin from './layouts/admin'
import Login from './layouts/login'
import NotFound from './layouts/notfound'

Vue.component('default-layout', Default)
Vue.component('notfound-layout', NotFound)
Vue.component('admin-layout', Admin)
Vue.component('login-layout', Login)

Vue.filter('str_limit', function (value, size) {
  if (!value) return '';
  value = value.toString();

  if (value.length <= size) {
    return value;
  }
  return value.substr(0, size) + ' ....';
});

Vue.prototype.moment = moment

Vue.config.productionTip = false
Vue.mixin({
  methods: {
    breadcumText: (str) => {
      const text = str.replace("_", " ");
      //return text.charAt(0).toUpperCase() + text.slice(1)
      return text
        .toLowerCase()
        .split(' ')
        .map(word => word.charAt(0).toUpperCase() + word.slice(1))
        .join(' ');
    }
  }
})

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')