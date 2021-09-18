<template>
  <v-dialog v-model="visible" max-width="500px" :persistent="true">
    <v-card>
      <v-card-title>
        <span>{{ title }}</span>
      </v-card-title>
      <v-card-text>
        <v-progress-circular
          indeterminate
          :size="70"
          :width="5"
          color="purple"
          style="vertical-align: middle"
          class="mr-4"
        ></v-progress-circular>
        <span>{{ text }}</span>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
import Progress from "../plugins/progress_dialog";
export default {
  data: () => ({
    title: null,
    text: null,
    visible: false,
  }),
  methods: {
    hide() {
      this.visible = false;
    },
    show(params) {
      this.visible = true;
      this.title = params.title;
      this.text = params.text;
    },
  },
  beforeMount() {
    Progress.EventBus.$on("progress.show", (params) => {
      this.show(params);
    });
    Progress.EventBus.$on("progress.hide", (/* params */) => {
      this.hide();
    });
  },
};
</script>

<style scoped>
</style>

