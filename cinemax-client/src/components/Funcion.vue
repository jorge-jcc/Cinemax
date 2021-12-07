<template>
  <v-card class="px-0 my-3">
    <v-row class="mx-0 px-0">
      <v-col cols="12" md="3">
        <v-img aria-hidden="" :src="image" />
      </v-col>
      <v-col>
        <v-card-title class="pa-0">
          <span>{{pelicula.nombre}}</span>
        </v-card-title>
        <v-chip class="ma-2" color="blue darken-3" x-small label text-color="white">
          {{pelicula.clasificacion}}
        </v-chip>
        <v-icon color="blue darken-3" class="ml-2 mr-0">mdi-clock-outline</v-icon><span class="text-caption">
          {{pelicula.duracionMinutos}} min
        </span>
        <v-icon color="blue darken-3" class="ml-2 mr-0">mdi-filmstrip</v-icon><span class="text-caption">
          {{pelicula.genero}}
        </span>
        <v-col cols="6" class="py-0 px-2">
          <span class="text-caption">Idioma: 
            <span class="font-weight-bold">
              {{pelicula.idioma}}
            </span>
          </span>
        </v-col>
        <v-col cols="6" class="py-0 px-2">
          <span class="text-caption">Subtítulos: <span class="font-weight-bold">
            {{pelicula.subtitulo}}
            </span></span>
        </v-col>
        <v-col>
          <v-row justify="center">
            <v-col class="px-2">
              <v-chip v-for="horario in pelicula.horarios" :key="horario" 
                class="mr-2 mb-2" 
                label 
                color="blue darken-4" 
                outlined
              >
                {{horario}}
              </v-chip>
            </v-col>
          </v-row>
        </v-col>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
import { API_HOST } from '../utils/constant'

export default {
  name: "Funcion",
  props: {
    pelicula: Object,
  },
  data: () => ({
    image: null,
  }),
  methods: {

  },
  computed:{

  },
  created: function(){
    this.image = `${API_HOST}/pelicula/image?pelicula_id=${this.pelicula.id}`

  }
}
</script>

<style>

.descripcion{
  text-align: justify;
  text-justify: inter-word;
}

</style>