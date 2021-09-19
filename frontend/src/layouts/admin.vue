<template>
  <v-app id="keep">
    <v-app-bar app color="brown darken-1">
      <v-app-bar-nav-icon
        @click="drawer = !drawer"
        class="d-lg-none d-xl-none d-xs-block d-sm-block d-md-block white--text"
      ></v-app-bar-nav-icon>
      <v-app-bar-nav-icon
        @click="mini = !mini"
        class="d-none d-xl-block d-lg-block white--text"
      ></v-app-bar-nav-icon>
      <span class="title ml-3 mr-5 white--text">
        Aplikasi&nbsp;
        <span class="font-weight-light">Inventory dan Surat</span>
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
            <v-btn color="primary" text @click="signOut()"
              >Sign Out<v-icon right>fas fa-sign-out-alt</v-icon></v-btn
            >
          </v-card-actions>
        </v-card>
      </v-menu>
    </v-app-bar>

    <v-navigation-drawer
      v-model="drawer"
      app
      color="brown darken-1"
      :mini-variant="mini"
    >
      <v-list class="text-center position-fixed brown darken-1">
        <img
          v-bind:class="{ 'logo-mini': mini, logo: !mini }"
          src="@/assets/logo.svg"
        />
        <v-list-item v-show="!mini">
          <v-list-item-content>
            <v-list-item-title class="title white--text"
              >POLRES JEMBRANA</v-list-item-title
            >
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list
        dense
        class="brown darken-1"
        v-bind:class="{ 'position-under-mini': mini, 'position-under': !mini }"
      >
        <v-list-item-group v-model="selectedItem" color="white">
          <template v-for="(item, i) in items">
            <v-row
              v-if="item.heading"
              :key="item.heading"
              align="center"
              v-show="!mini"
            >
              <v-col cols="6">
                <v-subheader class="white--text" v-if="item.heading">{{
                  item.heading
                }}</v-subheader>
              </v-col>
            </v-row>
            <v-divider
              v-else-if="item.divider"
              :key="i"
              light
              class="mx-2 white"
            ></v-divider>
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
            <div v-else :key="item.text">
              <v-list-item v-if="!mini" @click="to(item.href)" link>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <v-list-item-icon>
                  <v-icon left :size="item.size" color="white">{{
                    item.icon
                  }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title class="white--text">{{
                    item.text
                  }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-tooltip right v-else-if="mini">
                <template v-slot:activator="{ on, attrs }">
                  <v-list-item
                    @click="to(item.href)"
                    link
                    v-bind="attrs"
                    v-on="on"
                  >
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <v-list-item-icon>
                      <v-icon left :size="item.size" color="white">{{
                        item.icon
                      }}</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                      <v-list-item-title class="white--text">{{
                        item.text
                      }}</v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                </template>
                <span>{{ item.text }}</span>
              </v-tooltip>
            </div>
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
      { heading: "Master Data" },
      {
        icon: "fas fa-cookie-bite",
        text: "Kategori",
        href: "/admin/kategori",
        size: 20,
      },
      { heading: "Surat" },
      {
        icon: "fas fa-envelope",
        text: "Surat Masuk",
        href: "/admin/surat_masuk",
        size: 20,
      },
      {
        icon: "fas fa-envelope-open-text",
        text: "Surat Keluar",
        href: "/admin/surat_keluar",
        size: 20,
      },
      {
        icon: "fas fa-mail-bulk",
        text: "Juknis",
        href: "/admin/juknis",
        size: 17,
      },
      { heading: "Inventory" },
      {
        icon: "fas fa-shopping-basket",
        text: "Barang",
        href: "/admin/inventory",
        size: 16,
      },
      {
        icon: "fas fa-warehouse",
        text: "Asset Tetap",
        href: "/admin/asset_tetap",
        size: 16,
      },
      { heading: "Tranksasi" },
      {
        icon: "fas fa-database",
        text: "Inventory Masuk",
        href: "/admin/inventory_masuk",
        size: 16,
      },
      {
        icon: "fas fa-table",
        text: "Inventory Keluar",
        href: "/admin/inventory_keluar",
        size: 16,
      },
      { divider: true },
      { icon: "fas fa-users", text: "User", href: "/admin/user", size: 16 },
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
      if (this.$router.history.current.path !== href) {
        this.$router.push(href);
        if (menu !== undefined) {
          this.$store.dispatch("constant/menu", menu);
        }
      }
    },
  },
  mounted() {
    this.$store.dispatch("user/nama");
    this.$store.dispatch("user/foto");

    const path = this.$router.history.current.path;

    const routes = path.split("/");
    if (routes[2] === "kategori") {
      this.$store.dispatch("constant/menu", 1);
    } else if (routes[2] === "surat_masuk") {
      this.$store.dispatch("constant/menu", 2);
    } else if (routes[2] === "surat_keluar") {
      this.$store.dispatch("constant/menu", 3);
    } else if (routes[2] === "juknis") {
      this.$store.dispatch("constant/menu", 4);
    } else if (routes[2] === "inventory") {
      this.$store.dispatch("constant/menu", 5);
    } else if (routes[2] === "asset_tetap") {
      this.$store.dispatch("constant/menu", 6);
    } else if (routes[2] === "inventory_masuk") {
      this.$store.dispatch("constant/menu", 7);
    } else if (routes[2] === "inventory_keluar") {
      this.$store.dispatch("constant/menu", 8);
    } else if (routes[2] === "user") {
      this.$store.dispatch("constant/menu", 9);
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

.logo-mini {
  width: 45px;
  height: auto;
}

.position-fixed {
  position: fixed !important;
  z-index: 99;
  width: 100%;
}

.position-under {
  margin-top: 132px;
  z-index: 98;
}

.position-under-mini {
  margin-top: 80px;
  z-index: 98;
}
</style>