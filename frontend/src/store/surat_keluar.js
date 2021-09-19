import defaultState from './surat_keluar/state';
import getters from './surat_keluar/getters';
import actions from './surat_keluar/actions';
import mutations from './surat_keluar/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}