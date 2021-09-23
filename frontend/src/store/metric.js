import defaultState from './metric/state';
import getters from './metric/getters';
import actions from './metric/actions';
import mutations from './metric/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}