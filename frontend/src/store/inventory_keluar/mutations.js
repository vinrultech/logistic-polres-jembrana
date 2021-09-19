export default {
    PAGINATOR: (state, items) => {
        state.items.push(items);
    },
    ALL: (state, items) => {
        state.items = items;
    },
    ITEM: (state, item) => {
        state.item = item;
    },
    LIMIT: (state, limit) => {
        state.limit = limit;
    }
}