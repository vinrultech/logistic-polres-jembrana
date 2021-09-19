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
                    v-model="no_juknis"
                    :rules="[(v) => !!v || 'No Juknis dibutuhkan']"
                    label="No Juknis"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="tanggal_juknis"
                    :rules="[(v) => !!v || 'Tanggal Juknis dibutuhkan']"
                    label="Tanggal Juknis"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="perihal"
                    :rules="[(v) => !!v || 'Perihal Juknis dibutuhkan']"
                    label="Juknis"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12">
                  <v-textarea
                    v-model="isi"
                    :rules="[(v) => !!v || 'Isi dibutuhkan']"
                    label="Isi Juknis"
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
import { JUKNIS, EDIT } from "../../breadcum";
import utils from "../../utils";
export default {
  name: "edit_juknis",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(JUKNIS(false), EDIT),
    id: null,
    no_juknis: null,
    tanggal_juknis: null,
    perihal: null,
    isi: null,
    valid: true,
  }),
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("juknis/update", {
          item: {
            kode: this.kode,
            nama: this.nama,
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/juknis");
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["juknis/item"](parseInt(id));
    if (item != null || item != undefined) {
      (this.no_juknis = item.no_juknis),
        (this.tanggal_juknis = item.tanggal_juknis),
        (this.perihal = item.perihal),
        (this.isi = item.isi),
        (this.id = id);
    }
  },
};
</script>

<style scoped>
</style>