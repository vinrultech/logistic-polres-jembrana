<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12">
        <v-alert v-model="error" dismissible type="error">
          {{ errorMessage }}
        </v-alert>
        <v-card class="elevation-6">
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
                    :rules="[(v) => !!v || 'Username dibutuhkan',
                            (v) => !!/^[a-zA-Z0-9_]+$/.test(v) || 'Username harus huruf dan angka',]"
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
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-select
                    v-model="unit_kerja_id"
                    :items="unit_kerjas"
                    menu-props="auto"
                    item-text="nama"
                    item-value="id"
                    label="Unit Kerja"
                    hide-details
                    prepend-icon="fas fa-building"
                    single-line
                    :rules="[(v) => !!v || 'Role dibutuhkan']"
                    required
                  ></v-select>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="hp"
                    :rules="[(v) => !!v || 'No WA dibutuhkan']"
                    label="No WA"
                    required
                  ></v-text-field>
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
import { USER, ADD } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "create_user",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(USER(false), ADD),
    nama: null,
    username: null,
    password: null,
    confirm_password: null,
    hp: null,
    unit_kerja_id: null,
    valid: true,
  }),
  computed: {
    ...mapGetters({
      unit_kerjas: "unit_kerja/all_items",
      errorMessage: "constant/errorMessage",
    }),
    passwordConfirmationRule() {
      return () =>
        this.password === this.confirm_password || "Password tidak sama";
    },
    error: {
      get() {
        return this.$store.getters["constant/error"];
      },
      set(val) {
        this.$store.dispatch("constant/set_error", val);
      },
    },
  },
  mounted() {
    this.$store.dispatch("unit_kerja/all");
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        const unit_kerja = this.$store.getters["unit_kerja/find_all_item"](parseInt(this.unit_kerja_id));
        this.$store.dispatch("user/create",{
            username: this.username,
            nama: this.nama,
            hp: this.hp,
            unit_kerja: unit_kerja,
            role: "unit_kerja",
            password: CryptoJS.MD5(this.password).toString(),
          },);
      }
    },
    batal() {
      this.$router.push("/admin/user");
    }
  },
};
</script>

<style scoped>
</style>