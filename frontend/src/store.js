import Vue from 'vue'
import Vuex from 'vuex'

import user from './store/user'
import kategori from './store/kategori'
import surat_masuk from './store/surat_masuk'
import home from './store/home'
import constant from './store/constant';

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        user,
        kategori,
        surat_masuk,
        home,
        constant,
    }
});