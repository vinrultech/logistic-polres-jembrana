import defaultState from './asset_tetap/state';
import getters from './asset_tetap/getters';
import actions from './asset_tetap/actions';
import mutations from './asset_tetap/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}