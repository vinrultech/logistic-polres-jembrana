export default {
    items: (state) => {
        return state.display_items;
    },
    item: (state) => (row_id) => {
        return state.items.find(item => item.row_id === row_id);
    },
    get_item: (state) => {
        return state.item;
    },
    
    prev_show: (state) => {
        return state.prev_show;
    },
    dates: (state) => {
        return state.dates;
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