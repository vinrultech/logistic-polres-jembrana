import Vue from 'vue'
import Vuex from 'vuex'

import user from './store/user'
import berkas from './store/berkas'
import petugas_ukur from './store/petugas_ukur'
import petugas_gambar from './store/petugas_gambar'
import problem from './store/problem'
import home from './store/home'
import constant from './store/constant';

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        user,
        berkas,
        petugas_ukur,
        petugas_gambar,
        problem,
        home,
        constant,
    }
});