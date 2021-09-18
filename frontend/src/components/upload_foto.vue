<template>
  <div class="image-container">
    <img :src="`${host}upload/${image}`" v-if="image !== ''" />
    <img
      class="upload-image"
      src="@/assets/upload.png"
      v-else
      @click="$refs.file.click()"
    />
    <v-btn
      icon
      class="btn-hapus"
      v-if="image !== ''"
      color="red"
      @click="remove"
    >
      <v-icon>fas fa-times</v-icon>
    </v-btn>
    <input
      type="file"
      accept="image/*"
      @change="upload"
      ref="file"
      style="display: none"
    />
    <v-btn
      icon
      class="text-center"
      @click="$refs.file.click()"
      v-if="image !== ''"
    >
      <v-icon> fas fa-upload </v-icon>
    </v-btn>
  </div>
</template>

<script>
export default {
  props: {
    image: {
      type: String,
      required: true,
    },
    username: {
      type: String,
      required: true,
    },
  },
  data: (vm) => ({
    host: vm.$host,
  }),
  methods: {
    upload(e) {
      const image = e.target.files[0];
      const formData = new FormData();
      formData.append("image", image);

      this.$store
        .dispatch("user/upload_image", {
          formData,
          username: this.username,
          progress: this.$progress,
          axios: this.$axios,
          swal: this.$swal,
          toast: this.$toastr,
          router: this.$router,
        })
        .then((data) => {
          this.$emit("uploaded", data.filename);
        })
    },
    remove() {
      this.$store
        .dispatch("user/remove_image", {
          item: {
            filename: this.image,
            username: this.username,
          },
          progress: this.$progress,
          axios: this.$axios,
          swal: this.$swal,
          toast: this.$toastr,
          router: this.$router,
        })
        .then(() => {
          this.$emit("removed");
        });
    },
  },
};
</script>

<style scoped>
.image-container {
  position: relative;
  width: 100px;
  height: 150px;
}

.image-container img {
  display: flex;
  width: 100px;
  height: 100px;
  z-index: 99;
}

.image-container .btn-hapus {
  position: absoute;
  top: -65%;
  left: 65%;
  z-index: 100;
}

.image-container .upload-image {
  cursor: pointer;
}
</style>