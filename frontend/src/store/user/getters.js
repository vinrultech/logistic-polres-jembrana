export default {
    items: (state) => {
        return state.display_items;
    },
    item: (state) => (id) => {
        return state.items.find(item => item.id === id);
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
    last_id: (state) => {
        return state.last_id;
    },
    search: (state) => {
        return state.search;
    },
    filter_search: (state) => {
        return state.filter_search;
    },
    nama: (state) => {
        return state.nama;
    },
    account: (state) => {
        return state.account;
    },
    foto: (state) => {
        return state.foto === null ? '' : state.foto;
    },
    hp: (state) => {
        return state.hp;
    }
}