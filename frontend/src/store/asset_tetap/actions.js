import injector from 'vue-inject'

export default {
    fetch: injector.encase(['axios'], (axios) => (context, payload) => {
        const id = payload.id;
        axios.get(`/admin/asset_tetap/${id}`)
            .then((response) => {
                const item = response.data;
                context.commit('ITEM', item);
            })
    }),
    gets: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}`
        axios.get(`/admin/asset_tetap?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    search: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}&search=${payload.search}&filter=${payload.filter}`
        axios.get(`/admin/asset_tetap/search?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    add: injector.encase(['axios', 'router'], (axios, router) => (context, item) => {
        axios.post(`/admin/asset_tetap/add`, item)
            .then(() => {
                router.push("/admin/asset_tetap")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
    update: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        const id = payload.id;
        const item = payload.item;
        axios.put(`/admin/asset_tetap/update/${id}`, item)
            .then(() => {
                router.push("/admin/asset_tetap")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
}