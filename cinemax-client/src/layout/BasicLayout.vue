<template>
  <v-app id="inspire">
    <app-bar></app-bar>
    <v-main class="grey lighten-3">
      <v-container >
        <v-row justify="center">
          <v-col cols="12" sm="8" >
            <v-sheet min-height="70vh" rounded="lg">
              <slot></slot>
            </v-sheet>
          </v-col>
           <v-col cols="12" sm="4">
            <div style="position: sticky; top: 4rem;">
              <v-sheet rounded="lg" min-height="268" class="">
              <slot name="right"></slot>
            </v-sheet>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import AppBar from '../components/AppBar.vue';

export default {
  name: "BasicLayout",
  components: { AppBar },
    data: () => ({
  }),
}
</script>