<template>
  <div class="image-container">
    <v-row>
      <v-col
        cols="6"
        md="3"
        xs="6"
        sm="6"
        lg="3"
        xl="3"
        v-for="file in files"
        :key="file.file_id"
      >
        <a href="#" @click="download(file.url)">{{ file.filename }}</a>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  data: (vm) => ({
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      files: "files/files",
    }),
  },
  mounted() {},
  methods: {
    download(url) {
      window.open(`${this.host}${url}`);
    },
  },
};
</script>

<style scoped>
</style>