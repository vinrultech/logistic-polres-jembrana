<template>
  <div>
    <v-row>
      <v-col cols="12">
        <card-petugas-ukur
          :item="item"
          :is_detail="false"
          :is_home="true"
          @detail="detail"
          v-if="item.status === 0"
        />
        <card-petugas-gambar
          :item="item"
          :is_detail="false"
          :is_home="true"
          @detail="detail"
          v-if="item.status === 1"
        />
        <card-problem
          :item="item"
          :is_detail="false"
          :is_home="true"
          @detail="detail"
          v-if="item.status === 2"
        />
      </v-col>
    </v-row>
    <table-bidang-petugas-ukur
      :items="item.bidangs"
      :id="item.id"
      :is_home="true"
      v-if="item.status === 0"
    />
    <table-bidang-petugas-gambar
      :items="item.bidangs"
      :id="item.id"
      :is_home="true"
      v-if="item.status === 1"
    />
    <table-bidang-problem
      :items="item.bidangs"
      :id="item.id"
      :is_home="true"
      v-if="item.status === 2"
    />
  </div>
</template>

<script>
import CardPetugasGambar from "../components/card_petugas_gambar";
import CardPetugasUkur from "../components/card_petugas_ukur";
import CardProblem from "../components/card_problem";

import TableBidangPetugasGambar from "../components/table_bidang_petugas_gambar";
import TableBidangPetugasUkur from "../components/table_bidang_petugas_ukur";
import TableBidangProblem from "../components/table_bidang_problem";

export default {
  name: "home_detail_berkas",
  components: {
    CardPetugasGambar,
    CardPetugasUkur,
    CardProblem,
    TableBidangPetugasGambar,
    TableBidangPetugasUkur,
    TableBidangProblem,
  },
  data: () => ({
    modal_bidang: false,
    bidangs: [],
    no_berkas: null,
    id: null,
    item: {},
  }),

  methods: {
    bidang(item) {
      this.bidangs = item.bidangs;
      this.modal_bidang = true;
      this.no_berkas = item.no_berkas;
      this.id = item.id;
    },
    detail(item) {
      console.log(item);
    },
  },
  mounted() {
    const id = this.$route.params.id;
    const item = this.$store.getters["home/item"](parseInt(id));
    if (item != null || item != undefined) {
      this.item = item;
    }
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