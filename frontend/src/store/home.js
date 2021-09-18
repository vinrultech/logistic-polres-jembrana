import defaultState from './home/state';
import getters from './home/getters';
import actions from './home/actions';
import mutations from './home/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}