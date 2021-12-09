<template>
  <main-layout>
    <v-row justify="center">
      <v-col cols="12" md="10" lg="9">
        <v-sheet min-height="70vh" rounded="lg">
          <v-form ref="form" v-model="valid">
            <v-row class="px-5 pb-2 mt-0" justify="center">
              <v-col cols="12" class="text-center">
                <h1>Registrar Película</h1>
              </v-col>
              <v-col cols="12" sm="10">
                <v-img v-if="imageURL" :src="imageURL" contain/>
                <v-skeleton-loader  v-else type="image"></v-skeleton-loader>
                <v-file-input :rules="rules"
                  accept="image/*" 
                  label="Imagen"
                  v-model="pelicula.image"
                  @change="onFileChange"
                />
                <v-row>
                  <v-col cols="7">
                    <v-text-field v-model="pelicula.nombre"
                      :rules="rules" 
                      label="Nombre"
                    />
                  </v-col>
                  <v-col cols="5"> 
                    <v-text-field v-model="pelicula.director"
                      :rules="rules" 
                      label="Director"
                    />
                  </v-col>
                </v-row>
                <v-textarea v-model="pelicula.descripcion" 
                  :rules="rules"
                  auto-grow 
                  label="Descripcion" 
                  rows="3"
                />
                <v-row>
                  <v-col cols="7">
                    <v-slider
                      v-model="pelicula.duracionMinutos"
                      class="mt-5"
                      label="Duración en minutos"
                      min="60"
                      max="300"
                      thumb-label="always"
                      hide-details=""
                    />
                  </v-col>
                  <v-col cols="2">
                    <v-text-field v-model="pelicula.anio"
                      :rules="rules" 
                      label="Año"
                    />
                  </v-col>
                  <v-col cols="3">
                    <v-select v-model="pelicula.clasificacionId"
                      :rules="[v => !!v || 'Item requerido']"
                      :items="clasificaciones" 
                      label="Clasificación" 
                      item-text="clave" 
                      item-value="id" 
                    />
                  </v-col>
                </v-row>
                <v-row class="my-0">
                  <v-col>
                    <v-select v-model="pelicula.idiomaId"
                    :rules="[v => !!v || 'Item requerido']"
                    :items="idiomas" 
                    label="Idioma" 
                    item-text="nombre" 
                    item-value="id"
                  />
                  </v-col>
                  <v-col>
                    <v-select  v-model="pelicula.subtituloId" 
                      :rules="[v => !!v || 'Item requerido']"
                      :items="idiomas" 
                      label="Subtitulos" 
                      item-text="nombre" 
                      item-value="id"
                    />
                  </v-col>
                </v-row>
                <v-row class="mt-0">
                  <v-col cols="6">
                    <v-menu
                      v-model="menuFecha"
                      transition="scale-transition"
                      offset-y
                      min-width="auto"
                    >
                        <template v-slot:activator="{ on, attrs }">
                          <v-text-field v-model="pelicula.fechaDisponibilidad"
                            :rules="rules"
                            label="Fecha de disponiblidad"
                            prepend-icon="mdi-calendar"
                            readonly
                            v-bind="attrs"
                            v-on="on"
                          ></v-text-field>
                        </template>
                        <v-date-picker color="blue lighten-1"
                          v-model="pelicula.fechaDisponibilidad"
                          no-title
                          scrollable
                          lang="es"
                        />
                    </v-menu>
                  </v-col>
                  {{pelicula.generoId}}
                  <v-col>
                    <v-select  v-model="pelicula.generoId"
                      :rules="[v => !!v || 'Item requerido']"
                      :items="generos" 
                      label="Género" 
                      item-text="nombre" 
                      item-value="id"
                    />
                  </v-col>
                </v-row>
                <v-textarea v-model="pelicula.resena"
                  :rules="rules"
                  clearable 
                  clear-icon="mdi-close-circle" 
                  label="Reseña"
                  rows="3"
                ></v-textarea>
              </v-col>
              <v-col class="text-center" cols="12">
                <v-btn rounded color="primary" @click="savePelicula">Guardar</v-btn>
              </v-col>
            </v-row>
          </v-form>
        </v-sheet>
      </v-col>
    </v-row>
    <dialog-error />
  </main-layout>
</template>

<script>
import moment from 'moment'
import { mapMutations } from 'vuex'
import { CreatePeliculaAPI, getClasificaciones, getGeneros, getIdiomas, LoadImagePeliculaAPI } from '../api/pelicula'
import DialogError from '../components/DialogError.vue'
import MainLayout from '../layout/MainLayout.vue'

export default {
  name : 'CreatePelicula',
  components: { MainLayout, DialogError },
  data: () => ({
    activePcker: 'YEAR',
    menuYear: false,
    menuFecha: false,
    imageURL : null,
    fechaDisponibilidad: null,
    year: null,
    idiomas: [],
    clasificaciones: [],
    generos: [],
    rules: [
        v => !!v || 'El campo es requerido',
    ],
    valid: false,
    pelicula:{
      nombre: null,
      director: null,
      descripcion: null,
      duracionMinutos: null,
      anio: null,
      fechaDisponibilidad: null,
      image: null,
      resena: null,
      clasificacionId: null,
      idiomaId: null,
      subtituloId: null,
      generoId: null,
    }
  }),
  methods:{
    ...mapMutations(["setDialogError", "setDialogErrorMessage"]),
    onFileChange: function(file) {
      if (file)
        this.imageURL = URL.createObjectURL(file)
      else 
        this.imageURL = null
    },
    savePelicula: function(){
      if (!this.valid) {
        this.$refs.form.validate()
        return
      }
      moment.locale("es")
      CreatePeliculaAPI({
        ...this.pelicula,
        fechaDisponibilidad: moment(this.pelicula.fechaDisponibilidad).toISOString()
      }).then(res => {
        if(res.status == 200){
          LoadImagePeliculaAPI(this.pelicula.image, res.data.peliculaId).then(()=>{
            this.$router.push({ path: '/peliculas' })
          })
        }
      }).catch(()=>{
        this.setDialogErrorMessage("No se pudo registrar exitosamente la película")
        this.setDialogError(true)
      })
    }
  },
  computed: {
  },
  watch:{
    menuYear(val){
      val && this.$nextTick(() => (this.$refs.picker.internalActivePicker  = 'YEAR'))
    }
  },
  created: function(){
    getIdiomas().then((res) => {
      this.idiomas = res.data.idiomas
    })
    getClasificaciones().then((res) => {
      this.clasificaciones = res.data.clasificaciones
    })
    getGeneros().then((res) => {
      this.generos = res.data.generos
    })
  },
}
</script>