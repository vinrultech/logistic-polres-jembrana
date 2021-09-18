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
          <v-toolbar>
            <v-toolbar-title>Edit Berkas</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-tabs
              v-model="tab"
              background-color="blue-grey"
              centered
              dark
              icons-and-text
            >
              <v-tabs-slider></v-tabs-slider>

              <v-tab href="#tab-1">
                Berkas
                <v-icon>fas fa-id-card-alt</v-icon>
              </v-tab>

              <v-tab href="#tab-2">
                Bidang
                <v-icon>fas fa-file-invoice</v-icon>
              </v-tab>
            </v-tabs>

            <v-tabs-items v-model="tab">
              <v-tab-item value="tab-1">
                <v-card flat>
                  <v-card-text>
                    <v-form ref="form" v-model="valid">
                      <v-row>
                        <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                          <v-text-field
                            v-model="no_berkas"
                            :rules="[(v) => !!v || 'No berkas dibutuhkan']"
                            label="No Berkas"
                            prepend-icon="fas fa-file-code"
                            required
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                          <v-text-field
                            v-model="pemohon"
                            :rules="[(v) => !!v || 'Nama pemohon dibutuhkan']"
                            label="Nama Pemohon"
                            prepend-icon="fas fa-male"
                            required
                          ></v-text-field>
                        </v-col>
                      </v-row>
                      <v-row>
                        <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                          <v-text-field
                            v-model="wa"
                            label="No WA"
                            prepend-icon="fab fa-whatsapp"
                            
                          ></v-text-field>
                        </v-col>
                        <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                          <v-dialog
                            ref="dialog"
                            v-model="modal_date"
                            :return-value.sync="tanggal_input"
                            persistent
                            width="290px"
                          >
                            <template v-slot:activator="{ on, attrs }">
                              <v-text-field
                                v-model="tanggal_input"
                                label="Tanggal Input"
                                prepend-icon="fas fa-calendar-alt"
                                readonly
                                v-bind="attrs"
                                v-on="on"
                              ></v-text-field>
                            </template>
                            <v-date-picker v-model="tanggal_input" scrollable>
                              <v-spacer></v-spacer>
                              <v-btn
                                text
                                color="primary"
                                @click="modal_date = false"
                              >
                                Cancel
                              </v-btn>
                              <v-btn
                                text
                                color="primary"
                                @click="$refs.dialog.save(tanggal_input)"
                              >
                                OK
                              </v-btn>
                            </v-date-picker>
                          </v-dialog>
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
                        <v-icon dark left>fas fa-save</v-icon>Update
                      </v-btn>
                    </v-form>
                  </v-card-text>
                </v-card>
              </v-tab-item>
              <v-tab-item value="tab-2">
                <v-card flat>
                  <v-card-title>
                    <v-btn
                      rounded
                      dark
                      color="green"
                      @click="modal_bidang = true"
                    >
                      <v-icon dark left>fas fa-file-invoice</v-icon>Tambah
                      Bidang
                    </v-btn>
                  </v-card-title>
                  <v-card-text>
                    <v-simple-table>
                      <thead>
                        <tr>
                          <th class="text-center">No Bidang</th>
                          <th class="text-center">Jumlah</th>
                          <th class="text-center">Luas</th>
                          <th class="text-center action">Action</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="bidang in bidangs" :key="bidang.id">
                          <td>
                            {{ bidang.no_bidang }}
                          </td>
                          <td>{{ bidang.jumlah }} bidang</td>
                          <td>{{ bidang.luas }} m2</td>
                          <td class="text-center">
                            <v-tooltip left>
                              <template v-slot:activator="{ on }">
                                <v-btn
                                  x-small
                                  fab
                                  dark
                                  color="error"
                                  v-on="on"
                                  @click="remove(bidang.id)"
                                >
                                  <v-icon dark size="15">fas fa-trash</v-icon>
                                </v-btn>
                              </template>
                              <span>Hapus</span>
                            </v-tooltip>
                            <v-tooltip top>
                              <template v-slot:activator="{ on }">
                                <v-btn
                                  x-small
                                  fab
                                  dark
                                  color="primary"
                                  @click="edit(bidang)"
                                  v-on="on"
                                >
                                  <v-icon dark size="15">fas fa-edit</v-icon>
                                </v-btn>
                              </template>
                              <span>Edit</span>
                            </v-tooltip>
                          </td>
                        </tr>
                      </tbody>
                    </v-simple-table>
                  </v-card-text>
                </v-card>
              </v-tab-item>
            </v-tabs-items>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-dialog v-model="modal_bidang" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">
            <v-icon left>fas fa-file-invoice</v-icon>Tambah Bidang</span
          >
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-form ref="form_bidang" v-model="valid_bidang">
              <v-row>
                <v-col cols="12">
                  <v-text-field
                    v-model="no_bidang"
                    :rules="[(v) => !!v || 'No bidang dibutuhkan']"
                    label="No Bidang"
                    prepend-icon="fas fa-shield-virus"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="jumlah"
                    :rules="[(v) => !!v || 'Jumlah dibutuhkan']"
                    label="Jumlah"
                    type="number"
                    suffix="bidang"
                    prepend-icon="fas fa-sort-numeric-down"
                    required
                  ></v-text-field>
                </v-col>
                <v-col cols="12" md="6" lg="6" xl="6" sm="12" xs="12">
                  <v-text-field
                    v-model="luas"
                    :rules="[(v) => !!v || 'Luas dibutuhkan']"
                    label="Luas"
                    type="number"
                    suffix="m2"
                    prepend-icon="fas fa-th-large"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-form>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="close_bidang">
            Tutup<v-icon left>fas fa-window-close</v-icon>
          </v-btn>
          <v-btn
            color="blue darken-1"
            @click="save_bidang(edit_bidang)"
            :disabled="!valid_bidang"
          >
            <span v-if="edit_bidang">Update</span><span v-else>Simpan</span
            ><v-icon left>fas fa-save</v-icon>
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import moment from "moment";
import { v4 as uuidv4 } from "uuid";
import Breadcum from "../../components/breadcum";
import { BERKAS, EDIT } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "edit_berkas",
  components: {
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(BERKAS(false), EDIT),
    no_bidang: null,
    jumlah: null,
    luas: null,
    pemohon: null,
    no_berkas: null,
    wa: null,
    tanggal_input: moment().format("YYYY-MM-DD").substr(0, 10),
    modal_date: false,
    modal_bidang: false,
    valid: true,
    valid_bidang: true,
    edit_bidang: false,
    id_bidang: null,
    id: null,
    tab: null,
  }),
  computed: {
    ...mapGetters({
      bidangs: "berkas/bidangs",
    }),
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["berkas/item"](parseInt(id));
    if (item != null || item != undefined) {
      this.no_berkas = item.no_berkas;
      this.pemohon = item.pemohon;
      this.wa = item.wa;
      this.tanggal_input = item.tanggal_input;
      this.id = id;
      //this.bidangs = item.bidangs;
      this.$store.dispatch("berkas/add_bidangs", item.bidangs);
    }
  },
  methods: {
    simpan() {
      if (this.$refs.form.validate()) {
        this.$store.dispatch("berkas/edit_update_berkas", {
          item: {
            no_berkas: this.no_berkas,
            pemohon: this.pemohon,
            wa: this.wa,
            tanggal_input: this.tanggal_input,
          },
          vm: this,
          id: this.id,
        });
      }
    },
    save() {},
    batal() {
      this.$router.push("/admin/berkas");
    },
    edit(item) {
      this.id_bidang = item.id;
      this.no_bidang = item.no_bidang;
      this.luas = item.luas;
      this.jumlah = item.jumlah;
      this.modal_bidang = true;
      this.edit_bidang = true;
    },
    remove(id) {
      
      this.$swal({
        title: "Anda yakin?",
        text: "Apakah anda ingin menhapus bidang!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya",
        cancelButtonText: "Tidak",
      }).then((result) => {
        if (result.value) {
          this.$store.dispatch("berkas/edit_remove_bidang", {
            id: id,
            vm: this,
          });
        }
      });
    },
    save_bidang(edit_bidang) {
      if (this.$refs.form_bidang.validate()) {
        if (edit_bidang) {
          this.$store
            .dispatch("berkas/edit_update_bidang", {
              item: {
                no_bidang: this.no_bidang,
                jumlah: parseInt(this.jumlah),
                luas: parseFloat(this.luas),
                berkas_id: parseInt(this.id),
              },
              id: this.id_bidang,
              vm: this,
            })
            .then(() => {
              this.modal_bidang = false;
              this.no_bidang = null;
              this.jumlah = null;
              this.luas = null;
              this.id_bidang = null;
              this.edit_bidang = false;
            });
        } else {
          this.$store
            .dispatch("berkas/edit_add_bidang", {
              item: {
                id_bidang: uuidv4(),
                no_bidang: this.no_bidang,
                jumlah: parseInt(this.jumlah),
                luas: parseFloat(this.luas),
                berkas_id: parseInt(this.id),
              },
              vm: this,
            })
            .then(() => {
              this.modal_bidang = false;
              this.no_bidang = null;
              this.jumlah = null;
              this.luas = null;
              this.id_bidang = null;
              this.edit_bidang = false;
            });
        }
      }
    },
    close_bidang() {
      this.no_bidang = null;
      this.jumlah = null;
      this.luas = null;
      this.modal_bidang = false;
    },
  },
};
</script>

<style scoped>
</style>