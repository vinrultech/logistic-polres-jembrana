<template>
  <v-simple-table fixed-header height="300px">
    <thead>
      <tr>
        <th class="text-center">No Bidang</th>
        <th class="text-center">Jumlah</th>
        <th class="text-center">Luas</th>
        <th class="text-center">Status</th>
        <th class="text-center">Tanggal Setor</th>
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
          <v-icon v-if="item.status_ukur === 1" color="success"
            >fas fa-check-circle</v-icon
          >
          <v-icon v-else color="error">fas fa-times-circle</v-icon>
        </td>
        <td>
          <v-dialog
            v-model="modal_date"
            persistent
            width="290px"
            :retain-focus="false"
          >
            <v-date-picker v-model="tanggal_ukur" scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="modal_date = false">
                Cancel
              </v-btn>
              <v-btn text color="primary" @click="save(tanggal_ukur)">
                OK
              </v-btn>
            </v-date-picker>
          </v-dialog>
          <v-row>
            <v-col cols="6"
              ><span class="float-right mt-2">{{
                item.tanggal_ukur
              }}</span></v-col
            >
            <v-col cols="6"
              ><v-btn
                class="float-left"
                icon
                @click="openDate(item.id)"
                v-if="item.status_ukur === 0 && !is_home"
                ><v-icon>fas fa-calendar-alt</v-icon></v-btn
              ><v-btn class="float-left" icon v-else
                ><v-icon>fas fa-calendar-alt</v-icon></v-btn
              ></v-col
            >
          </v-row>
        </td>
        <td class="text-center" v-if="!is_home">
          <div v-if="item.status_ukur === 0">
            <v-tooltip right>
              <template v-slot:activator="{ on }">
                <v-btn
                  x-small
                  fab
                  dark
                  color="success"
                  @click="proses(item)"
                  v-on="on"
                >
                  <v-icon dark size="15">fas fa-microchip</v-icon>
                </v-btn>
              </template>
              <span>Proses</span>
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
    tanggal_ukur: null,
    modal_date: false,
    bidang_id: null,
  }),
  methods: {
    save(tanggal_ukur) {
      this.modal_date = false;
      this.$store.commit("petugas_ukur/CHANGE_DATE", {
        tanggal_ukur: tanggal_ukur,
        bidang_id: this.bidang_id,
        id: this.id,
      });
    },
    proses(item) {
      if (item.tanggal_ukur === "" || item.tanggal_ukur === null) {
        this.$swal({
          icon: "error",
          title: "Oops...",
          text: "Tanggal setor tidak boleh kosong",
        });
      } else {
        //console.log(item)

        this.$store.dispatch("petugas_ukur/proses", {
          item: {
            tanggal: item.tanggal_ukur,
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

