<template>
  <div>
    <v-row class="cari">
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
    <v-simple-table fixed-header height="250px">
      <thead>
        <tr>
          <th class="text-center">Nama</th>
          <th class="text-center">Foto</th>
          <th class="text-center action">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>
            {{ item.nama | str_limit(20) }}
          </td>
          <td class="text-center">
            <v-avatar color="primary" size="48">
              <img :src="`${host}upload/${item.foto}`" />
            </v-avatar>
          </td>
          <td class="text-center">
            <v-tooltip top>
              <template v-slot:activator="{ on }">
                <v-btn
                  x-small
                  fab
                  dark
                  color="primary"
                  @click="pilih(item)"
                  v-on="on"
                >
                  <v-icon dark size="15">fas fa-people-carry</v-icon>
                </v-btn>
              </template>
              <span>Pilih</span>
            </v-tooltip>
          </td>
        </tr>
      </tbody>
    </v-simple-table>
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
import { mapGetters } from "vuex";
import Paging from "./paging";
export default {
  components: {
    Paging,
  },
  data: (vm) => ({
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
      items: "petugas_gambar/peoples",
      prev_show: "petugas_gambar/prev_show",
      next_show: "petugas_gambar/next_show",
      first_no: "petugas_gambar/first_no",
      last_no: "petugas_gambar/last_no",
      limit: "petugas_gambar/limit",
      limits: "constant/limits",
    }),
  },
  watch: {
    limit: function () {
      //this.get("first");
      this.getFirst()
    },
    filterCari: function () {
      //this.get("first");
      this.getFirst()
    },
    searchText: function (val) {
      if (!(!val || /^\s*$/.test(val))) {
        this.isSearch = true;
      } else {
        this.isSearch = false;
        //this.get("first");
        this.getFirst()
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
      this.$store.dispatch("petugas_gambar/limit", limit);
    },
    pilih(item) {
      this.$emit("pilih", item);
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
      this.$store.dispatch("petugas_gambar/paginator_people", query);
    },
    previous() {
      this.get("prev");
    },
    next() {
      this.get("next");
    },
    cari() {
      //this.get("first");
      this.getFirst()
    },
    refresh() {
      this.searchText = "";
    },
    getFirst() {
      this.$store.dispatch("petugas_gambar/null_peoples").then(() => {
        this.get("first");
      });
    }
  },
  mounted() {
    this.getFirst()
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
.cari {
  margin-bottom: -20px;
  margin-top: -20px;
}
</style>