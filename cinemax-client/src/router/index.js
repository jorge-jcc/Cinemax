import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/funciones/create_funcion',
    name: 'CreateFunciones',
    component: () => import('../views/CreateFuncion.vue'),
  },
  {
    path: '/funciones',
    name: 'Funciones',
    component: () => import('../views/Funciones.vue'),
  },
  {
    path: '/peliculas/create_pelicula',
    name: 'CreatePeliculas',
    component: () => import('../views/CreatePelicula.vue')
  },
  {
    path: '/peliculas',
    name: 'Peliculas',
    component: () => import('../views/Peliculas.vue'),
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
