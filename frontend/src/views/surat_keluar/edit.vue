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
                    v-model="no_surat"
                    :rules="[(v) => !!v || 'No Surat dibutuhkan']"
                    label="No Surat"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="tanggal_surat"
                    :rules="[(v) => !!v || 'Tanggal Surat dibutuhkan']"
                    label="Tanggal Surat"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="tujuan"
                    :rules="[(v) => !!v || 'Tujuan dibutuhkan']"
                    label="Tujuan"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="perihal"
                    :rules="[(v) => !!v || 'Perihal Surat dibutuhkan']"
                    label="Perihal"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-textarea
                    v-model="isi"
                    :rules="[(v) => !!v || 'Isi dibutuhkan']"
                    label="Isi Surat"
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
import { SURAT_KELUAR, EDIT } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "edit_surat_keluar",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(SURAT_KELUAR(false), EDIT),
    id: null,
    no_surat: null,
    tanggal_surat: null,
    tujuan: null,
    perihal: null,
    isi: null,
    valid: true,
  }),
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("surat_keluar/update", {
          item: {
            kode: this.kode,
            nama: this.nama,
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/surat_keluar");
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["surat_keluar/item"](parseInt(id));
    if (item != null || item != undefined) {
      (this.no_surat = item.no_surat),
        (this.tanggal_surat = item.tanggal_surat),
        (this.tujuan = item.tujuan),
        (this.perihal = item.perihal),
        (this.isi = item.isi),
        (this.id = id);
    }
  },
};
</script>

<style scoped>
</style>