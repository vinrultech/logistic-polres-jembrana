<template>
  <v-simple-table fixed-header height="300px">
    <thead>
      <tr>
        <th class="text-center">No Bidang</th>
        <th class="text-center">Jumlah</th>
        <th class="text-center">Luas</th>
        <th class="text-center">Status</th>
        <th class="text-center">Tanggal Selesai</th>
        <th class="text-center action" v-if="!is_home">Action</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in items" :key="item.id">
        <td>
          {{ item.no_bidang }}
        </td>
        <td class="text-right">{{ item.jumlah }} bidang</td>
        <td class="text-right">{{ item.luas }} m2</td>
        <td class="text-center">
          <v-icon v-if="item.status_gambar === 1 && item.status_problem == 0" color="success"
          >fas fa-check-circle</v-icon>
          <v-icon v-else-if="item.status_problem === 1 && item.status_gambar === 0" color="error">fas fa-exclamation-triangle</v-icon>
          <v-icon v-else-if="item.status_gambar === 0 && item.status_problem == 0" color="error"
          >fas fa-times-circle</v-icon>
        </td>
        <td class="text-center">
          <v-dialog
            v-model="modal_date"
            persistent
            width="290px"
            :retain-focus="false"
          >
            <v-date-picker v-model="tanggal_selesai" scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal_date = false">
                Cancel
              </v-btn>
              <v-btn text color="primary" @click="save(tanggal_selesai)">
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
          <v-row>
            <v-col cols="6">
              <span v-if="item.status_gambar === 1 && item.status_problem == 0">{{ item.tanggal_gambar }}</span>
              <span v-else class="float-right mt-2">{{item.tanggal_selesai}}</span>
            </v-col>
            <v-col cols="6"
              ><v-btn
                class="float-left"
                icon
                @click="openDate(item.id)"
                v-if="item.status_problem === 1 && !is_home"
                ><v-icon>fas fa-calendar-alt</v-icon></v-btn
              ><v-btn class="float-left" icon v-else
                ><v-icon>fas fa-calendar-alt</v-icon></v-btn
              ></v-col
            >
          </v-row>
        </td>
        <td class="text-center" v-if="!is_home">
          <div v-if="item.status_problem === 1">
            <v-tooltip left>
              <template v-slot:activator="{ on }">
                <v-btn
                  x-small
                  fab
                  dark
                  color="warning"
                  @click="finish(item)"
                  v-on="on"
                >
                  <v-icon dark size="15">fas fa-print</v-icon>
                </v-btn>
              </template>
              <span>Print</span>
            </v-tooltip>
            <v-tooltip top>
              <template v-slot:activator="{ on }">
                <v-btn
                  x-small
                  fab
                  dark
                  color="primary"
                  @click="finish(item)"
                  v-on="on"
                >
                  <v-icon dark size="15">fab fa-whatsapp</v-icon>
                </v-btn>
              </template>
              <span>Whatsapp</span>
            </v-tooltip>
            <v-tooltip right>
              <template v-slot:activator="{ on }">
                <v-btn
                  x-small
                  fab
                  dark
                  color="success"
                  @click="finish(item)"
                  v-on="on"
                >
                  <v-icon dark size="15">fas fa-check-double</v-icon>
                </v-btn>
              </template>
              <span>Finish</span>
            </v-tooltip>
          </div>
        </td>
      </tr>
    </tbody>
  </v-simple-table>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
    id: {
      required: true,
    },
    is_home: {
      type: Boolean,
      default: false,
    },
  },
  data: (vm) => ({
    host: vm.$host,
    tanggal_selesai: null,
    modal_date: false,
    bidang_id: null,
  }),
  methods: {
    save(tanggal_selesai) {
      this.modal_date = false;
      this.$store.commit("problem/CHANGE_DATE", {
        tanggal_selesai: tanggal_selesai,
        bidang_id: this.bidang_id,
        id: this.id,
      });
    },
    finish(item) {
      if (item.tanggal_selesai === "" || item.tanggal_selesai === null) {
        this.$swal({
          icon: "error",
          title: "Oops...",
          text: "Tanggal selesai tidak boleh kosong",
        });
      } else {
        this.$store.dispatch("problem/proses", {
          item: {
            tanggal: item.tanggal_selesai,
          },
          vm: this,
          id: item.id,
        });
      }
    },
    openDate(id) {
      this.bidang_id = id;
      this.modal_date = true;
    },
  },
};
</script>

<style scoped>
</style>

