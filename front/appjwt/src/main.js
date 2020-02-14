import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Card from 'primevue/card';
import InputText from 'primevue/inputtext';
import Panel from 'primevue/panel';


var VueCookie = require('vue-cookie');


Vue.config.productionTip = false

Vue.component('InputText', InputText);
Vue.component('Card', Card);
Vue.component('Panel', Panel);
Vue.use(VueCookie);

new Vue({
  router,
  render: function (h) { return h(App) }
}).$mount('#app')
