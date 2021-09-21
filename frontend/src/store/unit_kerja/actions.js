import injector from 'vue-inject'
import _ from 'lodash';


export default {
    fetch: injector.encase(['axios'], (axios) => (context, payload) => {
        const id = payload.id;
        axios.get(`/admin/unit_kerja/${id}`)
            .then((response) => {
                const item = response.data;
                context.commit('ITEM', item);
            })
    }),
    gets: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, payload) => {
        let q = `limit=${payload.limit}&last_id=${payload.last_id}`
        axios.get(`/admin/unit_kerja?${q}`)
            .then((response) => {
                const items = response.data;
                if (items.length > 0) {
                    context.commit('PAGINATOR', items);
                    if (payload.last_id == 0) {
                        context.commit('PREV', false);
                    } else {
                        context.commit('PREV', true);
                    }
                } else {
                    if (payload.last_id !== 0) {
                        toastr.s("No more data next")
                    }
                }
            })
    }),
    search: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, payload) => {
        let q = `limit=${payload.limit}&last_id=${payload.last_id}&search=${payload.search}&filter=${payload.filter}`;
        axios.get(`/admin/unit_kerja/search?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('SEARCH', {
                    search: payload.search,
                    filter_search: payload.filter,
                })
                if (items.length > 0) {
                    context.commit('PAGINATOR', items);
                    if (payload.last_id == 0) {
                        context.commit('PREV', false);
                    } else {
                        context.commit('PREV', true);
                    }
                } else {
                    if (payload.last_id !== 0) {
                        toastr.s("No more data next")
                    }
                }
            })
    }),
    create: injector.encase(['axios', 'router'], (axios, router) => (context, item) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        axios.post(`/admin/unit_kerja/create`, item, {
                message: "Unit Kerja berhasil di tambah"
            })
            .then(() => {
                router.push("/admin/unit_kerja")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
    remove: injector.encase(['axios'], (axios) => (_, id) => {
        return new Promise((resolve) => {
            axios.delete(`/admin/unit_kerja/remove/${id}`, {
                    message: "Unit Kerja berhasil di hapus"
                })
                .then(() => {
                    resolve(true);
                })
        });
    }),
    update: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        const id = payload.id;
        const item = payload.item;
        axios.put(`/admin/unit_kerja/update/${id}`, item, {
                message: "Unit Kerja berhasil di update"
            })
            .then(() => {
                router.push("/admin/unit_kerja")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
    next: (context, is_search) => {
        const limit = context.getters.limit.value;
        const last_id = context.getters.last_id;
        const index = _.findIndex(context.state.items, function (o) {
            return o.id === last_id
        });
        //console.log(`start : ${index + 1} , end : ${index + limit}`)
        

        const items = _.slice(context.state.items, index + 1, index + 1 + limit)
        if (items.length < limit) {
            if (is_search) {
                const search = context.getters.search;
                const filter_search = context.getters.filter_search;
                context.dispatch('search', {
                    last_id: last_id,
                    limit: limit,
                    filter: filter_search,
                    search: search,
                });
            } else {
                context.dispatch('gets', {
                    last_id: last_id,
                    limit: limit,
                });
            }
            
        } else {
            context.commit('PREV', true);
            context.commit('DISPLAY_ITEMS', items)
        }


    },
    prev: injector.encase(['toastr'], (toastr) => (context) => {
        const limit = context.getters.limit.value;
        const last_id = context.getters.last_id;
        const index = _.findIndex(context.state.items, function (o) {
            return o.id === last_id
        });
        
        
        let start = index - (limit - 1) - limit;
        let end = index - (limit - 1);

        if (context.state.display_items.length < limit) {
            start = index - limit - (context.state.display_items.length - 1);
            end = index - (context.state.display_items.length - 1);
        }

        

        const items = _.slice(context.state.items, start, end)
        if (items.length > 0) {
            context.commit('DISPLAY_ITEMS', items)
        } else {
            context.commit('PREV', false);
            toastr.s("No more data previous")
        }
    }),
    reset: (context) => {
        context.commit('RESET');
    },
    limit: (context, payload) => {
        context.commit('LIMIT', payload)
        /*
        const limit = context.getters.limit.value;
        const last_id = context.getters.last_id;
        context.dispatch('gets', {
            last_id: last_id,
            limit: limit,
        });
        */
    },
}