export default {
    items: (state) => {
        return state.items;
    },
    bidangs: (state) => {
        return state.bidangs;
    },
    item: (state) => (id) => {
        return state.items.find(item => item.id === id);
    },
    last_no: (state) => {
        return state.last_no;
    },
    first_no: (state) => {
        return state.first_no;
    },
    prev_show: (state) => {
        return state.prev_show;
    },
    next_show: (state) => {
        return state.next_show;
    },
    limit: (state) => {
        return state.limit;
    },
    bidang: (state) => {
        return state.bidang;
    },
}