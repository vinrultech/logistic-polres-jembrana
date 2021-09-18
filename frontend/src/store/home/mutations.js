import _ from "lodash";
export default {
    PAGINATOR_ALL: (state, payload) => {
        state.items = _.concat(state.items, payload.items)
        state.last_no = payload.last_no
        state.first_no = payload.first_no
    },
    PAGINATOR_BY: (state, payload) => {
        state.items = _.concat(state.items, payload.items)
        state.last_no = payload.last_no
        state.first_no = payload.first_no
    },
    PAGINATOR_PETUGAS: (state, payload) => {
        state.items = _.concat(state.items, payload.items)
        state.last_no = payload.last_no
        state.first_no = payload.first_no
    },
    CHART_UKUR: (state, payload) => {
        state.ukurs = payload;
    },
    CHART_GAMBAR: (state, payload) => {
        state.gambars = payload;
    },
    NULL_ITEMS: (state) => {
        state.items = [];
        state.prev_show = false
        state.next_show = false
    },
    LIMIT: (state, payload) => {
        state.limit = payload;
    }
}