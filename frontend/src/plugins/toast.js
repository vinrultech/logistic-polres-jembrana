import Vue from 'vue'
// import plugin
import VueToastr from "vue-toastr";
// use plugin
Vue.use(VueToastr, {
    /* OverWrite Plugin Options if you need */
    defaultTimeout: 2000,
    defaultPosition: "toast-top-right",
    defaultStyle: { "background-color": "green" },
});