import defaultState from './petugas_gambar/state';
import getters from './petugas_gambar/getters';
import actions from './petugas_gambar/actions';
import mutations from './petugas_gambar/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}