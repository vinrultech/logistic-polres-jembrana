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

    
    <v-card class="elevation-6">
      <v-card-text>
        <v-simple-table>
          <thead>
            <tr>
              <th class="text-center">Unit Kerja</th>
              <th class="text-center">Kode Barang</th>
              <th class="text-center">Nama Barang</th>
              <th class="text-center">Kategori</th>
              <th class="text-center">Tanggal Keluar</th>
              <th class="text-center">Jumlah</th>
              <th class="text-center action">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>
                {{ item.unit_kerja.nama }}
              </td>
              <td>
                {{ item.barang.kode }}
              </td>
              <td>
                {{ item.barang.nama }}
              </td>
              <td>
                {{ item.barang.kategori.nama }}
              </td>
              <td>
                {{ item.tanggal }}
              </td>
              <td>
                {{ item.jumlah }} {{ item.barang.metric.nama }}
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
import { INVENTORY_KELUAR } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "inventory_keluar",
  components: {
    Paging,
    Breadcum,
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(INVENTORY_KELUAR(true)),
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
    filterCari: "nama",
    searchText: "",
    isSearch: false,
    host: vm.$host,
    
  }),
  computed: {
    ...mapGetters({
      items: "inventory_keluar/items",
      prev_show: "inventory_keluar/prev_show",
      next_show: "inventory_keluar/next_show",
      limit: "inventory_keluar/limit",
      limits: "constant/limits",
      last_id: "inventory_keluar/last_id",
    }),
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
      this.$store.dispatch("inventory_keluar/limit", limit);
    },
    add() {
      this.$router.push("/admin/inventory_keluar/create");
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
          this.$store.dispatch("inventory_keluar/remove", id).then(() => {
            this.refresh();
          });
        }
      });
    },
    get() {
      this.$store.dispatch("inventory_keluar/gets", {
        last_id: this.last_id,
        limit: this.limit.value,
      });
    },
    previous() {
      this.$store.dispatch("inventory_keluar/prev");
    },
    next() {
      this.$store.dispatch("inventory_keluar/next", this.isSearch);
    },
    cari() {
      this.$store.dispatch("inventory_keluar/reset");
      this.$store.dispatch("inventory_keluar/search", {
        last_id: this.last_id,
        limit: this.limit.value,
        search: this.searchText,
        filter: this.filterCari,
      });
    },
    refresh() {
      this.$store.dispatch("inventory_keluar/reset");
      this.searchText = "";
      this.isSearch = false;
      this.filterCari = "nama";
      this.get();
    }
  },
  mounted() {
    this.refresh();
    //this.$store.dispatch("inventory_keluar/gets", { last_id: this.last_id, limit : 10 });
  },
};
</script>

<style scoped>
.action {
  width: 130px;
}
</style>