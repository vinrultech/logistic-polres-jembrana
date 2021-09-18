<template>
  <v-dialog
    v-model="visible"
    persistent
    max-width="500px"
    :retain-focus="false"
  >
    <v-card>
      <v-card-title>
        <v-icon left>fas fa-file-invoice</v-icon>{{ no_bidang }}
        <v-spacer></v-spacer>
        <v-tooltip bottom>
          <template v-slot:activator="{ on }">
            <v-btn x-small fab dark color="error" @click="tutup" v-on="on">
              <v-icon dark size="15">fas fa-times</v-icon>
            </v-btn>
          </template>
          <span>Tutup</span>
        </v-tooltip>
      </v-card-title>
      <v-card-text style="height: 250px">
        <v-form ref="form" v-model="valid">
          <v-textarea
            v-model="keterangan"
            clearable
            clear-icon="fas fa-times"
            label="Keterangan"
            :rules="[(v) => !!v || 'Keterangan dibutuhkan']"
          ></v-textarea>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn :disabled="!valid" color="primary" @click="simpan">
          <v-icon dark left>fas fa-save</v-icon>Simpan
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    no_bidang: {
      required: true,
    },
    bidang_id: {
      required: true,
    },
    berkas_id: {
      required: true,
    },
  },

  data: () => ({
    valid: true,
    keterangan: "",
  }),
  methods: {
    tutup() {
      this.$emit("tutup");
    },
    simpan() {
      this.$store
        .dispatch("petugas_gambar/proses_masalah", {
          item: {
            bidang_id: parseInt(this.bidang_id),
            berkas_id: parseInt(this.berkas_id),
            keterangan: this.keterangan
          },
          vm: this,
        })
        .then(() => {
          this.$emit("tutup", true);
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

