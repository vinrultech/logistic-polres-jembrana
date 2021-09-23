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
                    v-model="no_surat"
                    :rules="[(v) => !!v || 'No Surat dibutuhkan']"
                    label="No Surat"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-dialog
                    ref="dialog"
                    v-model="modal"
                    :return-value.sync="tanggal_surat"
                    persistent
                    width="290px"
                  >
                    <template v-slot:activator="{ on, attrs }">
                      <v-text-field
                        v-model="tanggal_surat"
                        :rules="[(v) => !!v || 'Tanggal Surat dibutuhkan']"
                        label="Tanggal Surat"
                        required
                        readonly
                        v-bind="attrs"
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker v-model="tanggal_surat" scrollable>
                      <v-spacer></v-spacer>
                      <v-btn text color="primary" @click="modal = false">
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                        @click="$refs.dialog.save(tanggal_surat)"
                      >
                        OK
                      </v-btn>
                    </v-date-picker>
                  </v-dialog>
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
              <v-row>
                <v-col cols="12">
                  <multiple-upload :status="'add'" class="mb-6" />
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
import MultipleUpload from "../../components/multiple_upload.vue";
import { SURAT_KELUAR, EDIT } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "edit_surat_keluar",
  components: {
    Breadcum,
    MultipleUpload,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(SURAT_KELUAR(false), EDIT),
    id: null,
    no_surat: null,
    tanggal_surat: null,
    tujuan: null,
    perihal: null,
    isi: null,
    unit_kerja: null,
    valid: true,
    modal: false,
    row_id: null,
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
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("surat_keluar/update", {
          item: {
            row_id: this.row_id,
            no_surat: this.no_surat,
            tanggal_surat: this.tanggal_surat,
            tujuan: this.tujuan,
            perihal: this.perihal,
            isi: this.isi,
            unit_kerja: this.unit_kerja,
            files: this.$store.getters["files/files"],
          },
          row_id: this.row_id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/surat_keluar");
    },
  },
  mounted() {
    this.$store.dispatch("constant/set_error", false);
    const row_id = this.$route.params.row_id;
    const item = this.$store.getters["surat_keluar/item"](row_id);
    if (item != null || item != undefined) {
      (this.no_surat = item.no_surat),
        (this.tanggal_surat = item.tanggal_surat),
        (this.tujuan = item.tujuan),
        (this.perihal = item.perihal),
        (this.isi = item.isi),
        (this.row_id = row_id),
        (this.unit_kerja = item.unit_kerja),
        (this.id = item.id);
      this.$store.commit("files/SET", item.files);
    }
  },
};
</script>

<style scoped>
</style>