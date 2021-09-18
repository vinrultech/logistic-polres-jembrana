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
    <table-bidang-petugas-gambar :items="items" :id="id" @problem="problem" />
    <v-card>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          :color="is_problem ? 'error' : 'success'"
          @click="is_problem ? finish_masalah() : finish()"
          :disabled="!is_valid"
          >{{ is_problem ? 'Update Masalah' : 'Selesai'}}<v-icon right>
            {{ is_problem ? 'fas fa-exclamation-triangle' : 'fas fa-check'}}
          </v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
        
      

    <add-problem
      :berkas_id="id"
      :bidang_id="bidang_id"
      :no_bidang="no_bidang"
      :visible="modal_add_problem"
      @tutup="tutup_add_problem"
    />
  </v-dialog>
</template>

<script>
import AddProblem from "./add_problem";
import TableBidangPetugasGambar from "./table_bidang_petugas_gambar";
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
    AddProblem,
    TableBidangPetugasGambar,
  },
  computed: {
    is_problem() {
      return this.items.filter((el) => el.status_problem == 1).length > 0;
    },
    is_valid() {
      const gambar = this.items.filter((el) => el.status_gambar == 1).length;
      const masalah = this.items.filter((el) => el.status_problem == 1).length;
      return gambar + masalah === this.items.length;
    },
  },
  data: () => ({
    tanggal_setor: null,
    modal_date: false,
    bidang_id: null,
    modal_add_problem: false,
    no_bidang: null,
  }),
  methods: {
    tutup() {
      this.valid = false;
      this.$emit("tutup");
    },
    tutup_add_problem() {
      this.modal_add_problem = false;
      this.bidang_id = null;
      this.no_bidang = null;
    },
    problem(item) {
      this.modal_add_problem = true;
      this.bidang_id = item.id;
      this.no_bidang = item.no_bidang;
    },
    finish() {
      this.valid = false;
      this.$store
        .dispatch("petugas_gambar/finish", {
          item: {
            id: parseInt(this.id),
          },
          vm: this,
        }).then(() => {
           this.$emit("tutup");
        });
    },
    finish_masalah() {
      this.valid = false;
      this.$store
        .dispatch("petugas_gambar/finish_masalah", {
          item: {
            id: parseInt(this.id),
          },
          vm: this,
        }).then(() => {
           this.$emit("tutup");
        });
    },
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>

