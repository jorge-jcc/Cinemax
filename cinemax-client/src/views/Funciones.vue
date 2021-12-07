<template>
  <main-layout>
    <v-row justify="center">
      <v-col cols="12" class="my-2"><h1 class="text-center">Cartelera: {{fecha}}</h1></v-col>
      <v-col cols="8" class="text-center py-0">
        <v-btn block @click="crearFuncion">Agregar Función</v-btn>
      </v-col>
      <v-col cols="8" v-for="pelicula in peliculas" :key="pelicula.id">
        <funcion :pelicula="pelicula" />
      </v-col>
    </v-row>
  </main-layout>
</template>

<script>
import moment from 'moment'
import { getCarteleraAPI } from '../api/cartelera'
import Funcion from '../components/Funcion.vue'
import MainLayout from '../layout/MainLayout.vue'

export default {
  name : 'Funciones',
  components: { MainLayout, Funcion, },
  data: () => ({
    peliculas: [],
  }),
  methods:{
    crearFuncion: function(){
      this.$router.push({ path: '/funciones/create_funcion' })
    },
  },
  computed: {
    fecha: function(){
      moment.locale("es")
      return moment().format('LL');
    },
  },
  mounted() {
  },
  watch:{

  },
  created: function(){
    getCarteleraAPI().then(res => {
      this.peliculas = res.data.peliculas
    })
  },
}
</script>
<style scoped>

</style>