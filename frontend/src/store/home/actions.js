import utils from '../../utils'
export default {
    chart_petugas_ukur: async (context, payload) => {
        const axios = payload.vm.$axios;
        try {
            const {
                data
            } = await axios.get(`/chart/ukur`);
            context.commit('CHART_UKUR', data);
        } catch (ex) {
            console.log(ex)
        }
    },
    chart_petugas_gambar: async (context, payload) => {
        const axios = payload.vm.$axios;
        try {
            const {
                data
            } = await axios.get(`/chart/gambar`);
            context.commit('CHART_GAMBAR', data);
        } catch (ex) {
            console.log(ex)
        }
    },
    paginator_all: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        let q = `limit=${payload.limit}&action=${payload.action}&no=${payload.no}`
        if (payload.search) {
            q += `&search=${payload.search}&filter=${payload.filter}`
        }
        try {
            progress.show({
                title: 'Ambil data',
                text: 'Sedang ambil data dari server',
            })
            
            const {
                data
            } = await axios.get(`/berkas/all?${q}`)
            context.commit('PAGINATOR_ALL', data);
            progress.hide()
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    paginator_by: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        let q = `limit=${payload.limit}&action=${payload.action}&no=${payload.no}`
        if (payload.search) {
            q += `&search=${payload.search}&filter=${payload.filter}`
        }
        try {
            progress.show({
                title: 'Ambil data',
                text: 'Sedang ambil data dari server',
            })
            
            const {
                data
            } = await axios.get(`/berkas/by/${payload.status}?${q}`)
            context.commit('PAGINATOR_BY', data);
            progress.hide()
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    paginator_petugas: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        let q = `limit=${payload.limit}&action=${payload.action}&no=${payload.no}`
        if (payload.search) {
            q += `&search=${payload.search}&filter=${payload.filter}`
        }
        try {
            progress.show({
                title: 'Ambil data',
                text: 'Sedang ambil data dari server',
            })
            
            const {
                data
            } = await axios.get(`/berkas/petugas/${payload.role}/${payload.petugas_id}?${q}`)
            context.commit('PAGINATOR_PETUGAS', data);
            progress.hide()
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    null_items: (context) => {
        return new Promise((resolve) => {

            context.commit('NULL_ITEMS');
            resolve('success')

        });

    },
    limit: (context, payload) => {
        context.commit('LIMIT', payload)
    },
}