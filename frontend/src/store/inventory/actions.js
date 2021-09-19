import injector from 'vue-inject'

export default {
    fetch: injector.encase(['axios'], (axios) => (context, payload) => {
        const id = payload.id;
        axios.get(`/admin/inventory/${id}`)
            .then((response) => {
                const item = response.data;
                context.commit('ITEM', item);
            })
    }),
    gets: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}`
        axios.get(`/admin/inventory?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    search: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}&search=${payload.search}&filter=${payload.filter}`
        axios.get(`/admin/inventory/search?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    add: injector.encase(['axios', 'router'], (axios, router) => (context, item) => {
        axios.post(`/admin/inventory/add`, item)
            .then(() => {
                router.push("/admin/inventory")
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
        axios.put(`/admin/inventory/update/${id}`, item)
            .then(() => {
                router.push("/admin/inventory")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
}