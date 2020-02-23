import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios';
import Card from 'primevue/card';
import InputText from 'primevue/inputtext';
import Panel from 'primevue/panel';

//npm install vue-cookie --save
var VueCookie = require('vue-cookie');


Vue.config.productionTip = false

Vue.component('InputText', InputText);
Vue.component('Card', Card);
Vue.component('Panel', Panel);
//Vue.use(VueAxios, axios)
Vue.use({
  install(Vue) {   

    Vue.prototype.$http = axios.create({
      // baseURL: 'http://localhost:8000/',
      baseURL: 'http://192.168.1.141:8000/',
      // headers: {
      //   'Access-Control-Allow-Origin': '*',
      //   "Authorization": "1234"
      // },
    })  


    //https://stackoverflow.com/questions/55883984/vue-axios-cors-policy-no-access-control-allow-origin


    // axios.create({
    //   baseURL: 'http://localhost:8000',
    //   withCredentials: false,
    //   headers: {
    //     'Content-Type': 'application/json',
    //     'Authorization': 'Bearer ' + VueCookie.get('tokencliente'),
    //     'Access-Control-Allow-Origin': '*',
    //     'Accept' : 'application/json, text/plain, */*',
    //     'Access-Control-Allow-Methods' : 'GET, PUT, POST, DELETE, OPTIONS',
    //     'Access-Control-Allow-Credentials' : true
    //   }
    // })
    
    // Vue.prototype.$http  =  Axios;
    // const accessToken = VueCookie.get('tokencliente')

    // console.log("Token em main.js:" + accessToken)
    // if (accessToken) {
    //     Vue.prototype.$http.defaults.headers.common['Authorization'] =  accessToken
    // }

    Vue.prototype.$http.interceptors.request.use(config => {
      // console.log(config.method)
      return config
    }, error => Promise.reject(error))

    Vue.prototype.$http.interceptors.response.use(resp => {
      // const array = []
      // for (let chave in resp.data) {
      //   array.push ({id: chave, ...resp.data[chave]})
      // }
      // resp.data = array
      return resp
    }, error => Promise.reject(error))

  }
})


Vue.use(VueCookie);

new Vue({
  router,
  render: function (h) { return h(App) }
}).$mount('#app')
