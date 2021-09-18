export default {
    items: (state) => {
        return state.items;
    },
    ukurs: (state) => {
        return state.ukurs;
    },
    gambars: (state) => {
        return state.gambars;
    },
    item: (state) => (id) => {
        return state.items.find(item => item.id === id);
    },
    petugas_ukur: (state) => (id) => {
        return state.ukurs.find(item => item.id === id);
    },
    petugas_gambar: (state) => (id) => {
        return state.gambars.find(item => item.id === id);
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