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
      <v-col cols="4" md="2" xs="4" sm="4" lg="2" xl="2">
        <v-select
              v-model="filterBy"
              :items="filterBys"
              menu-props="auto"
              item-text="nama"
              item-value="id"
              label="Filter By"
              hide-details
              prepend-icon="fas fa-filter"
              single-line
              @change="change"
            ></v-select>
      </v-col>
      <v-col cols="4" md="2" xs="4" sm="4" lg="2" xl="3">
        <v-select
              v-model="filterValue"
              :items="filterValues"
              menu-props="auto"
              item-text="nama"
              item-value="id"
              label="Values"
              hide-details
              single-line
            ></v-select>
      </v-col>
      <v-col cols="4" md="2" xs="4" sm="4" lg="2" xl="2" style="margin-left: -15px">
        <v-btn elevation="4" fab small color="blue" @click="filter()" :disabled="!valid"
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
              <th class="text-center">Kode aset_tetap</th>
              <th class="text-center">Nama aset_tetap</th>
              <th class="text-center">Kategori</th>
              <th class="text-center">Jumlah</th>
              <th class="text-center">Status</th>
              <th class="text-center action">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>
                {{ item.unit_kerja.nama }}
              </td>
              <td>
                {{ item.kode }}
              </td>
              <td>
                {{ item.nama }}
              </td>
              <td>
                {{ item.kategori.nama }}
              </td>
              <td>
                {{ item.jumlah }} {{ item.metric.nama }}
              </td>
              <td>
                <v-icon v-if="item.status == 1" color="green">fas fa-check</v-icon>&nbsp;<span v-if="item.status === 1">Baik</span>
                <v-icon v-else-if="item.status == 2" color="red">fas fa-handshake-slash</v-icon>&nbsp;<span v-if="item.status === 2">Rusak</span>
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
import { ASSET_TETAP } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "aset_tetap",
  components: {
    Paging,
    Breadcum,
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(ASSET_TETAP(true)),
    filters: [
      {
        id: "nama",
        nama: "Nama",
      },
      {
        id: "kode",
        nama: "Kode",
      },
    ],
    filterBys: [
      {
        id: "kategori_id",
        nama: "Kategori",
      },
      {
        id: "metric_id",
        nama: "Satuan Metrik",
      },
    ],
    filterBy: null,
    filterValue: null,
    filterValues: [],
    filterCari: "nama",
    searchText: "",
    isSearch: false,
    host: vm.$host,
    
  }),
  computed: {
    ...mapGetters({
      items: "aset_tetap/items",
      prev_show: "aset_tetap/prev_show",
      next_show: "aset_tetap/next_show",
      limit: "aset_tetap/limit",
      limits: "constant/limits",
      last_id: "aset_tetap/last_id",
    }),
    valid : {
      get () {
        return this.filterBy !== null && this.filterValue !== null
      }
    }
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
    }
    
  },
  methods: {
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("aset_tetap/limit", limit);
    },
    add() {
      this.$router.push("/admin/aset_tetap/create");
    },
    edit(id) {
      this.$router.push(`/admin/aset_tetap/edit/${id}`);
    },
    async remove(id) {
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
          this.$store.dispatch("aset_tetap/remove", id).then(() => {
            this.refresh();
          });
        }
      });
    },
    get() {
      this.$store.dispatch("aset_tetap/gets", {
        last_id: this.last_id,
        limit: this.limit.value,
      });
    },
    previous() {
      this.$store.dispatch("aset_tetap/prev");
    },
    next() {
      this.$store.dispatch("aset_tetap/next", this.isSearch);
    },
    cari() {
      this.$store.dispatch("aset_tetap/reset");
      this.$store.dispatch("aset_tetap/search", {
        last_id: this.last_id,
        limit: this.limit.value,
        search: this.searchText,
        filter: this.filterCari,
      });
    },
    refresh() {
      this.$store.dispatch("aset_tetap/reset");
      this.searchText = "";
      this.isSearch = false;
      this.filterCari = "nama";
      this.get();
    },
    filter() {
      this.$store.commit("aset_tetap/SET_FILTERS", [this.filterBy, this.filterValue])
      if (this.isSearch) {
        this.cari()
      } else {
        this.refresh()
      }
    },
    reset() {
      this.filterBy = null
      this.filterValue = null
      this.filterValues = []

      this.$store.commit("aset_tetap/SET_FILTERS", [])
      if (this.isSearch) {
        this.cari()
      } else {
        this.refresh()
      }
    },
    change() {
      this.filterValue = null
      if (this.filterBy === "kategori_id") {
        if (this.$store.getters["kategori/all_items"].length > 0) {
          this.filterValues = this.$store.getters["kategori/all_items"]
        } else {
          this.$store.dispatch("kategori/all").then(() => {
            this.filterValues = this.$store.getters["kategori/all_items"]
          })
        }
      } else if (this.filterBy === "metric_id") {
        if (this.$store.getters["metric/all_items"].length > 0) {
          this.filterValues = this.$store.getters["metric/all_items"]
        } else {
          this.$store.dispatch("metric/all").then(() => {
            this.filterValues = this.$store.getters["metric/all_items"]
          })
        }
      } else if (this.filterBy === "unit_kerja_id") {
        if (this.$store.getters["unit_kerja/all_items"].length > 0) {
          this.filterValues = this.$store.getters["unit_kerja/all_items"]
        } else {
          this.$store.dispatch("unit_kerja/all").then(() => {
            this.filterValues = this.$store.getters["unit_kerja/all_items"]
          })
        }
      }
    }
  },
  mounted() {
    if (utils.role() === "superuser") {
      this.filterBys.push({
        id: "unit_kerja_id",
        nama: "Unit Kerja",
      })
    }
    this.$store.commit("aset_tetap/SET_FILTERS", [])
    this.refresh();
    //this.$store.dispatch("aset_tetap/gets", { last_id: this.last_id, limit : 10 });
  },
};
</script>

<style scoped>
.action {
  width: 130px;
}
</style>