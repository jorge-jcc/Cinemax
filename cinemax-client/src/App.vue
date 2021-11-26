<template>
  <v-app>
    <div v-if="user">
      <router-view />
    </div>
    <div v-else>
      <login/>
    </div>
  </v-app>
</template>

<script>
import { mapActions, mapState } from "vuex";
import Login from './views/Login.vue';

export default {
  name: 'App',
  components: { Login },
  data: () => ({

  }),
  computed: {
    ...mapState(["user"]),
  },
  methods: {
    ...mapActions(["userLogin"]),
  },
  mounted() {
    this.userLogin();
  },
};
</script>

<style>
.noselect {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}
</style>