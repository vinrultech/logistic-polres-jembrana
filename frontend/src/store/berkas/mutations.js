import _ from "lodash";
export default {
    PAGINATOR: (state, payload) => {
        //console.log(payload.items)
        state.items = _.concat(state.items, payload.items)
        //state.items = payload.items;
        state.last_no = payload.last_no
        state.first_no = payload.first_no
        //state.prev_show = payload.prev_show
        //state.next_show = payload.next_show
    },
    ITEM: (state, payload) => {
        state.bidang.no_berkas = payload.no_berkas;
        state.bidang.pemohon = payload.pemohon;
        state.bidang.wa = payload.wa;
        state.bidang.tanggal_input = payload.tanggal_input;
    },
    NULL_ITEM: (state) => {
        state.bidang = {
            no_berkas: null,
            pemohon: null,
            wa: null,
            tanggal_input: null,
            bidangs: []
        };
    },
    
    NULL_ITEMS: (state) => {
        state.items = [];
        state.prev_show = false
        state.next_show = false
    },
    ADD_BIDANG: (state, payload) => {
        //console.log(payload)
        state.bidang.bidangs.push(payload);
        //state.bidang.bidangs = _.concat(state.bidang.bidangs, payload)
    },
    ADD_BIDANGS: (state, payload) => {
        state.bidangs = payload;
    },
    ADD_BIDANG_EDIT: (state, payload) => {
        state.bidangs.push(payload);
    },
    UPDATE_BIDANG_EDIT: (state, payload) => {
        const bidang = state.bidangs.find(item => item.id === payload.id)

        const index = state.bidangs.indexOf(bidang)

        state.bidangs[index].no_bidang = payload.no_bidang
        state.bidangs[index].jumlah = payload.jumlah
        state.bidangs[index].luas = payload.luas
    },
    REMOVE_BIDANG_EDIT: (state, payload) => {
        const bidang = state.bidangs.find(item => item.id === payload)

        const index = state.bidangs.indexOf(bidang)

        state.bidangs.splice(index, 1);
    },
    UPDATE_BIDANG: (state, payload) => {
        const bidang = state.bidang.bidangs.find(item => item.id_bidang === payload.id)

        const index = state.bidang.bidangs.indexOf(bidang)

        state.bidang.bidangs[index].no_bidang = payload.item.no_bidang
        state.bidang.bidangs[index].jumlah = payload.item.jumlah
        state.bidang.bidangs[index].luas = payload.item.luas
    },
    REMOVE_BIDANG: (state, payload) => {
        const bidang = state.bidang.bidangs.find(item => item.id_bidang === payload)

        const index = state.bidang.bidangs.indexOf(bidang)

        state.bidang.bidangs.splice(index, 1);
    },
    ALL: (state, payload) => {
        state.items = payload;
    },
    LIMIT: (state, payload) => {
        state.limit = payload;
    }
}