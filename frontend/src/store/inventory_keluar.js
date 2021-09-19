import defaultState from './inventory_keluar/state';
import getters from './inventory_keluar/getters';
import actions from './inventory_keluar/actions';
import mutations from './inventory_keluar/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}