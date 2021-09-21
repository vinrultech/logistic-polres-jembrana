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
              <v-col cols="2" style="margin-left:-15px;">
                <v-btn elevation="4" fab small color="green" @click="cari()"><v-icon small>fas fa-search</v-icon></v-btn>
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
              <th class="text-center">Kode</th>
              <th class="text-center">Nama</th>
              <th class="text-center action">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>
                {{ item.kode }}
              </td>
              <td>
                {{ item.nama }}
              </td>
              <td class="text-center">
                <v-tooltip left>
                  <template v-slot:activator="{ on }">
                    <v-btn
                      x-small
                      fab
                      dark
                      color="error"
                      v-if="item.username !== username"
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
                      v-if="item.username !== username"
                      @click="edit(item.id)"
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
import { KATEGORI } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "kategori",
  components: {
    Paging,
    Breadcum,
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(KATEGORI(true)),
    filters: [
      {
        id: "kode",
        nama: "Kode",
      },
      {
        id: "nama",
        nama: "Nama",
      },
    ],
    filterCari: "nama",
    username: utils.username(),
    searchText: "",
    isSearch: false,
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      items: "kategori/items",
      prev_show: "kategori/prev_show",
      next_show: "kategori/next_show",
      limit: "kategori/limit",
      limits: "constant/limits",
      last_id: "kategori/last_id",
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
    },
  },
  methods: {
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("kategori/limit", limit);
    },
    add() {
      this.$router.push("/admin/kategori/create");
    },
    edit(id) {
      this.$router.push(`/admin/kategori/edit/${id}`);
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
          this.$store.dispatch("kategori/remove", id).then(() => {
            this.refresh();
          });
        }
      });
    },
    get() {
      this.$store.dispatch("kategori/gets", {
        last_id: this.last_id,
        limit: this.limit.value,
      });
    },
    previous() {
      this.$store.dispatch("kategori/prev");
    },
    next() {
      this.$store.dispatch("kategori/next", this.isSearch);
    },
    cari() {
      this.$store.dispatch("kategori/reset");
      this.$store.dispatch("kategori/search", {
        last_id: this.last_id,
        limit: this.limit.value,
        search: this.searchText,
        filter: this.filterCari,
      });
    },
    refresh() {
      this.$store.dispatch("kategori/reset");
      this.searchText = "";
      this.isSearch = false;
      this.filterCari = "nama";
      this.get();
    },
  },
  mounted() {
    this.refresh();
    //this.$store.dispatch("kategori/gets", { last_id: this.last_id, limit : 10 });
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>