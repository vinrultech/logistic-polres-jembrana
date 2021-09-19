import defaultState from './surat_masuk/state';
import getters from './surat_masuk/getters';
import actions from './surat_masuk/actions';
import mutations from './surat_masuk/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}