<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12">
        <v-card class="elevation-6" v-if="is_loaded">
          <v-card-text>
            <v-row>
              <v-col cols="4"> Unit Kerja </v-col>
              <v-col cols="8">
                {{ item.unit_kerja.nama }}
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="4"> No Juknis </v-col>
              <v-col cols="8">
                {{ item.no_surat }}
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="4"> Tanggal Juknis </v-col>
              <v-col cols="8">
                {{ item.tanggal_surat }}
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="4"> Perihal </v-col>
              <v-col cols="8">
                {{ item.perihal }}
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="4"> Isi </v-col>
              <v-col cols="8">
                {{ item.isi }}
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12">
                <multiple-files class="mb-6" />
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-btn color="error" class="mr-4" @click="back">
              <v-icon dark left>fas fa-arrow-left</v-icon>Kembali
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Breadcum from "../../components/breadcum";
import MultipleFiles from "../../components/multiple_files.vue";
import { JUKNIS, INFO } from "../../breadcum";
import { mapGetters } from "vuex";
import utils from "../../utils";
export default {
  name: "info_juknis",
  components: {
    Breadcum,
    MultipleFiles,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(JUKNIS(false), INFO),
    row_id: null,
    is_loaded: false,
  }),
  computed: {
    ...mapGetters({
      item: "juknis/get_item",
    }),
  },
  methods: {
    back() {
      this.$router.push("/admin/juknis");
    },
  },
  mounted() {
    const row_id = this.$route.params.row_id;
    this.$store.dispatch("juknis/fetch", row_id).then(() => {
      this.is_loaded = true;
    });
  },
};
</script>

<style scoped>
</style>