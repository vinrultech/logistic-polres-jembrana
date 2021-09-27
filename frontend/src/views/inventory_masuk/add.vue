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
                  <v-select
                    v-model="barang_id"
                    :items="barangs"
                    menu-props="auto"
                    item-text="nama"
                    item-value="row_id"
                    label="Barang"
                    hide-details
                    prepend-icon="fas fa-shopping-basket"
                    single-line
                    :rules="[(v) => !!v || 'Barang dibutuhkan']"
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
                  <v-dialog
                    ref="dialog"
                    v-model="modal"
                    :return-value.sync="tanggal"
                    persistent
                    width="290px"
                  >
                    <template v-slot:activator="{ on, attrs }">
                      <v-text-field
                        v-model="tanggal"
                        :rules="[(v) => !!v || 'Tanggal dibutuhkan']"
                        label="Tanggal"
                        required
                        readonly
                        v-bind="attrs"
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker v-model="tanggal" scrollable>
                      <v-spacer></v-spacer>
                      <v-btn text color="primary" @click="modal = false">
                        Cancel
                      </v-btn>
                      <v-btn
                        text
                        color="primary"
                        @click="$refs.dialog.save(tanggal)"
                      >
                        OK
                      </v-btn>
                    </v-date-picker>
                  </v-dialog>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  
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
                <v-icon dark left>fas fa-save</v-icon>Create
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
import { INVENTORY_MASUK, ADD } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "create_metric",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(INVENTORY_MASUK(false), ADD),
    tanggal: null,
    modal: false,
    jumlah: null,
    barang_id: null,
    valid: true,
  }),
  computed: {
    ...mapGetters({
      errorMessage: "constant/errorMessage",
      barangs: "barang/all_items",
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
    this.$store.dispatch("barang/all");
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("inventory_masuk/create", {
          barang_id: this.barang_id,
          tanggal: this.tanggal,
          jumlah: parseInt(this.jumlah),
        });
      }
    },
    batal() {
      this.$router.push("/admin/inventory_masuk");
    },
  },
};
</script>

<style scoped>
</style>