import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../components/login.vue'
import Hello from '../components/About.vue'
import Inicio from '../components/Inicio.vue'


Vue.use(VueRouter)

const routes = [
  // {
  //   path: '/',
  //   redirect: {
  //     name: "/login"
  //   }
  // },
  {
    path: '/',
    name: 'inicio',
    component: Inicio
  },
  {
    path: '/hello',
    name: 'hello',
    component: Hello
  },
 
  {
    path: '/home',
    name: 'home',
    component: Home
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router