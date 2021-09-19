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
    
    <v-card class="elevation-6">
      <v-card-text>
        <v-simple-table>
          <thead>
            <tr>
              <th class="text-center">Nama</th>
              <th class="text-center">Tanggal</th>
              <th class="text-center">Jumlah</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              
              <td>
                {{ item.nama }}
              </td>
              <td>
                {{ item.tanggal }}
              </td>
              <td>
                {{ item.jumlah }}
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
import { INVENTORY_MASUK } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "inventory_masuk",
  components: {
    Paging,
    Breadcum
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(INVENTORY_MASUK(true)),
    filters: [
      {
        id: "nama",
        nama: "Nama",
      },
    ],
    filterCari: "nama",
    searchText: "",
    isSearch: false,
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      items: "inventory_masuk/items",
      prev_show: "inventory_masuk/prev_show",
      next_show: "inventory_masuk/next_show",
      limit: "inventory_masuk/limit",
      limits: "constant/limits",
    }),
  },
  methods: {
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("inventory_masuk/limit", limit);
    },
    add() {
      this.$router.push("/admin/inventory_masuk/create");
    },
    edit(id) {
      this.$router.push(`/admin/inventory_masuk/edit/${id}`);
    },
    async remove(id) {
      console.log(id);
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
          /*
          this.$store.dispatch("user/remove", { vm: this, id: id }).then(() => {
            this.get("first");
          });
          */
          console.log("REMOVE");
        }
      });
    },
    get() {
      console.log("PAGINATOR")
    },
    previous() {
      console.log("prev")
    },
    next() {
      console.log("next")
    },
    cari() {
      console.log("cari")
    },
  },
  mounted() {
    //this.get("first");
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>