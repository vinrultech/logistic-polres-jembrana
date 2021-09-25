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
                    :rules="[(v) => !!v || 'Kode Aset Tetap dibutuhkan']"
                    label="Kode Aset Tetap"
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

              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-select
                    v-model="kategori_id"
                    :items="kategoris"
                    menu-props="auto"
                    item-text="nama"
                    item-value="id"
                    label="Kategori"
                    hide-details
                    prepend-icon="fas fa-cookie-bite"
                    single-line
                    :rules="[(v) => !!v || 'Kategori dibutuhkan']"
                    required
                  ></v-select>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="jumlah"
                    :rules="[(v) => !!v || 'Jumlah dibutuhkan']"
                    label="Jumlah"
                    type="number"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>

              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-select
                    v-model="metric_id"
                    :items="metrics"
                    menu-props="auto"
                    item-text="nama"
                    item-value="id"
                    label="Satuan Metrik"
                    hide-details
                    prepend-icon="fas fa-sort-numeric-down"
                    single-line
                    :rules="[(v) => !!v || 'Satuan Metrik dibutuhkan']"
                    required
                  ></v-select>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <span>Status Aset</span>
                  <v-radio-group v-model="status" row>
                    <v-radio label="Baik" value="1"></v-radio>
                    <v-radio label="Rusak" value="2"></v-radio>
                  </v-radio-group>
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
import { mapGetters } from "vuex";
export default {
  name: "edit_metric",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(ASSET_TETAP(false), EDIT),
    row_id: null,
    nama: null,
    kode: null,
    jumlah: null,
    kategori_id: null,
    metric_id: null,
    unit_kerja: null,
    status: "1",
    valid: true,
  }),
  computed: {
    ...mapGetters({
      errorMessage: "constant/errorMessage",
      kategoris: "kategori/all_items",
      metrics: "metric/all_items",
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
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        const kategori = this.$store.getters["kategori/find_all_item"](
          parseInt(this.kategori_id)
        );
        const metric = this.$store.getters["metric/find_all_item"](
          parseInt(this.metric_id)
        );
        this.$store.dispatch("aset_tetap/update", {
          item: {
            nama: this.nama,
            kode: this.kode,
            jumlah: parseInt(this.jumlah),
            status: parseInt(this.status),
            kategori: kategori,
            metric: metric,
            unit_kerja: this.unit_kerja,
          },
          row_id: this.row_id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/aset_tetap");
    },
  },
  mounted() {
    this.$store.dispatch("constant/set_error", false);
    const row_id = this.$route.params.row_id;
    Promise.all([
      this.$store.dispatch("kategori/all"),
      this.$store.dispatch("metric/all"),
    ]).then(() => {
      const item = this.$store.getters["aset_tetap/item"](row_id);
      if (item != null || item != undefined) {
        this.nama = item.nama;
        this.kode = item.kode;
        this.status = `${item.status}`;
        this.jumlah = parseInt(item.jumlah);
        this.kategori_id = item.kategori.id;
        this.metric_id = item.metric.id;
        this.unit_kerja = item.unit_kerja;
        this.row_id = row_id;
      }
    });
  },
};
</script>

<style scoped>
</style>