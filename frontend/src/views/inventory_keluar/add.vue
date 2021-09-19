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
                    label="Nama Barang"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="tanggal"
                    :rules="[(v) => !!v || 'Tanggal dibutuhkan']"
                    label="Tanggal"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="jumlah"
                    :rules="[(v) => !!v || 'Jumlah dibutuhkan']"
                    label="Jumlah"
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
import Breadcum from "../../components/breadcum";
import { INVENTORY_KELUAR, ADD } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "create_inventory_keluar",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(INVENTORY_KELUAR(false), ADD),
    nama: null,
    tanggal: null,
    jumlah: null,
    valid: true,
  }),
  mounted() {},
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("inventory_keluar/add", {
          item: {
            kode: this.kode,
            nama: this.nama,
          }
        });
      }
    },
    batal() {
      this.$router.push("/admin/inventory_keluar");
    }
  },
};
</script>

<style scoped>
</style>