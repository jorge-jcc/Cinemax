<template>
  <v-app id="inspire">
    <app-bar></app-bar>
    <v-main class="grey lighten-3">
      <v-container>
        <slot></slot>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import AppBar from '../components/AppBar.vue';

export default {
  name: "MainLayout",
  components: { AppBar },
    data: () => ({
  }),
}
</script>