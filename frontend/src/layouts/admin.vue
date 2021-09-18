<template>
  <v-app id="keep">
    <v-app-bar app color="green darken-3">
      <v-app-bar-nav-icon
        @click="drawer = !drawer"
        class="d-lg-none d-xl-none d-xs-block d-sm-block d-md-block"
      ></v-app-bar-nav-icon>
      <v-app-bar-nav-icon
        @click="mini = !mini"
        class="d-none d-xl-block d-lg-block"
      ></v-app-bar-nav-icon>
      <span class="title ml-3 mr-5">
        Aplikasi&nbsp;
        <span class="font-weight-light">Berkas Tracker</span>
      </span>
      <v-spacer />
      <v-menu bottom left>
        <template v-slot:activator="{ on }">
          <v-btn icon v-on="on">
            <v-avatar class="white" size="40">
              <v-icon v-if="foto === '' || foto == 'null'" color="primary"
                >fas fa-user-shield</v-icon
              >
              <img v-else :src="`${host}upload/${foto}`" />
            </v-avatar>
          </v-btn>
        </template>
        <v-card tile>
          <v-list rounded>
            <v-subheader>{{ nama }}</v-subheader>
            <v-list-item-group color="primary">
              <v-list-item @click="to(`/admin/change`, 0)">
                <v-list-item-icon>
                  <v-icon>fas fa-user-lock</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>Ubah Password</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-list-item @click="to(`/admin/account`, 0)">
                <v-list-item-icon>
                  <v-icon>fas fa-user-cog</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>Ubah Account</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="signOut()">Sign Out<v-icon right>fas fa-sign-out-alt</v-icon></v-btn>
          </v-card-actions>
        </v-card>
      </v-menu>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      app
      color="green darken-2"
      :mini-variant="mini"
    >
      <v-list-item>
        <v-list-item-avatar>
          <v-img :src="logo"></v-img>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title class="title"
            >Badan Pertanahan Nasional</v-list-item-title
          >
          <v-list-item-subtitle>Kabupaten Jembrana</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>
      <v-list dense class="green darken-1">
        <v-list-item-group v-model="selectedItem" color="white">
          <template v-for="item in items">
            <v-row v-if="item.heading" :key="item.heading" align="center">
              <v-col cols="6">
                <v-subheader v-if="item.heading">{{
                  item.heading
                }}</v-subheader>
              </v-col>
              <v-col cols="6" class="text-center">
                <a href="#!" class="body-2 black--text">EDIT</a>
              </v-col>
            </v-row>
            <v-list-group
              v-else-if="item.children"
              :key="item.text"
              v-model="item.model"
              :append-icon="item.model ? item.icon : item['icon-alt']"
              prepend-icon
            >
              <template v-slot:activator>
                <v-list-item>
                  <v-list-item-avatar v-if="item['icon-main']" left size="20">
                    <v-icon left size="20">{{ item["icon-main"] }}</v-icon>
                  </v-list-item-avatar>
                  <v-list-item-content>
                    <v-list-item-title>{{ item.text }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </template>
              <v-list-item
                v-for="(child, i) in item.children"
                :key="i"
                @click="to(child.href)"
              >
                <!-- <v-list-item-icon v-if="child.icon">
              </v-list-item-icon>-->
                &nbsp;&nbsp;&nbsp;&nbsp;
                <v-list-item-avatar v-if="child.icon" size="20">
                  <v-icon left size="20">{{ child.icon }}</v-icon>
                </v-list-item-avatar>
                <v-list-item-content>
                  <v-list-item-title>{{ child.text }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-group>
            <v-list-item v-else :key="item.text" @click="to(item.href)" link>
              &nbsp;&nbsp;&nbsp;&nbsp;
              <v-list-item-icon>
                <v-icon left :size="item.size">{{ item.icon }}</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ item.text }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </template>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container fluid>
        <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import logoPng from "../assets/logo.svg";
import utils from "../utils";
import { mapGetters } from "vuex";

export default {
  props: {
    source: String,
  },
  data: (vm) => ({
    drawer: null,
    mini: false,
    logo: logoPng,
    items: [
      { icon: "fas fa-home", text: "Dashboard", href: "/admin", size: 20 },
      
      { icon: "fas fa-users", text: "User", href: "/admin/user", size: 18 },
      /*
      { icon: "fas fa-id-card-alt", text: "Berkas", href: "/admin/berkas", size: 20 },
      {
        icon: "fas fa-person-booth",
        text: "Petugas Ukur",
        href: "/admin/petugas_ukur",
        size: 20
      },
      {
        icon: "fas fa-chalkboard-teacher",
        text: "Petugas Gambar",
        href: "/admin/petugas_gambar",
        size: 18
      },
      {
        icon: "fas fa-exclamation-triangle",
        text: "Problem",
        href: "/admin/problem",
        size: 20
      },
      */
      /*
      {
        icon: "fas fa-cogs",
        text: "Settings",
        href: "/admin/settings",
        size: 18
      },
      */
    ],
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      nama: "user/account",
      foto: "user/foto",
    }),
    selectedItem: {
      get() {
        return this.$store.getters["constant/menu"];
      },
      set(val) {
        this.$store.dispatch("constant/menu", val);
      },
    },
  },
  methods: {
    signOut() {
      utils.removeSession();
      this.$store.dispatch("constant/menu", 0);
      this.$router.push("/login");
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
  mounted() {
    this.$store.dispatch("user/nama");
    this.$store.dispatch("user/foto");
    //console.log(this.$router.history.current.path);

    const path = this.$router.history.current.path;

    const routes = path.split("/");

    if (routes[2] === "user") {
      this.$store.dispatch("constant/menu", 1);
    } else if (routes[2] === "berkas") {
      this.$store.dispatch("constant/menu", 2);
    } else if (routes[2] === "petugas_ukur") {
      this.$store.dispatch("constant/menu", 3);
    } else if (routes[2] === "petugas_gambar") {
      this.$store.dispatch("constant/menu", 4);
    } else if (routes[2] === "problem") {
      this.$store.dispatch("constant/menu", 5);
    } else {
      this.$store.dispatch("constant/menu", 0);
    }
  },
};
</script>

<style>
#keep .v-navigation-drawer__border {
  display: none;
}

.company {
  height: 200px;
  background-color: #607d8b;
  color: white;
}
.logo {
  width: 54px;
  height: auto;
}
</style>