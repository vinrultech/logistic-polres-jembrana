import Vue from 'vue'
import Vuex from 'vuex'

import user from './store/user'
import kategori from './store/kategori'
import unit_kerja from './store/unit_kerja'
import metric from './store/metric'
import barang from './store/barang'
import files from './store/files'
import inventory from './store/inventory'
import inventory_masuk from './store/inventory_masuk'
import inventory_keluar from './store/inventory_keluar'
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
        unit_kerja,
        files,
        surat_masuk,
        surat_keluar,
        juknis,
        inventory,
        inventory_masuk,
        inventory_keluar,
        asset_tetap,
        metric,
        barang,
        home,
        constant,
    }
});