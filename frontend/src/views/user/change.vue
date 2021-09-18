<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="6" lg="6" xl="12" sm="12" cols="6">
        <v-alert v-model="error" dismissible type="error">
          {{ errorMessage }}
        </v-alert>
        <v-card class="elevation-6 rounded-lg">
          <v-card-text>
            <v-form ref="form" v-model="valid">
              <v-text-field
                v-model="username"
                :rules="[(v) => !!v || 'Username is required']"
                label="Username"
                :disabled="true"
                required
              ></v-text-field>
              <v-text-field
                v-model="old_password"
                type="password"
                :rules="[(v) => !!v || 'Password lama dibutuhkan']"
                label="Old Password"
                required
              ></v-text-field>
              <v-text-field
                v-model="password"
                type="password"
                :rules="[(v) => !!v || 'Password dibutuhkan']"
                label="Password"
                required
              ></v-text-field>
              <v-text-field
                v-model="confirm_password"
                type="password"
                :rules="[
                  (v) => !!v || 'Konfirmasi password dibutuhkan',
                  passwordConfirmationRule,
                ]"
                label="Konfirmasi Password"
                required
              ></v-text-field>
              <v-btn
                :disabled="!valid"
                color="primary"
                class="mr-4 rounded-lg"
                @click="simpan"
              >
                <v-icon dark left>fas fa-save</v-icon>Ubah
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import CryptoJS from "crypto-js";
import utils from "../../utils";
import Breadcum from "../../components/breadcum";
import { PASSWORD } from "../../breadcum";
import { mapGetters } from "vuex";
export default {
  name: "change",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumOne(PASSWORD),
    username: null,
    password: null,
    confirm_password: null,
    old_password: null,
    valid: true,
  }),
  computed: {
    ...mapGetters({
      errorMessage: "constant/errorMessage",
    }),
    error: {
      get() {
        return this.$store.getters["constant/error"];
      },
      set(val) {
        this.$store.dispatch("constant/set_error", val);
      },
    },
    passwordConfirmationRule() {
      return () =>
        this.password === this.confirm_password || "Password tidak sama";
    },
  },
  mounted() {
    this.username = utils.username();
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store
          .dispatch("user/change", {
            item: {
              username: this.username,
              password: CryptoJS.MD5(this.password).toString(),
              confirm_password: CryptoJS.MD5(this.confirm_password).toString(),
              old_password: CryptoJS.MD5(this.old_password).toString(),
            }
          });
      }
    },
  },
};
</script>

<style scoped>
</style>