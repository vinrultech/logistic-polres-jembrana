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
                v-model="nama"
                type="text"
                :rules="[(v) => !!v || 'Nama',
                (v) => v.length > 3 || 'Nama harus lebih dari tiga huruf']"
                label="Nama"
                required
              ></v-text-field>
              <v-text-field
                v-model="hp"
                type="text"
                label="No WA"
              ></v-text-field>

              <upload-foto
                class="mb-6"
                :image="foto"
                :username="username"
                @uploaded="uploaded"
                @removed="removed"
              />

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
import utils from "../../utils";
import UploadFoto from "../../components/upload_foto";
import Breadcum from "../../components/breadcum";
import { ACCOUNT } from "../../breadcum";
import { mapGetters } from "vuex";

export default {
  name: "account",
  components: {
    UploadFoto,
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumOne(ACCOUNT),
    valid: true,
    username: utils.username(),
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
    nama: {
      get() {
        return this.$store.getters["user/nama"];
      },
      set(value) {
        this.$store.dispatch("user/set_nama", value);
      },
    },
    hp: {
      get() {
        return this.$store.getters["user/hp"];
      },
      set(value) {
        this.$store.dispatch("user/set_hp", value);
      },
    },
    foto: {
      get() {
        return this.$store.getters["user/foto"];
      },
      set(value) {
        this.$store.dispatch("user/set_foto", value);
      },
    },
  },
  mounted() {
    this.$store.dispatch("user/fetch", {
      username: this.username,
      progress: this.$progress,
      axios: this.$axios,
      swal: this.$swal,
      router: this.$router,
    });
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("user/account", {
            item: {
              username: this.username,
              nama: this.nama,
              foto: this.foto,
              hp: this.hp,
            }
          });
      }
    },
    uploaded(foto) {
      this.foto = foto;
    },
    removed() {
      this.foto = "";
    },
  },
};
</script>

<style scoped>
</style>