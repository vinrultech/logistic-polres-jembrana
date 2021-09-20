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
                    v-model="kode"
                    :rules="[(v) => !!v || 'Kode dibutuhkan']"
                    label="Kode"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="nama"
                    :rules="[(v) => !!v || 'Nama dibutuhkan']"
                    label="Nama"
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
                <v-icon dark left>fas fa-save</v-icon>Create
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
import { KATEGORI, ADD } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "create_kategori",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(KATEGORI(false), ADD),
    kode: null,
    nama: null,
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
  },
  mounted() {
    this.$store.dispatch("constant/set_error", false);
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("kategori/create", {
          kode: this.kode,
          nama: this.nama,
        });
      }
    },
    batal() {
      this.$router.push("/admin/kategori");
    },
  },
};
</script>

<style scoped>
</style>