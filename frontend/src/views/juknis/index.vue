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
          <v-col cols="12" md="4" xs="12" sm="12" lg="4" xl="4">
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
          <v-col cols="12" md="8" xs="12" sm="12" lg="8" xl="8">
            <v-row>
              <v-col cols="10">
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
              <v-col cols="2" style="margin-left: -15px">
                <v-btn elevation="4" fab small color="green" @click="cari()"
                  ><v-icon small color="white">fas fa-search</v-icon></v-btn
                >
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <v-row style="margin-top: -60px">
      <v-col cols="8" md="3" xs="8" sm="8" lg="3" xl="3">
        <v-dialog
          ref="dialog"
          v-model="modal"
          persistent
          width="290px"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="dateRangeText"
              label="Filter by Tanggal"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker v-model="dates" range>
            <v-spacer></v-spacer>
            <v-btn text color="primary" @click="modal = false"> Cancel </v-btn>
            <v-btn text color="primary" @click="saveDate(dates)">
              OK
            </v-btn>
          </v-date-picker>
        </v-dialog>
      </v-col>
      <v-col cols="4" md="2" xs="4" sm="4" lg="2" xl="2" style="margin-left: -15px">
        <v-btn elevation="4" fab small color="blue" @click="filter()"
          ><v-icon small color="white">fas fa-filter</v-icon></v-btn
        >
        &nbsp;&nbsp;
        <v-btn elevation="4" fab small color="red" @click="reset()"
          ><v-icon small color="white">fas fa-undo</v-icon></v-btn
        >
      </v-col>
      
    </v-row>

    <v-card class="elevation-6">
      <v-card-text>
        <v-simple-table>
          <thead>
            <tr>
              <th class="text-center">Unit Kerja</th>
              <th class="text-center">No Juknis</th>
              <th class="text-center">Tanggal Juknis</th>
              <th class="text-center">Perihal</th>
              <th class="text-center action">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>
                {{ item.unit_kerja.nama }}
              </td>
              <td>
                {{ item.no_surat }}
              </td>
              <td>
                {{ item.tanggal_surat }}
              </td>
              <td>
                {{ item.perihal }}
              </td>
              <td class="text-center">
                <v-tooltip left>
                  <template v-slot:activator="{ on }">
                    <v-btn
                      x-small
                      fab
                      dark
                      color="error"
                      v-on="on"
                      @click="remove(item.row_id)"
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
                      @click="edit(item.row_id)"
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
                      color="green"
                      @click="info(item.row_id)"
                      v-on="on"
                    >
                      <v-icon dark size="15">fas fa-info</v-icon>
                    </v-btn>
                  </template>
                  <span>Detail</span>
                </v-tooltip>
              </td>
            </tr>
          </tbody>
        </v-simple-table>
      </v-card-text>
    </v-card>
    <paging
      :limits="limits"
      :limit="limit"
      :prev-show="prev_show"
      :next-show="next_show"
      @next="next"
      @previous="previous"
      @changeLimit="changeLimit"
    />
  </div>
</template>

<script>
import Paging from "../../components/paging";
import Breadcum from "../../components/breadcum";
import { JUKNIS } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
import moment from 'moment';
export default {
  name: "juknis",
  components: {
    Paging,
    Breadcum,
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(JUKNIS(true)),
    filters: [
      {
        id: "no_surat",
        nama: "No Surat",
      },
      {
        id: "perihal",
        nama: "Perihal",
      },
    ],
    filterCari: "perihal",
    searchText: "",
    isSearch: false,
    host: vm.$host,
    modal: false,
    dates: [],
  }),
  computed: {
    ...mapGetters({
      items: "juknis/items",
      prev_show: "juknis/prev_show",
      next_show: "juknis/next_show",
      limit: "juknis/limit",
      limits: "constant/limits",
      last_id: "juknis/last_id",
    }),
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
  },
  watch: {
    limit: function () {
      this.refresh();
    },
    filterCari: function () {
      this.cari();
    },
    searchText: function (val) {
      if (!(!val || /^\s*$/.test(val))) {
        this.isSearch = true;
      } else {
        this.isSearch = false;
        this.refresh();
      }
    },
  },
  methods: {
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("juknis/limit", limit);
    },
    add() {
      this.$router.push("/admin/juknis/create");
    },
    edit(id) {
      this.$router.push(`/admin/juknis/edit/${id}`);
    },
    info(id) {
      this.$router.push(`/admin/juknis/info/${id}`);
    },
    async remove(row_id) {
      //console.log(id);
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
          this.$store.dispatch("juknis/remove", row_id).then(() => {
            this.refresh();
          });
        }
      });
    },
    filter() {
      this.$store.commit("juknis/SET_DATES", this.dates)
      if (this.isSearch) {
        this.cari()
      } else {
        this.refresh()
      }
    },
    reset() {
      this.dates = []
      this.$store.commit("juknis/SET_DATES", [])
      if (this.isSearch) {
        this.cari()
      } else {
        this.refresh()
      }
    },
    get() {
      this.$store.dispatch("juknis/gets", {
          last_id: this.last_id,
          limit: this.limit.value,
        });
      
    },
    previous() {
      this.$store.dispatch("juknis/prev");
    },
    next() {
      this.$store.dispatch("juknis/next", this.isSearch);
    },
    cari() {
      this.$store.dispatch("juknis/reset");
      this.$store.dispatch("juknis/search", {
        last_id: this.last_id,
        limit: this.limit.value,
        search: this.searchText,
        filter: this.filterCari,
      });
    },
    refresh() {
      this.$store.dispatch("juknis/reset");
      this.searchText = "";
      this.isSearch = false;
      this.filterCari = "perihal";
      this.get();
    },
    saveDate(dates) {
      if (dates.length >= 2) {
        const date1 = moment(dates[0])
        const date2 = moment(dates[1])
        const jml = date2.diff(date1)
        
        if (jml < 0) {
          this.dates = []
          this.$toastr.e("Pilihan tanggal kedua harus lebih dari tanggal kedua")
        } else {
          this.dates = dates
          this.modal = false
          this.filter()
        }
        
      } else {
        this.$toastr.e("Harus pilih dua tanggal")
      }
      
    }
  },
  mounted() {
    this.$store.commit("juknis/SET_DATES", [])
    this.refresh();
  },
};
</script>

<style scoped>
.action {
  width: 130px;
}
</style>