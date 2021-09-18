import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import id from 'vuetify/es5/locale/id';

Vue.use(Vuetify);

export default new Vuetify({
    lang: {
        locales: {
            id
        },
        current: 'id',
    },
    icons: {
        iconfont: 'fa',
    },
});