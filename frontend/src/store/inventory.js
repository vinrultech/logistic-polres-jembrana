import defaultState from './inventory/state';
import getters from './inventory/getters';
import actions from './inventory/actions';
import mutations from './inventory/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}