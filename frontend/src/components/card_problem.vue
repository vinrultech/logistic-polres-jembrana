<template>
  <v-card class="elevation-6">
    <v-card-title>
      {{ item.no_berkas }}
      <v-spacer></v-spacer>
      <v-avatar color="primary" size="48">
        <img :src="`${host}upload/${item.foto_petugas_gambar}`" />
      </v-avatar>
    </v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="12" xl="4" lg="4" md="4" sm="12" xs="12">
          Pemohon :
        </v-col>
        <v-col cols="12" xl="8" lg="8" md="8" sm="12" xs="12">
          {{ item.pemohon }} / {{ item.wa }}
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" xl="4" lg="4" md="4" sm="12" xs="12">
          Jumlah Bidang :
        </v-col>
        <v-col cols="12" xl="8" lg="8" md="8" sm="12" xs="12">
          {{ item.bidangs.length }}
        </v-col>
      </v-row>
      <v-stepper alt-labels>
        <v-stepper-header>
          <v-stepper-step step="1" complete color="success">
            <div class="text-center text-step">
              Input<br />
              {{ item.tanggal_input }}
            </div>
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step complete step="2" color="success">
            <div class="text-center text-step">Ukur</div>
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step step="3">
            <div class="text-center text-step">
              Gambar<br />
              <span v-if="item.nama_petugas_gambar">{{
                item.nama_petugas_gambar
              }}</span
              ><span v-else>{{ item.petugas }}</span
              ><br />
              <div
                v-if="
                  item.bidangs.filter((el) => el.status_gambar === 1).length > 0
                "
              >
                <v-icon color="success" size="20">fas fa-user-check</v-icon
                ><span class="done">
                  {{
                    item.bidangs.filter((el) => el.status_gambar === 1).length
                  }}
                  bidang</span
                >
              </div>
              <div
                v-if="
                  item.bidangs.filter(
                    (el) => el.status_gambar === 0 && el.status_problem === 0
                  ).length > 0
                "
              >
                <v-icon color="error" size="20">fas fa-user-times</v-icon
                ><span class="undone">
                  {{
                    item.bidangs.filter(
                      (el) => el.status_gambar === 0 && el.status_problem === 0
                    ).length
                  }}
                  bidang</span
                >
              </div>
            </div>
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step :rules="[() => false]" step="4">
            Problem
            <small
              v-if="
                item.bidangs.filter(
                  (el) => el.status_problem === 1 && el.status_gambar === 0
                ).length > 0
              "
            >
              <v-icon color="error" size="16"
                >fas fa-exclamation-triangle</v-icon
              >

              <span class="undone">
                {{
                  item.bidangs.filter(
                    (el) => el.status_problem === 1 && el.status_gambar === 0
                  ).length
                }}
                bidang</span
              >
            </small>
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step step="5" v-if="!is_detail">
            <div class="text-center text-step">Selesai</div>
          </v-stepper-step>
        </v-stepper-header>
      </v-stepper>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <div v-if="!is_detail">
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-btn x-small fab dark color="error" @click="tutup" v-on="on">
              <v-icon dark size="15">fas fa-times</v-icon>
            </v-btn>
          </template>
          <span>Kembali</span>
        </v-tooltip>
      </div>
      <div v-if="is_detail">
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-btn
              x-small
              fab
              dark
              color="indigo"
              @click="detail(item)"
              v-on="on"
            >
              <v-icon dark size="15">fas fa-info</v-icon>
            </v-btn>
          </template>
          <span>Detail</span>
        </v-tooltip>
      </div>
      <div v-if="!is_home">
        <v-tooltip top>
          <template v-slot:activator="{ on }">
            <v-btn
              x-small
              fab
              dark
              color="teal"
              @click="bidang(item)"
              v-on="on"
            >
              <v-icon dark size="15">fas fa-file-invoice</v-icon>
            </v-btn>
          </template>
          <span>Bidang</span>
        </v-tooltip>
      </div>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
    is_detail: {
      type: Boolean,
      required: true,
    },
    is_home: {
      type: Boolean,
      default: false,
    },
  },
  data: (vm) => ({
    host: vm.$host,
  }),
  methods: {
    tutup() {
      this.$router.go(-1);
    },
    bidang(item) {
      this.$emit("bidang", item);
    },
    detail(item) {
      this.$emit("detail", item);
    },
  },
};
</script>

<style scoped>
</style>

