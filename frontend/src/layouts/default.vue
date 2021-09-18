<template>
  <v-app id="keep">
    <v-app-bar app color="blue-grey">
      <span class="title ml-3 mr-5">
        <v-avatar size="60">
          <v-img src="@/assets/logo.svg"></v-img>
        </v-avatar>
        Aplikasi&nbsp;
        <span class="font-weight-light">Berkas Tracker</span>
      </span>
      <v-spacer />
      
      <v-btn dark text to="/" class="mr-2"> <v-icon left>fas fa-home</v-icon>Home</v-btn>
      <v-btn dark text to="/berkas" class="mr-12"> <v-icon left>fas fa-id-card-alt</v-icon>Berkas</v-btn>
      

      <v-btn :dark=" is_login ? false : true" :to=" is_login ? '/admin' : '/login'" :color="is_login ? 'grey lighten-5' : 'success'">
        {{ is_login ? 'Admin' : 'Login'}}
        <v-icon right>{{ is_login ? 'fas fa-users-cog' : 'fas fa-user-lock'}}</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import utils from "../utils"
export default {
  props: {
    source: String,
  },
  data: (vm) => ({
    host: vm.$host,
  }),
  computed: {
    is_login() {
      return utils.session() !== null;
    }
  },
  methods: {},
  mounted() {},
};
</script>

<style>
#keep .v-navigation-drawer__border {
  display: none;
}

.company {
  height: 64px;
  width: 256px;
  position: fixed;
  background-color: #607d8b;
  color: white;
}
</style>