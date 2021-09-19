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
import { KATEGORI, EDIT } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "edit_kategori",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(KATEGORI(false), EDIT),
    id: null,
    nama: null,
    kode: null,
    valid: true,
  }),
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("kategori/update", {
          item: {
            kode: this.kode,
            nama: this.nama,
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/kategori");
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["kategori/item"](parseInt(id));
    if (item != null || item != undefined) {
      this.kode = item.kode;
      this.nama = item.nama;
      this.id = id;
    }
  },
};
</script>

<style scoped>
</style>