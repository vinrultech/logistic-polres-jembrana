import defaultState from './kategori/state';
import getters from './kategori/getters';
import actions from './kategori/actions';
import mutations from './kategori/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}