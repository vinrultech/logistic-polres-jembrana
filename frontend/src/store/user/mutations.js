export default {
    PAGINATOR: (state, payload) => {
        state.items = payload.items;
        state.last_no = payload.last_no
        state.first_no = payload.first_no
        state.prev_show = payload.prev_show
        state.next_show = payload.next_show
    },
    ALL: (state, payload) => {
        state.items = payload;
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