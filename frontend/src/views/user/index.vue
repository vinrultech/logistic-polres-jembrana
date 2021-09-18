<template>
  <div>
    <v-row>
      <v-col cols="12">
        <breadcum :items="breadcums" />
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
        <v-btn rounded dark color="green" @click="add()" class="mt-5">
          <v-icon dark left>fas fa-plus</v-icon>Tambah
        </v-btn>
        <v-btn rounded dark color="deep-purple" @click="refresh()" class="mt-5">
          <v-icon dark left>fas fa-sync-alt</v-icon>Refresh
        </v-btn>
      </v-col>
      <v-col cols="12" md="6" xs="12" sm="12" lg="6" xl="6">
        <v-row>
          <v-col cols="4">
            <v-select
              v-model="filterCari"
              :items="filters"
              menu-props="auto"
              item-text="nama"
              item-value="id"
              label="Filter"
              hide-details
              prepend-icon="fas fa-filter"
              single-line
            ></v-select>
          </v-col>
          <v-col cols="8">
            <v-text-field
              v-model="searchText"
              :placeholder="'Cari'"
              rounded
              solo
              prepend-inner-icon="fas fa-search"
              :append-icon="isSearch ? 'fas fa-times' : ''"
              single-line
              @keydown.enter="cari()"
              @click:append="refresh()"
              @click:prepend-inner="refresh()"
            ></v-text-field>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    
    <v-card class="elevation-6">
      <v-card-text>
        <v-simple-table>
          <thead>
            <tr>
              <th class="text-center">Nama</th>
              <th class="text-center">Username</th>
              <th class="text-center">Role</th>
              <th class="text-center">No WA</th>
              <th class="text-center">Foto</th>
              <th class="text-center action">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in items" :key="item.id">
              <td>
                {{ item.nama | str_limit(20) }}
              </td>
              <td>
                {{ item.username }}
              </td>
              <td>
                {{ item.role }}
              </td>
              <td>
                {{ item.hp }}
              </td>
              <td class="text-center">
                <v-avatar size="48">
                  <img :src="`${host}upload/${item.foto}`" v-if="item.foto !== ''"/>
                  <v-icon size="46" v-else>fas fa-user</v-icon>
                </v-avatar>
              </td>
              <td class="text-center">
                <v-tooltip left>
                  <template v-slot:activator="{ on }">
                    <v-btn
                      x-small
                      fab
                      dark
                      color="error"
                      v-if="item.username !== username"
                      v-on="on"
                      @click="remove(item.id)"
                    >
                      <v-icon dark size="15">fas fa-trash</v-icon>
                    </v-btn>
                  </template>
                  <span>Hapus</span>
                </v-tooltip>
                <v-tooltip top>
                  <template v-slot:activator="{ on }">
                    <v-btn
                      x-small
                      fab
                      dark
                      color="primary"
                      v-if="item.username !== username"
                      @click="edit(item.id)"
                      v-on="on"
                    >
                      <v-icon dark size="15">fas fa-edit</v-icon>
                    </v-btn>
                  </template>
                  <span>Edit</span>
                </v-tooltip>
              </td>
            </tr>
          </tbody>
        </v-simple-table>
      </v-card-text>
    </v-card>
    <paging
      :limits="limits"
      :limit="limit"
      :prev-show="prev_show"
      :next-show="next_show"
      @next="next"
      @previous="previous"
      @changeLimit="changeLimit"
    />
  </div>
</template>

<script>
import Paging from "../../components/paging";
import Breadcum from "../../components/breadcum";
import { USER } from "../../breadcum";
import utils from "../../utils";
import { mapGetters } from "vuex";
export default {
  name: "user",
  components: {
    Paging,
    Breadcum
  },
  data: (vm) => ({
    breadcums: utils.breadcumOne(USER(true)),
    filters: [
      {
        id: "nama",
        nama: "Nama",
      },
      {
        id: "username",
        nama: "Username",
      },
    ],
    filterCari: "nama",
    username: utils.username(),
    searchText: "",
    isSearch: false,
    host: vm.$host,
  }),
  computed: {
    ...mapGetters({
      items: "user/items",
      prev_show: "user/prev_show",
      next_show: "user/next_show",
      first_no: "user/first_no",
      last_no: "user/last_no",
      limit: "user/limit",
      limits: "constant/limits",
    }),
  },
  watch: {
    limit: function () {
      this.get("first");
    },
    filterCari: function () {
      this.get("first");
    },
    searchText: function (val) {
      if (!(!val || /^\s*$/.test(val))) {
        this.isSearch = true;
      } else {
        this.isSearch = false;
        this.get("first");
      }
    },
  },
  methods: {
    changeSearchText(val) {
      this.searchText = val;
    },
    changeLimit(val) {
      const limit = {
        value: val,
        text: val.toString(),
      };
      this.$store.dispatch("user/limit", limit);
    },
    add() {
      this.$router.push("/admin/user/create");
    },
    edit(id) {
      this.$router.push(`/admin/user/edit/${id}`);
    },
    find_items(id) {
      return this.items.find((item) => {
        return item.id === id;
      });
    },
    async remove(id) {
      const item = this.find_items(id);

      if (item.is_delete === 1) {
        this.$swal({
          title: "Peringatan",
          text: `${item.pegawai.jabatan} tidak bisa di hapus`,
          type: "warning",
        });

        return;
      }
      this.$swal({
        title: "Anda yakin?",
        text: "Apakah anda ingin menhapus data!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya",
        cancelButtonText: "Tidak",
      }).then((result) => {
        if (result.value) {
          this.$store.dispatch("user/remove", { vm: this, id: id }).then(() => {
            this.get("first");
          });
        }
      });
    },
    get(action) {
      let no = 0;
      if (action === "next") {
        no = this.last_no;
      } else if (action === "prev") {
        no = this.first_no;
      } else {
        no = 0;
      }
      let query = {
        limit: this.limit.value,
        action: action,
        no: no,
        vm: this,
      };
      if (
        !(
          !this.searchText || /^\s*$/.test(this.searchText)
        ) /*this.search !== "" || this.search !== null*/
      ) {
        query = {
          ...query,
          search: this.searchText,
          filter: this.filterCari,
        };
      }
      this.$store.dispatch("user/paginator", query);
    },
    previous() {
      this.get("prev");
    },
    next() {
      this.get("next");
    },
    cari() {
      this.get("first");
    },
    refresh() {
      this.searchText = "";
      //this.get("first");
    },
  },
  mounted() {
    this.get("first");
  },
};
</script>

<style scoped>
.action {
  width: 100px;
}
</style>