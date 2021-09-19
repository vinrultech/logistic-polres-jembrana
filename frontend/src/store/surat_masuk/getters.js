export default {
    items: (state) => {
        return state.items;
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
    }
}