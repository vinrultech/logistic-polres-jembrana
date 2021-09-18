<template>
  <v-app id="keep">
    <!-- <v-app-bar app color="green lighten-1">
      
      <span class="title ml-3 mr-5">
        Aplikasi&nbsp;
        <span class="font-weight-light">Pelayanan Desa</span>
      </span>
      <v-spacer />

    </v-app-bar> -->
    

    

    <v-main>
      <v-container fluid>
        <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

import utils from "../utils";


export default {
  props: {
    source: String,
  },
  data: (vm) => ({
    
    items: [
      { icon: "fas fa-home", text: "Dashboard", href: "/admin" },
      { icon: "fas fa-users", text: "User", href: "/admin/user" },
    ],
    host: vm.$host,
  }),
  computed: {
    
  },
  methods: {
    signOut() {
      utils.removeSession();
      this.$router.push("/login");
    },
    to(href) {
      //console.log(href);
      this.$router.push(href);
    },
  },
  mounted() {
    this.$store.dispatch("user/nama");
  },
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