<template>
  <v-row v-resize="onResize">
    <v-col cols="12" xl="6" lg="6" md="6" sm="12" xs="12">
      <chart
        :items="ukurs"
        id="petugas_ukur"
        :width="width"
        :height="height"
        title="Bidang yang belum di ukur"
        x_label="Petugas Ukur"
        y_label="Bidang"
        @pilih="select_ukur"
      />
    </v-col>
    <v-col cols="12" xl="6" lg="6" md="6" sm="12" xs="12">
      <chart
        :items="gambars"
        id="petugas_gambar"
        :width="width"
        :height="height"
        title="Bidang yang belum di gambar"
        x_label="Petugas Gambar"
        y_label="Bidang"
        @pilih="select_gambar"
      />
    </v-col>
  </v-row>
</template>

<script>
import Chart from "../components/chart";

import eventbus from "../eventbus";

export default {
  name: "home",
  components: {
    Chart,
  },
  data: () => ({
    width: 720,
    height: 500,
  }),
  computed: {
    ukurs() {
      return this.$store.getters["home/ukurs"];
    },
    gambars() {
      return this.$store.getters["home/gambars"];
    },
  },
  mounted() {
    this.onResize();
    this.dispatchUkur();
    this.dispatchGambar();
    eventbus.$on('update', (payload) => {
      if (payload.action === "ukur") {
        this.dispatchUkur();
      } else if (payload.action === "gambar") {
        this.dispatchGambar();
      } else {
        this.dispatchUkur();
        this.dispatchGambar();
      }
    });
  },
  methods: {
    onResize() {
      this.width = window.innerWidth / 2;
      this.height = window.innerHeight - 120;
    },
    select_gambar(item) {
      this.$router.push(`/petugas_gambar/${item.id}`);
    },
    select_ukur(item) {
      this.$router.push(`/petugas_ukur/${item.id}`);
    },
    dispatchUkur() {
      this.$store.dispatch("home/chart_petugas_ukur", {
        vm: this,
      });
    },
    dispatchGambar() {
      this.$store.dispatch("home/chart_petugas_gambar", {
        vm: this,
      });
    },
  },
};
</script>

<style scoped>
</style>
