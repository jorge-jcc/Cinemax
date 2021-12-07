<template>
  <main-layout>
    <v-row justify="center">
      <v-col cols="12" sm="8" >
        <v-sheet min-height="70vh" rounded="lg">
          <v-row class="px-5 pb-2 mt-0">
            <v-col cols="12" class="text-center">
              <h1 class="display-3 font-weight-bold noselect">
                Cinemax
              </h1>
            </v-col>
            <v-col cols="12" md="7">
              <combobox title="Película"
                label="Seleccione la pelicula" 
                :items="peliculas" 
                name="nombre"
                @select="getFunciones">
              </combobox>
            </v-col>
            <v-col cols="12" md="5">
              <combobox title="Función"
                label="Seleccione la función" 
                :items="funciones" 
                name="horario"
                @select="getAsientos">
              </combobox>
            </v-col>
            <div v-if="sala">
              <v-row justify="space-between" align="center" class="mb-2 mx-1">
                <v-col cols="auto">
                  <h3 class="my-0">{{sala.nombre}}</h3>
                  <span class="text-caption my-0">{{sala.ubicacion}}</span>
                </v-col>
                <v-col cols="auto">
                  <v-btn class="ma-2" x-large text icon color="grey darken-2" @click="refreshAsientos">
                    <v-icon>mdi-refresh</v-icon>
                  </v-btn>
                </v-col>
              </v-row>
              <v-img src="@/assets/pantalla.svg" class="mb-5"/>
              <v-row v-if="asientosLocal!=null" justify="center" class="mb-5">
                <div class="grid-container mx-5">
                  <asiento v-for="(asiento, index) in asientosLocal" :key="index" 
                  :asiento="asiento" 
                  v-on:update="refreshAsientos"
                />
                </div>
              </v-row>
            </div>
            <dialog-error></dialog-error>
          </v-row>
        </v-sheet>
      </v-col>
      <v-col cols="12" sm="4">
        <div style="position: sticky; top: 4rem;">
          <v-sheet rounded="lg" min-height="268" class="">
            <ticket :pelicula="peliculaNombre" :horario="peliculaHorario" :sala="peliculaSala"/>
          </v-sheet>
        </div>
      </v-col>
    </v-row>
  </main-layout>
</template>

<script>
import { mapActions, mapState } from "vuex"
import Combobox from '../components/Combobox.vue'
import { getPeliculasEnCartelera } from "../api/pelicula"
import { getFuncionesByPeliculaId } from "../api/funcion"
import { getInformacionSala } from "../api/sala"
import { getAsientosByFuncion } from "../api/asientos"
import Asiento from '../components/Asiento.vue'
import DialogError from '../components/DialogError.vue'
import Ticket from "../components/Ticket.vue"
import MainLayout from '../layout/MainLayout.vue'

export default {
  name : 'Home',
  components: { Combobox, Asiento, DialogError, Ticket, MainLayout },
  data: () => ({
    peliculas: [],
    peliculaID: null,
    funciones: [],
    funcionId: null,
    sala: null,
    asientosLocal: null,
    dialog: true,
  }),
  methods:{
    ...mapActions(["deshacerTransaccionAPI", "getPreciosBoletos"]),
    getFunciones: async function(peliculaID){
      await this.reset()
      this.peliculaID = this.peliculas[peliculaID].id;
      getFuncionesByPeliculaId(this.peliculaID).then(res => {
        this.funciones = res.data.funciones
      }).catch(() => {})
    },
    getAsientos: async function(funcionId){
      await this.reset()
      this.funcionId = this.funciones[funcionId].id;
      getInformacionSala(this.funcionId).then(res =>{
        if (res.status == 200){
          this.sala = res.data.sala;
        }
      })
      this.refreshAsientos()
    },
    refreshAsientos: function(){
      getAsientosByFuncion(this.funcionId).then(res =>{
        const a = res.data.asientos.map((a, i) => ({...a, "i": i}))
        this.asientos.forEach((e) => a[e.i].StatusAsiento === "EN PROCESO" && e.id === a[e.i].id ? a[e.i] = e: null);
        this.asientosLocal = a
      })
    },
    async reset(){  
      await this.deshacerTransaccionAPI()
      this.funcionId = null
      this.asientosLocal = null
      this.sala = null
    },
  },
  computed: {
    ...mapState(["asientos", "updateAsientos"]),
    peliculaNombre: function(){
      return this.peliculas?.find(p => p.id == this.peliculaID)?.nombre
    },
    peliculaHorario: function(){
      return this.funciones?.find(f => f.id == this.funcionId)?.horario.split(" ")[0]
    },
    peliculaSala: function(){
      return this.sala?.nombre
    },
  },
  mounted() {
    getPeliculasEnCartelera().then(res =>{
      if (res.status == 200){
        this.peliculas = res.data.peliculas
      }
    }).catch(() => {})
  },
  watch:{
    updateAsientos: function(){
      this.refreshAsientos()
    }
  },
  created: function(){
    this.getPreciosBoletos()
  },
}
</script>
<style scoped>
.grid-container{
  display: grid;
  grid-template-columns: repeat(15, 1fr);
}

</style>