<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12">
        <v-card class="elevation-6">
          <v-toolbar>
            <v-toolbar-title>Tambah User</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form" v-model="valid">
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="nama"
                    :rules="[(v) => !!v || 'Nama dibutuhkan']"
                    label="Nama"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="username"
                    :rules="[(v) => !!v || 'Username dibutuhkan']"
                    label="Username"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="password"
                    type="password"
                    :rules="[(v) => !!v || 'Password dibutuhkan']"
                    label="Password"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
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
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-select
                    v-model="role"
                    :items="roles"
                    menu-props="auto"
                    item-text="nama"
                    item-value="id"
                    label="Role"
                    hide-details
                    prepend-icon="fas fa-users-cog"
                    single-line
                    :rules="[(v) => !!v || 'Role dibutuhkan']"
                    required
                  ></v-select>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="hp"
                    :rules="[(v) => !!v || 'No WA dibutuhkan']"
                    label="No WA"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <upload-foto
                    :image="foto"
                    @uploaded="uploaded"
                    @removed="removed"
                  />
                </v-col>
              </v-row>
              <v-btn color="error" class="mr-4" @click="batal">
                <v-icon dark left>fas fa-arrow-left</v-icon>Batal
              </v-btn>
              <v-btn
                :disabled="!valid"
                color="primary"
                class="mr-4"
                @click="simpan"
              >
                <v-icon dark left>fas fa-save</v-icon>Simpan
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
import Breadcum from "../../components/breadcum";
import UploadFoto from "../../components/upload_foto";
import { USER, ADD } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "create_user",
  components: {
    Breadcum,
    UploadFoto,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(USER(false), ADD),
    roles: [
      {
        id: "admin",
        nama: "Admin",
      },
      {
        id: "petugas_ukur",
        nama: "Petugas Ukur",
      },
      {
        id: "petugas_gambar",
        nama: "Petugas Gambar",
      },
    ],
    nama: null,
    username: null,
    password: null,
    confirm_password: null,
    hp: null,
    foto: "",
    role: null,
    valid: true,
  }),
  computed: {
    passwordConfirmationRule() {
      return () =>
        this.password === this.confirm_password || "Password tidak sama";
    },
  },
  mounted() {},
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("user/add", {
          item: {
            username: this.username,
            nama: this.nama,
            hp: this.hp,
            foto: this.foto,
            role: this.role,
            password: CryptoJS.MD5(this.password).toString(),
          },
          vm: this,
        });
      }
    },
    batal() {
      this.$router.push("/admin/user");
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