import defaultState from './aset_tetap/state';
import getters from './aset_tetap/getters';
import actions from './aset_tetap/actions';
import mutations from './aset_tetap/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}