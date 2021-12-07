<template>
  <main-layout>
    <v-row justify="center">
      <v-col cols="12" md="10" lg="9">
        <v-sheet min-height="70vh" rounded="lg">
          <v-form ref="form" v-model="valid">
            <v-row class="px-5 pb-2 mt-0" justify="center">
              <v-col cols="12" class="text-center">
                <h1>Crear una Función</h1>
              </v-col>
              <v-col cols="8">
                <v-card-text>
                  <v-autocomplete
                    v-model="peliculaModel"
                    :items="peliculas"
                    :loading="isLoading"
                    :search-input.sync="search"
                    hide-no-data
                    hide-selected
                    :rules="[v => !!v || 'Item requerido']"
                    item-text="nombre"
                    item-value="id"
                    label="Seleccionar Película"
                    placeholder="Escribe el nombre de la película"
                    prepend-icon="mdi-database-search"
                    return-object
                  ></v-autocomplete>
                </v-card-text>

                <v-row class="px-4">
                  <v-col cols="12" sm="6">
                    <v-menu transition="scale-transition" offset-y min-width="auto">
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field
                          :rules="[v => !!v || 'Campo requerido']"
                          v-model="Fecha"
                          label="Seleccionar fecha"
                          prepend-icon="mdi-calendar"
                          readonly
                          v-bind="attrs"
                          v-on="on"
                        ></v-text-field>
                      </template>
                      <v-date-picker v-model="Fecha" no-title scrollable lang="es">
                      </v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" sm="6">
                    <v-menu
                      ref="menu"
                      v-model="menuHora"
                      :close-on-content-click="false"
                      :nudge-right="40"
                      :return-value.sync="hora"
                      transition="scale-transition"
                      offset-y
                      max-width="290px"
                      min-width="290px"
                    >
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field
                          v-model="hora"
                          :rules="[v => !!v || 'Campo requerido']"
                          label="Seleccionar hora"
                          prepend-icon="mdi-clock-time-four-outline"
                          readonly
                          v-bind="attrs"
                          v-on="on"
                        ></v-text-field>
                      </template>
                      <v-time-picker
                        v-if="menuHora"
                        v-model="hora"
                        full-width
                        ampm-in-title
                        @click:minute="$refs.menu.save(hora)"
                      ></v-time-picker>
                    </v-menu>
                  </v-col>
                </v-row>
                <v-radio-group v-model="tipoFuncion" mandatory row class="px-4" dense>
                  <template v-slot:label>
                    <div>Tipo de Función</div>
                  </template>
                  <v-radio label="Normal" value="1"></v-radio>
                  <v-radio label="3D" value="2"></v-radio>
                </v-radio-group>
                <div class="mx-4">
                  <v-combobox
                    v-model="sala"
                    label="Seleccionar sala"
                    :rules="[v => !!v || 'Item requerido']"
                    :items="salas"
                    item-text="nombre"
                    item-value="id"
                  />
                </div>
                <v-row class="px-4">
                  <v-col cols="12" v-if="funciones.length != 0">
                    <span class="text--secondary">Otras funciones:</span>
                    <v-row justify="center" class="mx-0 my-0">
                      <v-col class="pb-0">
                        <v-chip v-for="horario in funciones" :key="horario" 
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
                  <v-col cols="12" class="mt-2">
                    <v-btn color="primary" block class="text-center" @click="crearFuncion">
                      Crear función
                    </v-btn>
                  </v-col>
                </v-row>
              </v-col>
              <v-col cols="4" v-if="peliculaObject">
                <pelicula :pelicula="peliculaObject" />
              </v-col>
              {{horaInicio}}
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
import { createFuncionAPI, getFuncionesByPeliculaId } from '../api/funcion'
import { getPeliculasIdAPI, getPeliculasAPI } from '../api/pelicula'
import { getSalasDisponibles } from '../api/sala'
import DialogError from '../components/DialogError.vue'
import Pelicula from '../components/Pelicula.vue'
import MainLayout from '../layout/MainLayout.vue'

export default {
  name : 'CreateFuncion',
  components: { MainLayout, DialogError, Pelicula },
  data: () => ({
    valid: false,
    peliculaModel: null,
    peliculaObject: null,
    peliculas: [],
    isLoading: false,
    search: null,
    menuFecha: null,
    Fecha: null,
    menuHora: null,
    hora: null,
    tipoFuncion: null,
    funciones: [],
    salas: [],
    sala: null,
    aux: 0,
  }),
  computed: {
    getSala: function(){
      if(this.peliculaModel && this.Fecha !== null && this.hora !==null) 
        this.updateAux()
      return this.aux
    },
    horaInicio: function(){
      moment.locale("es")
      return moment(`${this.Fecha} ${this.hora}`).format()
    }
  },
  methods:{
    ...mapMutations(["setDialogError", "setDialogErrorMessage"]),
    updateAux: function(){
      this.aux++
    },
    getPeliculaById: async function() {
      const res = await getPeliculasAPI(this.peliculaModel.nombre)
      this.peliculaObject = res.data.peliculas[0]
      getFuncionesByPeliculaId(this.peliculaModel.id).then(res => {
        this.funciones = res.data.funciones.map(f => (f.horario))
      })
    },
    crearFuncion(){
       if (!this.valid) {
        this.$refs.form.validate()
        return
      }
      createFuncionAPI(this.peliculaObject.id, this.sala.id, this.horaInicio, this.tipoFuncion)
      .then(() =>{
        this.$router.push({ path: '/funciones' })
      })
      .catch(() => {
        this.setDialogErrorMessage("No se pudo registrar exitosamente la película")
        this.setDialogError(true)
      })
    }
  },
  watch: {
    search () {
      if (this.isLoading) return
      this.isLoading = true
      getPeliculasIdAPI()
      .then(res => {
        this.peliculas = res.data.peliculas
      })
      .finally(() => (this.isLoading = false))
    },
    peliculaModel: function(){
      if (this.peliculaModel)
        this.getPeliculaById()
    },
    getSala: function(){
      getSalasDisponibles(this.horaInicio, this.peliculaModel.id).then(res => {
        this.salas = res.data.salas
      })
    },
  },
  created: function(){
    getPeliculasIdAPI()
    .then(res => {
      this.peliculas = res.data.peliculas
    })
  },
}
</script>
<style scoped>

</style>