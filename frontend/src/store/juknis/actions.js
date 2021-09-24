import injector from 'vue-inject'
import _ from 'lodash';


export default {
    fetch: injector.encase(['axios'], (axios) => (context, row_id) => {
        return new Promise((resolve) => {
            axios.get(`/admin/juknis/${row_id}`)
                .then((response) => {
                    const item = response.data;
                    context.commit('ITEM', item);
                    context.commit("files/SET", item.files, {
                        root: true
                    });
                    resolve(true)
                })
        })

    }),
    gets: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, payload) => {
        let q = `limit=${payload.limit}&last_id=${payload.last_id}`

        if (context.getters.dates.length > 1) {
            const dates = context.getters.dates;
            q += `&start=${dates[0]}&end=${dates[1]}`
        }

        axios.get(`/admin/juknis?${q}`)
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
                        toastr.i("No more data next")
                    }
                }
            })
    }),
    search: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, payload) => {
        let q = `limit=${payload.limit}&last_id=${payload.last_id}&search=${payload.search}&filter=${payload.filter}`;
        if (context.getters.dates.length > 1) {
            const dates = context.getters.dates;
            q += `&start=${dates[0]}&end=${dates[1]}`
        }
        axios.get(`/admin/juknis/search?${q}`)
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
                        toastr.i("No more data next")
                    }
                }
            })
    }),
    create: injector.encase(['axios', 'router'], (axios, router) => (context, item) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        axios.post(`/admin/juknis/create`, item, {
                message: "Juknis berhasil di tambah"
            })
            .then(() => {
                router.push("/admin/juknis")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
    remove: injector.encase(['axios'], (axios) => (_, row_id) => {
        return new Promise((resolve) => {
            axios.delete(`/admin/juknis/remove/${row_id}`, {
                    message: "Juknis berhasil di hapus"
                })
                .then(() => {
                    resolve(true);
                })
        });
    }),
    update: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        const row_id = payload.row_id;
        const item = payload.item;
        axios.put(`/admin/juknis/update/${row_id}`, item, {
                message: "Juknis berhasil di update"
            })
            .then(() => {
                router.push("/admin/juknis")
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
            toastr.i("No more data previous")
        }
    }),
    reset: (context) => {
        context.commit('RESET');
    },
    limit: (context, payload) => {
        context.commit('LIMIT', payload)
    },
}