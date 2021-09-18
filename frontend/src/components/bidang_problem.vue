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
    <table-bidang-problem :items="items" :id="id" />
    <v-card>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="success" @click="finish()" :disabled="!is_valid"
          >Selesai<v-icon right> fas fa-check </v-icon>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import TableBidangProblem from "./table_bidang_problem";
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
    TableBidangProblem,
  },
  computed: {
    is_valid() {
      return (
        this.items.filter((el) => el.status_problem == 0).length ===
        this.items.length
      );
    },
  },
  data: () => ({
  }),
  methods: {
    tutup() {
      this.$emit("tutup");
    },
    finish() {
      this.$store
        .dispatch("problem/finish", {
          item: {
            id: parseInt(this.id),
          },
          vm: this,
        })
        .then(() => {
          this.$emit("tutup");
        });
    }
  },
};
</script>

<style scoped>
.action {
  width: 130px;
}
</style>

