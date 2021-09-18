import defaultState from './problem/state';
import getters from './problem/getters';
import actions from './problem/actions';
import mutations from './problem/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}