<template>
  <v-app-bar dark app color="red darken-4" flat>
    <v-avatar size="40">
      <v-img src="@/assets/Imagen3.svg"></v-img>
    </v-avatar>
    <h2 class="noselect" style="position:absolute; left: 62px">Cinemax</h2>
    <v-tabs centered class="ml-n9">
      <v-tab v-for="link in links" :key="link">
        {{ link }}
      </v-tab>
    </v-tabs>
    <router-link to="" style="position:absolute; right: 16px">
      <v-btn dark class="ma-2" elevation="2" outlined  @click.native="userLogout">
        Cerrar Sesión
      </v-btn>
    </router-link>
  </v-app-bar>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "AppBar",
  data: () => ({
    links: [
      'Taquilla',
      'Funciones',
      'Peliculas',
    ],
  }),
  methods:{
    ...mapActions(["userLogout"]),
  }
}
</script>