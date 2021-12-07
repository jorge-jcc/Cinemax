<template>
  <v-dialog v-model="dialog" persistent width="35rem">
    <v-card>
      <v-card-title>
        <span class="text-h5">Realizar Pago</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <span class="text-subtitle-1 font-weight-bold">Tipo de compra</span>
          <v-radio-group v-model="metodoPago" mandatory dense class="mt-0 mb-4">
            <v-radio label="Efectivo" value="efectivo"></v-radio>
            <v-radio label="Tarjeta de Crédito/Débito" value="tarjeta"></v-radio>
          </v-radio-group>
          <div v-if="metodoPago =='efectivo'">
            <v-row justify="center">
              <v-col cols="8" class="py-0">
                <v-row class="my-0 py-0" align="center">
                  <v-col cols="6" class="my-0 py-0 text-end">
                    <p class="text-subtitle-2">Total a Pagar:</p>
                  </v-col>
                  <v-col cols="3" class="my-0 py-0">
                    <v-text-field class="my-0" :value="total" dense prefix="$" disabled></v-text-field>
                  </v-col>

                  <v-col cols="6" class="my-0 py-0 text-end">
                    <p class="text-subtitle-2">Paga con:</p>
                  </v-col>
                  <v-col cols="3" class="my-0 py-0">
                    <v-text-field class="my-0" v-model="paga" dense prefix="$"></v-text-field>
                  </v-col>

                  <v-col cols="6" class="my-0 py-0 text-end">
                    <p class="text-subtitle-2">Cambio:</p>
                  </v-col>
                  <v-col cols="3" class="my-0 py-0">
                    <v-text-field :value="paga-total" dense prefix="$" disabled></v-text-field>
                  </v-col>
                </v-row>
              </v-col>  
            </v-row>
          </div>
          <div v-else class="text-center">
            <v-progress-circular indeterminate color="primary"></v-progress-circular>
          </div>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="$emit('cerrar')">
          Cancelar
        </v-btn>
        <v-btn color="red darken-1" text @click="pagar" :disabled="(paga-total)<0">
          Pagar
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapActions, mapState } from 'vuex'
export default {
  props: ["dialog", "total", "boletos"],
  data () {
    return {
      metodoPago: null,
      paga: null
    }
  },
  computed:{
    ...mapState(["preciosBoletos"]),
  },
  methods:{
    ...mapActions(["crearTicket"]),
    pagar: async function(){
      await this.crearTicket({
        total: this.total, 
        boletos: this.boletos.map((b) => ({
          asientoId: b.id,
          tipoBoletoId: (this.preciosBoletos.find(e => e.clave == b.tipoBoleto)).id
        }))
      })
      this.paga = null
      this.$emit('cerrar')
    },
  },
}
</script>