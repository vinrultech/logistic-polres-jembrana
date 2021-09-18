<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
        <v-btn rounded dark color="green" @click="add()" class="mt-5">
          <v-icon dark left>fas fa-plus</v-icon>Tambah
        </v-btn>
        <v-btn rounded dark color="deep-purple" @click="refresh()" class="mt-5">
          <v-icon dark left>fas fa-sync-alt</v-icon>Refresh
        </v-btn>
      </v-col>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
        <v-row>
          <v-col cols="4">
            <v-select
              v-model="filterCari"
              :items="filters"
              menu-props="auto"
              item-text="nama"
              item-value="id"
              label="Filter"
              hide-details
              prepend-icon="fas fa-filter"
              single-line
            ></v-select>
          </v-col>
          <v-col cols="8">
            <v-text-field
              v-model="searchText"
              :placeholder="'Cari'"
              rounded
              solo
              prepend-inner-icon="fas fa-search"
              :append-icon="isSearch ? 'fas fa-times' : ''"
              single-line
              @keydown.enter="cari()"
              @click:append="refresh()"
              @click:prepend-inner="refresh()"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row>
      <template v-for="item in items">
        <v-col cols="12" xl="6" lg="6" md="6" sm="12" xs="12" :key="item.id">
          <v-card class="elevation-6">
            <v-card-title>{{ item.no_berkas }}</v-card-title>
            <v-card-text>
              <v-row>
                <v-col cols="12" xl="4" lg="4" md="4" sm="12" xs="12">
                  Pemohon :
                </v-col>
                <v-col cols="12" xl="8" lg="8" md="8" sm="12" xs="12">
                  {{ item.pemohon }} / {{ item.wa }}
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" xl="4" lg="4" md="4" sm="12" xs="12">
                  Jumlah Bidang :
                </v-col>
                <v-col cols="12" xl="8" lg="8" md="8" sm="12" xs="12">
                  {{ item.bidangs.length }}
                </v-col>
              </v-row>
              <v-stepper alt-labels>
                <v-stepper-header>
                  <v-stepper-step step="1" complete color="success">
                    <div class="text-center">
                      Input<br />
                      {{ item.tanggal_input }}
                    </div>
                  </v-stepper-step>
                  <v-divider></v-divider>
                  <v-stepper-step step="2">
                    <div class="text-center">
                      Petugas Ukur<br />
                      {{ item.nama_petugas_ukur }}
                    </div>
                  </v-stepper-step>
                  <v-divider></v-divider>
                  <v-stepper-step step="3"> Petugas Gambar </v-stepper-step>
                  <v-divider></v-divider>
                  <v-stepper-step step="4">Selesai</v-stepper-step>
                </v-stepper-header>
              </v-stepper>
            </v-card-text>
            <v-card-actions>
              <v-spacer />
              <v-tooltip top>
                <template v-slot:activator="{ on }">
                  <v-btn
                    x-small
                    fab
                    dark
                    color="error"
                    v-on="on"
                    @click="remove(item.id)"
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
                    @click="edit(item.id)"
                    v-on="on"
                  >
                    <v-icon dark size="15">fas fa-edit</v-icon>
                  </v-btn>
                </template>
                <span>Edit</span>
              </v-tooltip>
              <v-tooltip top>
                <template v-slot:activator="{ on }">
                  <v-btn
                    x-small
                    fab
                    dark
                    color="indigo"
                    @click="detail(item)"
                    v-on="on"
                  >
                    <v-icon dark size="15">fas fa-info</v-icon>
                  </v-btn>
                </template>
                <span>Bidang</span>
              </v-tooltip>
              <v-tooltip top>
                <template v-slot:activator="{ on }">
                  <v-btn
                    x-small
                    fab
                    dark
                    color="success"
                    @click="send_wa(item)"
                    v-on="on"
                  >
                    <v-icon dark size="15">fab fa-whatsapp</v-icon>
                  </v-btn>
                </template>
                <span>Whatsapp</span>
              </v-tooltip>
            </v-card-actions>
          </v-card>
        </v-col>
      </template>
    </v-row>
    <paging
      :limits="limits"
      :limit="limit"
      :prev-show="prev_show"
      :next-show="next_show"
      @next="next"
      @previous="previous"
      @changeLimit="changeLimit"
    />
    <v-dialog v-model="modal_detail" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">
            <v-icon left>fas fa-file-invoice</v-icon>Bidang</span
          >
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-simple-table>
              <thead>
                <tr>
                  <th class="text-center">No Bidang</th>
                  <th class="text-center">Jumlah</th>
                  <th class="text-center">Luas</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="bidang in bidangs" :key="bidang.id">
                  <td>
                    {{ bidang.no_bidang }}
                  </td>
                  <td>{{ bidang.jumlah }} bidang</td>
                  <td>{{ bidang.luas }} m2</td>
                </tr>
              </tbody>
            </v-simple-table>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" @click="modal_detail = false">
            Tutup<v-icon left>fas fa-window-close</v-icon>
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import Paging from "../../components/paging";
import Breadcum from "../../components/breadcum";
import { BERKAS } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "berkas",
  components: {
    Paging,
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumOne(BERKAS(true)),
    filters: [
      {
        id: "no_berkas",
        nama: "No Berkas",
      },
      {
        id: "pemohon",
        nama: "Pemohon",
      },
    ],
    filterCari: "no_berkas",
    searchText: "",
    isSearch: false,
    modal_detail: false,
    bidangs: [],
  }),
  computed: {
    ...mapGetters({
      items: "berkas/items",
      prev_show: "berkas/prev_show",
      next_show: "berkas/next_show",
      first_no: "berkas/first_no",
      last_no: "berkas/last_no",
      limit: "berkas/limit",
      limits: "constant/limits",
    }),
  },
  watch: {
    limit: function () {
      //this.get("first");
      this.getFirst();
    },
    filterCari: function () {
      //this.get("first");
      this.getFirst();
    },
    searchText: function (val) {
      if (!(!val || /^\s*$/.test(val))) {
        this.isSearch = true;
      } else {
        this.isSearch = false;
        //this.get("first");
        this.getFirst();
      }
    },
  },
  methods: {
    changeSearchText(val) {
      this.searchText = val;
    },
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("berkas/limit", limit);
    },
    add() {
      this.$router.push("/admin/berkas/create");
    },
    edit(id) {
      this.$router.push(`/admin/berkas/edit/${id}`);
    },
    detail(item) {
      this.bidangs = item.bidangs;
      this.modal_detail = true;
    },
    send_wa(item) {
      /*
      this.$store.dispatch("berkas/send_wa", {
        vm: this,
        item: {
          no: item.wa,
          message: `Bapak/ibu ${item.pemohon} dengan no berkas ${item.no_berkas} di harap hadir`,
        },
      });
      */
      const text = `Bapak/ibu ${item.pemohon} dengan no berkas ${item.no_berkas} di harap hadir`
      const message = text.replace(" ", "%20")
      const url = `https://api.whatsapp.com/send?phone=${item.wa}&text=${message}`
      window.open(url)
    },
    async remove(id) {
      this.$swal({
        title: "Anda yakin?",
        text: "Apakah anda ingin menhapus data!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya",
        cancelButtonText: "Tidak",
      }).then((result) => {
        if (result.value) {
          this.$store
            .dispatch("berkas/remove", { vm: this, id: id })
            .then(() => {
              //this.get("first");
              this.getFirst();
            });
        }
      });
    },
    get(action) {
      let no = 0;
      if (action === "next") {
        no = this.last_no;
      } else if (action === "prev") {
        no = this.first_no;
      } else {
        no = 0;
      }
      let query = {
        limit: this.limit.value,
        action: action,
        no: no,
        vm: this,
      };
      if (
        !(
          !this.searchText || /^\s*$/.test(this.searchText)
        ) /*this.search !== "" || this.search !== null*/
      ) {
        query = {
          ...query,
          search: this.searchText,
          filter: this.filterCari,
        };
      }
      this.$store.dispatch("berkas/paginator", query);
    },
    previous() {
      this.get("prev");
    },
    next() {
      this.get("next");
    },
    cari() {
      //this.get("first");
      this.getFirst();
    },
    refresh() {
      this.searchText = "";
      //this.get("first");
    },
    getFirst() {
      this.$store.dispatch("berkas/null_items").then(() => {
        this.get("first");
      });
    },
  },
  mounted() {
    this.getFirst();
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>