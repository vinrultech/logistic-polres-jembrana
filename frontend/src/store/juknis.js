import defaultState from './juknis/state';
import getters from './juknis/getters';
import actions from './juknis/actions';
import mutations from './juknis/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}