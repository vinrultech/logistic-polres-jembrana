<template>
  <v-row align="center" justify="center">
    <v-col cols="12" >
      <h1>Not Found 404</h1>
    </v-col>
  </v-row>
</template>

<script>
import CryptoJS from "crypto-js";

export default {
  name: "login",
  data: () => ({
    username: null,
    password: null,
    valid: true
  }),
  methods: {
    login() {
      this.$store.dispatch("user/login", {
        username: this.username,
        password: CryptoJS.MD5(this.password).toString(),
        vm: this
      });
    }
  }
};
</script>