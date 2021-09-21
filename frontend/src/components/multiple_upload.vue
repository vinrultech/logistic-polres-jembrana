<template>
  <div class="image-container">
    <input type="file" @change="upload" ref="file" style="display: none" />
    <v-btn @click="$refs.file.click()" color="blue">
      Upload file<v-icon right>fas fa-upload</v-icon>
    </v-btn>
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
        v-show="file.status !== 'delete'"
      >
        {{ file.filename }}
        <v-btn icon color="red" @click="remove(file.file_id, file.filename, file.status)">
          <v-icon>fas fa-times</v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { v1 as uuidv1 } from "uuid";

export default {
  props: {
    status: {
      type: String,
      required: true,
    },
  },
  data: (vm) => ({
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      files: "files/files",
    }),
  },
  mounted() {
    if (this.status === "add") {
      this.$store.dispatch("files/reset");
    }
  },
  methods: {
    upload(e) {
      const file = e.target.files[0];
      const filename = e.target.files[0].name;
      const formData = new FormData();
      formData.append("file", file);
      formData.append("file_id", uuidv1())
      formData.append("filename", filename)
      formData.append("status", "new")

      this.$store.dispatch("files/upload", formData)
      //console.log(filename)
      /*
      const item = {
        file_id: uuidv1(),
        url: filename,
        filename: filename,
        status: "new",
      };

      this.$store.commit("files/ADD_FILE", item);
      */
    },
    remove(file_id, filename, status) {
      //this.$store.commit("files/REMOVE_FILE", file_id);
      const item = {
        file_id: file_id,
        filename: filename,
        status: status
      }

      this.$store.dispatch("files/remove", item)

    },
  },
};
</script>

<style scoped>
</style>