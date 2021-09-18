<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col md="6" lg="6" xl="12" sm="12" cols="6">
        <v-card class="elevation-6 rounded-lg">
          <v-card-text>
            <v-img :src="`${host}${qrcode}`"></v-img>
            <v-btn
              block
              color="primary"
              class="mr-4 rounded-lg"
              @click="login_wa"
            >
              <v-icon dark left>fab fa-whatsapp</v-icon>Login WA
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import utils from "../../utils";
import Breadcum from "../../components/breadcum";
import { SETTINGS } from "../../breadcum";
import eventBus from "../../eventbus";

export default {
  name: "wa",
  components: {
    Breadcum,
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(SETTINGS),
    qrcode: null,
    host: vm.$host,
  }),

  mounted() {
    eventBus.$on("wa_login", (qrcode) => {
      this.qrcode = qrcode;
    });

    eventBus.$on("wa_success", () => {
      this.$toastr.s("Berhasil login wa");
    });
  },
  methods: {
    async login_wa() {
      try {
            
            const token = utils.session();
            await this.$axios.get(`/admin/wa/login`, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            });
        } catch (ex) {
            console.log(ex);
        }
    },
  },
};
</script>

<style scoped>
</style>