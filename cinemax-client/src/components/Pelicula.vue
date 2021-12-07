<template>
  <v-card :loading="loading" class="mx-auto mb-5" max-width="374">
    <template slot="progress">
      <v-progress-linear color="deep-purple" height="10" indeterminate/>
    </template>
    <v-img height="300" :src="image">
      <template v-slot:placeholder>
          <v-row class="fill-height ma-0" align="center" justify="center">
            <v-progress-circular indeterminate color="blue lighten-3"></v-progress-circular>
          </v-row>
        </template>
    </v-img>

    <v-card-title>
      <span class="col-12 pa-0" style="font-size:0.9em;">{{pelicula.nombre}}</span>
      <span class="text-caption font-weight-light my-0 pb-0">{{fecha}}</span>
    </v-card-title>

    <v-card-text class="pb-2">
      <v-row align="center" class="mx-0">
        <v-rating :value="4.5" color="amber" dense half-increments readonly size="20"/>
        <div class="grey--text ms-4">
          4.5 (413)
        </div>
      </v-row>
      <div class="mt-3 mb-2 text-subtitle-1">
        {{pelicula.director}}
      </div>

      <v-divider class="my-1"></v-divider>
    </v-card-text>

<v-card-text class="caption py-0">
      Clasificación: <span class="font-weight-medium">{{pelicula.clasificacion.clave}}</span>
    </v-card-text>
    <v-card-text class="caption py-0">
      Género: <span class="font-weight-medium">{{pelicula.genero.nombre}}</span>
    </v-card-text>
    <v-card-text class="caption py-0">
      <v-row>
        <v-col cols="">
        Idioma(s):<span class="font-weight-medium">{{pelicula.idioma.nombre}}</span>
        -
        Subtitulo(s):<span class="font-weight-medium">{{pelicula.subtitulos.nombre}}</span>
      </v-col>
      </v-row>
    </v-card-text>
    <!--
    <v-card-text>
      <div class="pb-0 descripcion">Por primera vez en la historia cinematográfica de Spider-Man, nuestro amistoso héroe y vecino es desenmascarado, y ya no puede separar su vida normal de los altos riesgos de ser un súper héroe</div>
    </v-card-text>
    -->


  </v-card>
</template>

<script>
import moment from 'moment'
import { API_HOST } from '../utils/constant'
export default {
  name: "Pelicula",
  props: {
    pelicula: Object,
  },
  data: () => ({
    loading: false,
    selection: 1,
  }),
  methods: {
    reserve () {
      this.loading = true
      setTimeout(() => (this.loading = false), 2000)
    },
  },
  computed:{
    fecha: function(){
      moment.locale("es")
      return moment(this.pelicula.fechaDisponibilidad).format('LL');
    },
    image: function(){
      if (!this.pelicula.id) return null
      return `${API_HOST}/pelicula/image?pelicula_id=${this.pelicula.id}`
    }
  },
  created: function(){
  }
}
</script>

<style>

.descripcion{
  text-align: justify;
  text-justify: inter-word;
}

</style>