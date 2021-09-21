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
                    :rules="[
                      (v) => !!v || 'Username dibutuhkan',
                      (v) =>
                        !!/^[a-zA-Z0-9_]+$/.test(v) ||
                        'Username harus huruf dan angka',
                    ]"
                    label="Username"
                    required
                    disabled
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
                @click="update"
              >
                <v-icon dark left>fas fa-save</v-icon>Update
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Breadcum from "../../components/breadcum";
import { USER, EDIT } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
import { v1 as uuidv1 } from "uuid";
export default {
  name: "create_user",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(USER(false), EDIT),
    id: null,
    nama: null,
    username: null,
    role: null,
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
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        const unit_kerja = this.$store.getters["unit_kerja/find_all_item"](parseInt(this.unit_kerja_id));
        this.$store.dispatch("user/update", {
          item: {
            nama: this.nama,
            hp: this.hp,
            unit_kerja : unit_kerja,
            username: this.username,
            role: this.role,
            password: uuidv1()
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/user");
    },
  },
  mounted() {
    const id = this.$route.params.id;
    this.$store.dispatch("unit_kerja/all").then(() => {
      const item = this.$store.getters["user/item"](parseInt(id));
      if (item != null || item != undefined) {

        this.unit_kerja_id = item.unit_kerja.id;
        this.username = item.username;
        this.nama = item.nama;
        this.hp = item.hp;
        this.role = item.role;
        this.id = id;
      }
    });
  },
};
</script>

<style scoped>
</style>