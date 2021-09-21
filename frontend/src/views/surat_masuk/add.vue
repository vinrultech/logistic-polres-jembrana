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
                  <!-- <v-text-field
                    v-model="tanggal_surat"
                    :rules="[(v) => !!v || 'Tanggal Surat dibutuhkan']"
                    label="Tanggal Surat"
                    required
                  ></v-text-field> -->
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
                    v-model="dari"
                    :rules="[(v) => !!v || 'Dari dibutuhkan']"
                    label="Dari"
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
import MultipleUpload from "../../components/multiple_upload.vue";
import { SURAT_MASUK, ADD } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "create_surat_masuk",
  components: {
    Breadcum,
    MultipleUpload,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(SURAT_MASUK(false), ADD),
    no_surat: null,
    tanggal_surat: null,
    dari: null,
    perihal: null,
    isi: null,
    valid: true,
    modal: false,
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
        this.$store.dispatch("surat_masuk/create", {
          no_surat: this.no_surat,
          tanggal_surat: this.tanggal_surat,
          dari: this.dari,
          perihal: this.perihal,
          isi: this.isi,
          files: this.$store.getters["files/files"],
        });
      }
    },
    batal() {
      this.$router.push("/admin/surat_masuk");
    },
  },
};
</script>

<style scoped>
</style>