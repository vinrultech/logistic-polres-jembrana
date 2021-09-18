<template>
  <div>
    <v-row>
      <v-col cols="12">
        <v-avatar color="primary" size="60">
          <img :src="`${host}upload/${petugas.foto}`" />
        </v-avatar>
        {{ petugas.nama }}
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
        <v-btn rounded dark color="error" @click="tutup" class="mt-5">
          <v-icon dark left>fas fa-arrow-left</v-icon>Kembali
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
          <card-petugas-ukur
            :item="item"
            :is_detail="true"
            :is_home="true"
            @detail="detail"
            v-if="item.status === 0"
          />
          <card-petugas-gambar
            :item="item"
            :is_detail="true"
            :is_home="true"
            @detail="detail"
            v-if="item.status === 1"
          />
          <card-problem
            :item="item"
            :is_detail="true"
            :is_home="true"
            @detail="detail"
            v-if="item.status === 2"
          />
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
  </div>
</template>

<script>
import Paging from "../components/paging";
import CardPetugasGambar from "../components/card_petugas_gambar";
import CardPetugasUkur from "../components/card_petugas_ukur";
import CardProblem from "../components/card_problem";

import { mapGetters } from "vuex";
export default {
  name: "home_petugas_ukur",
  components: {
    Paging,
    CardPetugasGambar,
    CardPetugasUkur,
    CardProblem,
  },
  data: (vm) => ({
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
    modal_bidang: false,
    bidangs: [],
    no_berkas: null,
    id: null,
    petugas: {},
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      items: "home/items",
      prev_show: "home/prev_show",
      next_show: "home/next_show",
      first_no: "home/first_no",
      last_no: "home/last_no",
      limit: "home/limit",
      limits: "constant/limits",
    }),
  },
  watch: {
    limit: function () {
      this.getFirst()
    },
    filterCari: function () {
      this.getFirst()
    },
    searchText: function (val) {
      if (!(!val || /^\s*$/.test(val))) {
        this.isSearch = true;
      } else {
        this.isSearch = false;
        this.getFirst()
      }
    },
  },
  methods: {
    tutup() {
      this.$router.go(-1);
    },
    changeSearchText(val) {
      this.searchText = val;
    },
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("home/limit", limit);
    },
    find_items(id) {
      return this.items.find((item) => {
        return item.id === id;
      });
    },
    bidang(item) {
      this.bidangs = item.bidangs;
      this.modal_bidang = true;
      this.no_berkas = item.no_berkas;
      this.id = item.id;
    },
    detail(item) {
      this.$router.push(`/berkas/${item.id}`);
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
        petugas_id: this.petugas.id,
        "role": "ukur",
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
      this.$store.dispatch("home/paginator_petugas", query);
    },
    previous() {
      this.get("prev");
    },
    next() {
      this.get("next");
    },
    cari() {
      this.getFirst()
    },
    refresh() {
      this.searchText = "";
      //this.get("first");
    },
    getFirst() {
      this.$store.dispatch("home/null_items").then(() => {
        this.get("first");
      });
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["home/petugas_ukur"](parseInt(id));
    if (item != null || item != undefined) {
      this.petugas = item;
    }
    this.getFirst()
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}

.text-step {
  font-size: 12px;
}

.done {
  color: green;
}
.undone {
  color: red;
}
</style>