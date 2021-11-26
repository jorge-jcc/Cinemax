<template>
  <v-dialog v-model="showDialog" max-width="290">
      <v-card>
        <v-card-title class="text-h6">
          Ocurrio un error
        </v-card-title>
        <v-card-text>
          {{transactionErrorMessage}}
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red darken-1" text @click="showDialog = false">
            Aceptar
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script>
import { mapState, mapMutations } from "vuex"

export default {
  name: "DialogError",
  data: () => ({

  }),
  computed: {
    ...mapState(["transactionError", "transactionErrorMessage"]),
    showDialog: {
      get: function () {
        return this.transactionError
      },
      set: function () {
        this.setTransactionError(false)
      },
    },
  },
  methods:{
    ...mapMutations(["setTransactionError"]),
  }
}
</script>