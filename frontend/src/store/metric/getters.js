export default {
    items: (state) => {
        return state.display_items;
    },
    all_items: (state) => {
        return state.all_items;
    },
    item: (state) => (id) => {
        return state.items.find(item => item.id === id);
    },
    find_all_item: (state) => (id) => {
        return state.all_items.find(item => item.id === id);
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
    }
}