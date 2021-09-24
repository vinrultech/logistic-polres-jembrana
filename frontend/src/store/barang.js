import defaultState from './barang/state';
import getters from './barang/getters';
import actions from './barang/actions';
import mutations from './barang/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}