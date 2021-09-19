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
    },
    NAMA: (state, payload) => {
        state.nama = payload;
    },
    ACCOUNT: (state, payload) => {
        state.account = payload;
    },
    FOTO: (state, payload) => {
        state.foto = payload;
    },
    HP: (state, payload) => {
        state.hp = payload;
    }
}