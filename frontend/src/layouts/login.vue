<template>
  <v-app class="bg" id="inspire">
    <v-main>
      <v-container class="fill-height" fluid>
        <slot />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  props: {
    source: String,
  },
  data: () => ({
    drawer: null,
  }),
};
</script>

<style scoped>
  .bg {
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    /*background-color: #66BB6A; */
    
  }
</style>