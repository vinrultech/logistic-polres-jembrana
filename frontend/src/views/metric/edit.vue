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
                <v-col cols="12">
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
import { METRIC, EDIT } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "edit_metric",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(METRIC(false), EDIT),
    id: null,
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
  methods: {
    update() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("metric/update", {
          item: {
            nama: this.nama
          },
          id: this.id,
        });
      }
    },
    batal() {
      this.$router.push("/admin/metric");
    },
  },
  mounted() {
    this.$store.dispatch("constant/set_error", false);
    const id = this.$route.params.id;
    const item = this.$store.getters["metric/item"](parseInt(id));
    if (item != null || item != undefined) {
      this.nama = item.nama;
      this.id = id;
    }
  },
};
</script>

<style scoped>
</style>