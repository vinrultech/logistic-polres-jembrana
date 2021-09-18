<template>
  <div id="app">
    <component :is="layout">
      <router-view />
    </component>
    <progress-dialog />
  </div>
</template>

<script>
const default_layout = "notfound";
import utils from "./utils";

export default {
  name: "App",
  computed: {
    layout() {
      return (this.$route.meta.layout || default_layout) + "-layout";
    },
  },
  mounted() {
    const vm = this;

    this.$axios.interceptors.request.use(
      function (config) {
        let title = "Get";
        let text = "Sedang mendapatkan data dari server";
        if (config.url === "user/login") {
          title = "Login";
          text = "Sedang login ke server";
        } else {
          const token = utils.session();
          if (config.method === "post") {
            title = "Tambah";
            text = "Sedang tambah data ke server";
          } else if (config.method === "put") {
            title = "Update";
            text = "Sedang update data ke server";
            if (config.url == "/admin/user/change") {
              title = "Profile";
              text = "Sedang merubah password";
            }
          } else if (config.method === "delete") {
            title = "Hapus";
            text = "Sedang hapus data ke server";
          }
          config.headers = {
            "Content-Type": `application/json`,
            Authorization: `Bearer ${token}`,
          };
        }
        vm.$progress.show({
          title: title,
          text: text,
        });
        return config;
      },
      function (error) {
        return Promise.reject(error);
      }
    );

    this.$axios.interceptors.response.use(
      function (response) {
        vm.$progress.hide();
        if (response.config.message) {
          vm.$toastr.s(response.config.message);
        }
        return response;
      },
      function (error) {
        vm.$progress.hide();
        const code = parseInt(error.response.status);
        if (code === 400 || code === 401 || code === 403) {
          if (vm.$swal !== undefined && vm.$router !== undefined) {
            utils.error(error, vm.$swal, vm.$router);
          }
        } else if (error.response.config.method === "get") {
          if (vm.$swal !== undefined && vm.$router !== undefined) {
            utils.error(error, vm.$swal, vm.$router);
          }
        }

        return Promise.reject(error);
      }
    );
  },
};
</script>
