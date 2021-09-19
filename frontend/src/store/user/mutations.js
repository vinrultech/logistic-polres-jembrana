export default {
    PAGINATOR: (state, items) => {
        state.items.push(items);
    },
    ALL: (state, payload) => {
        state.items = payload;
    },
    ITEM: (state, item) => {
        state.item = item;
    },
    LIMIT: (state, payload) => {
        state.limit = payload;
    }
}