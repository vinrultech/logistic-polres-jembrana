import defaultState from './petugas_ukur/state';
import getters from './petugas_ukur/getters';
import actions from './petugas_ukur/actions';
import mutations from './petugas_ukur/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}