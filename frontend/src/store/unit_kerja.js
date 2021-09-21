import defaultState from './unit_kerja/state';
import getters from './unit_kerja/getters';
import actions from './unit_kerja/actions';
import mutations from './unit_kerja/mutations';

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}