<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
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
          <card-petugas-ukur :item="item" :is_detail="true" @bidang="bidang" @detail="detail" />
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
    <bidang-petugas-ukur
      :no_berkas="no_berkas"
      :id="id"
      :visible="modal_bidang"
      :items="bidangs"
      @tutup="tutup"
    />
  </div>
</template>

<script>
import Paging from "../../components/paging";
import BidangPetugasUkur from "../../components/bidang_petugas_ukur";
import CardPetugasUkur from "../../components/card_petugas_ukur";
import Breadcum from "../../components/breadcum";
import { PETUGAS_UKUR } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "petugas_ukur",
  components: {
    Paging,
    BidangPetugasUkur,
    CardPetugasUkur,
    Breadcum
  },
  data: () => ({
    breadcums: utils.breadcumOne(PETUGAS_UKUR(true)),
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
  }),
  computed: {
    ...mapGetters({
      items: "petugas_ukur/items",
      prev_show: "petugas_ukur/prev_show",
      next_show: "petugas_ukur/next_show",
      first_no: "petugas_ukur/first_no",
      last_no: "petugas_ukur/last_no",
      limit: "petugas_ukur/limit",
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
      this.$store.dispatch("petugas_ukur/limit", limit);
    },
    bidang(item) {
      this.bidangs = item.bidangs;
      this.modal_bidang = true;
      this.no_berkas = item.no_berkas;
      this.id = item.id;
    },
    tutup(is_pilih) {
      if (is_pilih) {
        this.getFirst();
      }
      this.modal_bidang = false

    },
    detail(item) {
      this.$router.push(`/admin/petugas_ukur/detail/${item.id}`);
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
      this.$store.dispatch("petugas_ukur/paginator", query);
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
      this.$store.dispatch("petugas_ukur/null_items").then(() => {
        this.get("first");
      });
    }
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