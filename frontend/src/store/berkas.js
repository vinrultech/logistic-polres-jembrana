import defaultState from './berkas/state';
import getters from './berkas/getters';
import actions from './berkas/actions';
import mutations from './berkas/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}