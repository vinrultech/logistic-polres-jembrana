import defaultState from './inventory_masuk/state';
import getters from './inventory_masuk/getters';
import actions from './inventory_masuk/actions';
import mutations from './inventory_masuk/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}