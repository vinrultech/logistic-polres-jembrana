//import { union } from '../../utils';
import utils from '../../utils'

export default {
    PAGINATOR: (state, items) => {
        state.display_items = items;
        if (items.length > 0) {
            state.last_id = items[items.length - 1].id
            
        }
        state.items = utils.arrayUnion(state.items, items, 'id');
    },
    ALL: (state, all_items) => {
        state.all_items = all_items;
    },
    DISPLAY_ITEMS: (state, items) => {
        if (items.length > 0) {
            state.last_id = items[items.length - 1].id
        }
        state.display_items = items;
    },
    ITEM: (state, item) => {
        state.item = item;
    },
    LIMIT: (state, limit) => {
        state.limit = limit;
    },
    PREV: (state, prev) => {
        state.prev_show = prev;
    },
    SEARCH: (state, item) => {
        state.search = item.search;
        state.filter_search = item.filter_search;
    },
    SET_FILTERS: (state, filters) => {
        state.filters = filters
    },
    RESET: (state) => {
        state.last_id = 0;
        state.items = [];
        state.display_items = [];
        state.item = {};
        state.search = ""
        state.filter_search = ""
    },
}