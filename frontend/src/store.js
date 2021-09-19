import Vue from 'vue'
import Vuex from 'vuex'

import user from './store/user'
import kategori from './store/kategori'
import inventory from './store/inventory'
import asset_tetap from './store/asset_tetap'
import surat_masuk from './store/surat_masuk'
import surat_keluar from './store/surat_keluar'
import juknis from './store/juknis'
import home from './store/home'
import constant from './store/constant';

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        user,
        kategori,
        surat_masuk,
        surat_keluar,
        juknis,
        inventory,
        asset_tetap,
        home,
        constant,
    }
});