import _ from "lodash";
export default {
    PAGINATOR: (state, payload) => {
        state.items = _.concat(state.items, payload.items)
        state.last_no = payload.last_no
        state.first_no = payload.first_no
        //state.prev_show = payload.prev_show
        //state.next_show = payload.next_show
    },
    LIMIT: (state, payload) => {
        state.limit = payload;
    },
    NULL_ITEMS: (state) => {
        state.items = [];
        state.prev_show = false
        state.next_show = false
    },
    CHANGE_DATE: (state, payload) => {
        
        const item = state.items.find(item => item.id === payload.id);
        const index = state.items.indexOf(item)

        const bidang = item.bidangs.find(item => item.id === payload.bidang_id)

        const idx = item.bidangs.indexOf(bidang)
        

        state.items[index].bidangs[idx].tanggal_ukur = payload.tanggal_ukur

    },
    CHANGE_STATUS: (state, payload) => {
        
        const item = state.items.find(item => item.id === payload.id);
        const index = state.items.indexOf(item)

        const bidang = item.bidangs.find(item => item.id === payload.bidang_id)

        const idx = item.bidangs.indexOf(bidang)
        

        state.items[index].bidangs[idx].status_ukur = payload.status_ukur

    }

}