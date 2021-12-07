<template>
  <main-layout>
    <v-row justify="center">
      <v-col cols="8">
        <v-row class="mt-2 my-0">
          <v-col cols="9">
            <v-text-field dense 
              clearable
              label="Buscar una película"
              placeholder="Escribe la película"
              prepend-icon="mdi-magnify"
            ></v-text-field>
          </v-col>
          <v-col class="text-end"> 
            <v-btn class="mx-2" fab dark dense small color="red darken-4" @click="crearPelicula">
              <v-icon dark>mdi-plus</v-icon>
            </v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" md="4" lg="3" v-for="pelicula in peliculas" :key="pelicula.id">
        <pelicula :pelicula="pelicula" />
      </v-col>
    </v-row>
  </main-layout>
</template>

<script>
import MainLayout from '../layout/MainLayout.vue'
import Pelicula from '../components/Pelicula.vue'
import { getPeliculasAPI } from '../api/pelicula'

export default {
  name : 'Peliculas',
  components: { MainLayout, Pelicula },
  data: () => ({
    peliculas: [],
  }),
  methods:{
    crearPelicula: function(){
      this.$router.push({ path: '/peliculas/create_pelicula' })
    },
  },
  computed: {

  },
  mounted() {
  },
  watch:{

  },
  created: function(){
    getPeliculasAPI().then(res => {
      if (res.status == 200){
        this.peliculas = res.data.peliculas
      }
    })
  },
}
</script>
<style scoped>

</style>