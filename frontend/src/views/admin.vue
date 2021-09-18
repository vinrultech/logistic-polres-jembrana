<template>
  <v-row>
    <template v-for="item in items">
      <v-col cols="12" md="4" xs="12" sm="12" lg="4" xl="4" :key="item.text">
        <v-card class="text-center pb-4">
          <v-icon size="100">{{ item.icon }}</v-icon>
          <br />
          <v-btn color="primary" class="mt-4" @click="to(item.href)">{{ item.text }}</v-btn>
        </v-card>
      </v-col>
    </template>
  </v-row>
</template>

<script>
export default {
  name: "admin",
  data: () => ({
    items: [
      { icon: "fas fa-users", text: "User", href: "/admin/user", size: 18 },
      {
        icon: "fas fa-id-card-alt",
        text: "Berkas",
        href: "/admin/berkas",
        size: 20,
      },
      {
        icon: "fas fa-person-booth",
        text: "Petugas Ukur",
        href: "/admin/petugas_ukur",
        size: 20,
      },
      {
        icon: "fas fa-chalkboard-teacher",
        text: "Petugas Gambar",
        href: "/admin/petugas_gambar",
        size: 18,
      },
      {
        icon: "fas fa-exclamation-triangle",
        text: "Problem",
        href: "/admin/problem",
        size: 20,
      },
    ],
  }),
  mounted() {
    this.$store.dispatch("constant/menu", 0);
  },
  methods: {
    test() {
      //this.$store.dispatch('user/test', { vm : this});
    },
    to(href, menu) {
      //console.log(href);
      //console.log(this.$router.history.current.path)
      if (this.$router.history.current.path !== href) {
        this.$router.push(href);
        if (menu !== undefined) {
          this.$store.dispatch("constant/menu", menu);
        }
      }

      //this.$store.dispatch("constant/menu", menu);
    },
  },

};
</script>
