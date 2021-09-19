import injector from 'vue-inject'

export default {
    fetch: injector.encase(['axios'], (axios) => (context, payload) => {
        const id = payload.id;
        axios.get(`/admin/kategori/${id}`)
            .then((response) => {
                const item = response.data;
                context.commit('ITEM', item);
            })
    }),
    gets: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}`
        axios.get(`/admin/kategori?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    search: injector.encase(['axios'], (axios) => (context, payload) => {
        let q = `limit=${payload.limit}&last_no=${payload.last_no}&search=${payload.search}&filter=${payload.filter}`
        axios.get(`/admin/kategori/search?${q}`)
            .then((response) => {
                const items = response.data;
                context.commit('PAGINATOR', items);
            })
    }),
    add: injector.encase(['axios', 'router'], (axios, router) => (context, item) => {
        axios.post(`/admin/kategori/add`, item)
            .then(() => {
                router.push("/admin/kategori")
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
        axios.put(`/admin/kategori/update/${id}`, item)
            .then(() => {
                router.push("/admin/kategori")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
}