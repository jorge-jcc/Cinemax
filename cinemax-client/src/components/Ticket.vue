<template>
  <v-layout class="noselect">
    <v-row class="pa-2">
      <v-col>
        <h2 class="text-center">Ticket de Venta</h2>
        <v-col class="text-end py-0"><span class="text-caption">{{time}}</span></v-col>
        <v-col class="py-0"><span class="text-body-2">Película:</span></v-col>
        <v-col class="py-0"><span class="text-body-2">Horario:</span></v-col>
        <v-col class="py-0"><span class="text-body-2">Sala:</span></v-col>
        <v-data-table hide-default-footer dense :headers="headers" :items="boletos">
          <!-- eslint-disable-next-line -->
          <template v-slot:item.tipoBoleto="props" >
            <v-edit-dialog :return-value.sync="props.item.name">
                <v-select
                  :items="items"
                  dense
                  v-model="props.item.tipoBoleto"
                ></v-select>
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
        <v-col class="py-0"><span class="text-caption">Fue atendido por: {{user.email}}</span></v-col>
      </v-col>
    </v-row>
  </v-layout>
</template>

<script>
import moment from "moment"
import { mapActions, mapState } from "vuex";

export default {
  name: "Ticket",
  data: () => ({
    items: ['Normal', 'Calibri', 'Courier', 'Verdana'],
    headers: [
      { text: '#', value: 'num', sortable: false, },
      { text: 'Asiento', value: 'asiento', sortable: false, },
      { text: 'Tipo Boleto', value: 'tipoBoleto', sortable: false, },
      { text: 'Precio', value: 'precio',sortable: false, }
    ],
  }),
  computed: {
    ...mapState(["asientos", "user"]),
    boletos: function(){
      return this.asientos.map((a, i) => ({
        num: i,
        asiento: a.clave,
        tipoBoleto: "Normal",
        precio: 50
      }))
    },
    total: function(){
      let t = 0
      this.boletos.forEach(e => { t += e.precio });
      return t
    },
    time: function(){
      return moment().format('lll');
    }
  },
  methods:{
    ...mapActions(["userLogout"]),
    
  }
}
</script>