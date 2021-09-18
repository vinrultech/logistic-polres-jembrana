<template>
  <v-row>
    <v-col cols="12">
      <div class="d-flex flex-row float-right">
        <v-select
          v-model="limitData"
          :items="limits"
          item-value="value"
          item-text="text"
          label="Rows per page"
          single-line
          class="mr-2 limits"
          dense
          small-chips
          flat
        />
        <v-btn
          @click="previous()"
          :disabled="!prevShow"
          class="primary"
          fab
          small
        >
          <v-icon>fas fa-arrow-left</v-icon>
        </v-btn>
        <v-btn @click="next()" :disabled="!nextShow" class="primary" fab small>
          <v-icon>fas fa-arrow-right</v-icon>
        </v-btn>
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  props: {
    limits: {
      type: Array,
      required: true,
    },
    limit: {
      type: Object,
      required: true,
    },
    prevShow: {
      type: Boolean,
      required: true,
    },
    nextShow: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    limitData: {
      get() {
        return this.limit;
      },
      set(val) {
        this.$emit("changeLimit", val);
      },
    },
  },
  methods: {
    previous() {
      this.$emit("previous");
    },
    next() {
      this.$emit("next");
    },
  },
};
</script>

<style scoped>
.limits {
  max-width: 80px;
}
</style>
