<template>
  <v-breadcrumbs :items="items" class="grey lighten-4" large>
    <template v-slot:item="props">
      <v-breadcrumbs-item
        @click="open_breadcum(props.item.href, props.item.menu)"
        :class="{
          'breadcums-disabled': props.item.disabled,
          'breadcums-enabled': !props.item.disabled,
        }"
      >
        <v-icon left>{{ props.item.icon }}</v-icon>
        {{ breadcumText(props.item.text) }}
      </v-breadcrumbs-item>
    </template>
    <template v-slot:divider>
      <v-icon>fas fa-chevron-right</v-icon>
    </template>
  </v-breadcrumbs>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      required: true,
    },
  },
  data: () => ({}),
  methods: {
    open_breadcum(href, menu) {
      this.$router.push(href);
      this.$store.dispatch("constant/menu", menu);
    },
  },
};
</script>

<style scoped>
</style>

