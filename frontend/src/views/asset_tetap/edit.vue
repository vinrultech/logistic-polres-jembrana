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
                    :rules="[(v) => !!v || 'Kode Barang dibutuhkan']"
                    label="Kode Barang"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="nama"
                    :rules="[(v) => !!v || 'Nama Barang dibutuhkan']"
                    label="Nama Barang"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="asset_tetap"
                    :rules="[(v) => !!v || 'asset_tetap dibutuhkan']"
                    label="asset_tetap"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="jumlah"
                    :rules="[(v) => !!v || 'Jumlah dibutuhkan']"
                    label="Jumlah"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
               <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="status"
                    :rules="[(v) => !!v || 'Status dibutuhkan']"
                    label="Status"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-textarea
                    v-model="spesifikasi"
                    :rules="[(v) => !!v || 'Spesifikasi dibutuhkan']"
                    label="Spesifikasi"
                    required
                  ></v-textarea>
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
import { ASSET_TETAP, EDIT } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "edit_asset_tetap",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(ASSET_TETAP(false), EDIT),
    id: null,
    kode: null,
    nama: null,
    asset_tetap: null,
    jumlah: null,
    status: null,
    spesifikasi: null,
    valid: true,
  }),
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("asset_tetap/update", {
          item: {
            kode: this.kode,
            nama: this.nama,
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/asset_tetap");
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["asset_tetap/item"](parseInt(id));
    if (item != null || item != undefined) {
      this.kode = item.kode;
      this.nama = item.nama;
      this.jumlah = item.jumlah;
      this.kategori = item.kategori;
      this.spesifikasi = item.spesifikasi;
      this.status = item.status;
      this.id = id;
    }
  },
};
</script>

<style scoped>
</style>