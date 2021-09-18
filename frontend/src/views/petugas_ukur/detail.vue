<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <card-petugas-ukur
          :item="item"
          :is_detail="false"
          @bidang="bidang"
          @detail="detail"
        />
      </v-col>
    </v-row>
    <bidang-petugas-ukur
      :no_berkas="no_berkas"
      :id="id"
      :visible="modal_bidang"
      :items="bidangs"
      @tutup="modal_bidang = false"
    />
  </div>
</template>

<script>
import BidangPetugasUkur from "../../components/bidang_petugas_ukur";
import CardPetugasUkur from "../../components/card_petugas_ukur";
import Breadcum from "../../components/breadcum";
import { PETUGAS_UKUR, INFO } from "../../breadcum";
import utils from "../../utils";

export default {
  name: "detail_petugas_ukur",
  components: {
    BidangPetugasUkur,
    CardPetugasUkur,
    Breadcum,
  },
  data: () => ({
    breadcums: utils.breadcumTwo(PETUGAS_UKUR(false), INFO),
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
    modal_bidang: false,
    bidangs: [],
    no_berkas: null,
    id: null,
    item: {
      bidangs : []
    },
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
    const item = this.$store.getters["petugas_ukur/item"](parseInt(id));
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