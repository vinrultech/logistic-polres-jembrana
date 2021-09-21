import defaultState from './files/state';
import getters from './files/getters';
import actions from './files/actions';
import mutations from './files/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}