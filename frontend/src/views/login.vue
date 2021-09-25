<template>
  <v-row align="center" justify="center">
    <v-col cols="12" sm="8" md="8">
      <v-card class="elevation-12">
        <v-row>
          <v-col cols="12" md="4" class="brown accent-3">
            <v-card-text>
              <div class="text-center">
                <img src="@/assets/logo.svg" :width="imgWidth"/>
                <h1 class="text-center text-h5 text-lg-h4 text-xl-h4 teal--text text--accent-3">
                  POLRES JEMBRANA
                </h1>
              </div>
            </v-card-text>
          </v-col>
          <v-col cols="12" md="8">
            <v-card-text class="mt-lg-12 mt-xl-12 mt-md-12 mt-xs-0">
              <h1 class="text-center text-h5 text-lg-h4 text-xl-h4 teal--text text--accent-3">
                Login Aplikasi Inventory dan Surat
              </h1>
              <h4 class="text-center mt-4">
                <v-alert v-model="error" dismissible type="error">
                  {{ errorMessage }}
                </v-alert>
                <div v-show="!error">
                  Pastikan username dan password dengan benar
                </div>
              </h4>
              <v-form>
                <v-text-field
                  label="Username"
                  name="Username"
                  prepend-icon="fas fa-user-lock"
                  type="text"
                  color="teal accent-3"
                  v-model="username"
                />

                <v-text-field
                  id="password"
                  label="Password"
                  name="password"
                  prepend-icon="fas fa-lock"
                  type="password"
                  color="teal accent-3"
                  v-model="password"
                />
              </v-form>
            </v-card-text>
            <div class="text-center mt-3">
              <v-btn rounded color="teal accent-3" dark @click="login"
                >SIGN IN</v-btn
              >
            </div>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
  <!-- <v-row align="center" justify="center">
    <v-col cols="12" sm="12" md="6" xs="12" xl="6">
      <v-row align="center" justify="center">
        <v-col cols="12">
          <div class="text-center title"><h2>Selamat Datang di</h2></div>
          <div class="text-center title"><h1>APLIKASI SURAT</h1></div>
          <div class="text-center">
            <img class="my-6 logo" src="@/assets/logo.svg" />
          </div>
          <div class="text-center title">LOGISTIK</div>
          <div class="text-center title">POLRES JEMBRANA - POLDA BALI</div>
        </v-col>
      </v-row>
    </v-col>
    <v-col cols="12" sm="12" md="6" xs="12" xl="6">
      <v-card class="elevation-12 rounded-xl">
        <v-card-text>
          <v-alert v-model="error" dismissible type="error">
            {{ errorMessage }}
          </v-alert>
          <v-form>
            <v-text-field
              label="Login"
              name="login"
              prepend-icon="fas fa-user-lock"
              type="text"
              v-model="username"
            ></v-text-field>
            <v-text-field
              id="password"
              label="Password"
              name="password"
              prepend-icon="fas fa-lock"
              type="password"
              v-model="password"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions class="pl-3 pr-5">
          <v-btn
            class="rounded-lg"
            x-large
            color="primary"
            @click="login"
            dark
            block
            elevation="5"
            style="width: 50%"
          >
            Sign In<v-icon dark right>fas fa-sign-in-alt</v-icon>
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row> -->
</template>

<script>
import CryptoJS from "crypto-js";
import { mapGetters } from "vuex";
export default {
  name: "login",
  data: () => ({
    username: null,
    password: null,
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
    imgWidth() {
      switch (this.$vuetify.breakpoint.name) {
        case "xs":
          return 100;
        case "sm":
          return 100;
        case "md":
          return 150;
        case "lg":
          return 150;
        case "xl":
          return 150;
        default:
          return 150;
      }
    },
  },
  mounted() {
    this.$store.dispatch("constant/remove_error");
  },
  methods: {
    login() {
      this.$store.dispatch("user/login", {
        username: this.username,
        password: CryptoJS.MD5(this.password).toString(),
      });
    },
  },
};
</script>

<style scoped>
.title {
  font-weight: bold;
  font-size: 24px;
  color: #eceff1;
}

.logo_polres {
  width: 150px;
  height: auto;
}

.logo {
  width: 80px;
  height: auto;
}
</style>