<template>
  <v-layout class="noselect">
    <v-row class="pa-2">
      <v-col>
        <h2 class="text-center">Ticket de Venta</h2>
        <v-col class="text-end py-0"><span class="text-body-2">{{ago}}</span></v-col>
        <v-col class="py-0">
          <span class="text--secondary text-caption">Película 
            <span class=" font-weight-bold">{{pelicula}}</span>
          </span>
        </v-col>
        <v-col class="py-0">
          <span class="text--secondary text-caption">Horario:
            <span class=" font-weight-bold">{{horario}}</span>
          </span>
        </v-col>
        <v-col class="py-0">
          <span class="text--secondary text-caption">Sala:
            <span class=" font-weight-bold">{{sala}}</span>
          </span>
        </v-col>
        <v-data-table hide-default-footer dense :headers="headers" :items="boletos">
          <!-- eslint-disable-next-line -->
          <template v-slot:item.tipoBoleto="props">
            <v-edit-dialog :return-value.sync="props.item.name">
                <v-select :items="tipoBoletos" dense v-model="props.item.tipoBoleto" @change="update"></v-select>
            </v-edit-dialog>
          </template>
          <!-- eslint-disable-next-line -->
          <template v-slot:body.append="{headers}">
            <tr class="font-weight-bold">
              <td v-for="(header, i) in headers" :key="i">
                <div v-if="header.value == 'tipoBoleto'">Total</div>
                <div v-if="header.value == 'precio'">{{total}}</div>
              </td>
            </tr>
          </template>
        </v-data-table>
        <v-row justify="center" class="text-center my-1">
          <v-col cols="12">
            <v-btn color="grey darken-1" rounded
              :disabled="boletos.length==0" 
              :dark="boletos.length!=0" 
              @click="iniciarCompra"
            >
              Confirmar compra
            </v-btn>
          </v-col>
        </v-row>
        <v-col class="py-0"><span class="text-caption">Fue atendido por: {{user.email}}</span></v-col>
      </v-col>
    </v-row>
    <compra :dialog="dialogCompra" :total="total" :boletos="boletos" @cerrar="dialogCompra = false"/>
  </v-layout>
</template>

<script>
import moment from "moment"
import { mapActions, mapState } from "vuex"
import { iniciarCompraAPI } from "../api/ticket"
import Compra from './Compra.vue'

export default {
  components: { Compra },
  props: ["pelicula", "horario", "sala"],
  name: "Ticket",
  data: () => ({
    boletos: [],
    headers: [
      { text: '#', value: 'num', sortable: false, },
      { text: 'Asiento', value: 'asiento', sortable: false, },
      { text: 'Tipo Boleto', value: 'tipoBoleto', sortable: false, },
      { text: 'Precio', value: 'precio',sortable: false, }
    ],
    dialogCompra: false,
    ago: null,
  }),
  created: function(){
    this.interval = setInterval(() =>{
      this.ago = moment().format("lll")
    }, 1000)
  },
  computed: {
    ...mapState(["asientos", "user", "preciosBoletos", "transactionId"]),
    tipoBoletos: function(){
      return this.preciosBoletos.map((p) => p.clave)
    },
    total: function(){
      let t = 0
      this.boletos.forEach(e => { t += e.precio });
      return t
    },
    time: function(){
      return moment().format('lll');
    },
  },
  watch:{
    asientos: function(newVal){
      this.boletos = newVal.map((a, i) => ({
        num: i,
        id: a.id,
        asiento: a.clave,
        tipoBoleto: "TRADICIONAL",
        precio: (this.preciosBoletos.find(e => e.clave == "TRADICIONAL")).precio
      }))
    },
    
  },
  methods:{
    ...mapActions(["deshacerTransaccionAPI"]),
    update: function(){
      this.boletos = this.boletos.map((b) => ({
        ...b,
        precio: (this.preciosBoletos.find(e => e.clave == b.tipoBoleto)).precio
      }))
    },
    iniciarCompra: async function(){
      let compra = {
        transaccionId: this.transactionId,
        boletos: this.boletos.map((b) => b.id)
      }
      await iniciarCompraAPI(compra).then(() => {})
      this.dialogCompra = true
    }
  }
}
</script>