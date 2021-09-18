<template>
  <v-dialog v-model="visible" persistent max-width="800px">
    <v-toolbar>
      <v-toolbar-title
        ><v-icon left>fas fa-file-invoice</v-icon
        >{{ no_berkas }}</v-toolbar-title
      >
      <v-spacer></v-spacer>
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn x-small fab dark color="error" @click="tutup" v-on="on">
            <v-icon dark size="15">fas fa-times</v-icon>
          </v-btn>
        </template>
        <span>Tutup</span>
      </v-tooltip>
    </v-toolbar>
    <v-stepper v-model="step">
      <v-stepper-header>
        <v-stepper-step :complete="step > 1" step="1">
          Proses Bidang
        </v-stepper-step>
        <v-divider></v-divider>
        <v-stepper-step :complete="step > 2" step="2">
          Pilih Petugas Gambar
        </v-stepper-step>
      </v-stepper-header>
      <v-stepper-items>
        <v-stepper-content step="1">
          <table-bidang-petugas-ukur :items="items" :id="id" />
          <v-btn
            color="success"
            @click="step = 2"
            :disabled="
              !(
                items.filter((el) => el.status_ukur === 1).length ===
                items.length
              )
            "
          >
            Pilih Petugas Gambar<v-icon right>fas fa-chalkboard-teacher</v-icon>
          </v-btn>
        </v-stepper-content>
        <v-stepper-content step="2">
          <list-petugas-gambar @pilih="pilih" />
        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>
  </v-dialog>
</template>

<script>
import ListPetugasGambar from "./list_petugas_gambar";
import TableBidangPetugasUkur from "./table_bidang_petugas_ukur";
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
    visible: {
      type: Boolean,
      required: true,
    },
    no_berkas: {
      required: true,
    },
    id: {
      required: true,
    },
  },
  components: {
    ListPetugasGambar,
    TableBidangPetugasUkur,
  },
  data: () => ({
    step: 1,
    valid: false,
  }),
  methods: {
    tutup() {
      this.step = 1;
      this.valid = false;
      this.$emit("tutup", false);
    },

    pilih(item) {
      this.valid = false;
      this.$store
        .dispatch("petugas_ukur/pilih", {
          item: {
            id: parseInt(this.id),
            petugas_gambar: parseInt(item.id),
          },
          vm: this,
        })
        .then(() => {
          this.step = 1;
          this.$emit("tutup", true);
        });
    },
    checkValid() {
      let valid = 0;
      for (let i = 0; i < this.items.length; i++) {
        const item = this.items[i];
        if (item.status === 1) {
          valid++;
        }
      }

      if (valid === this.items.length) {
        this.valid = true;
      }
    },
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>

