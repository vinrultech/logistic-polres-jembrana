import utils from '../../utils'

export default {
    paginator: async (context, payload) => {
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
            const token = utils.session();
            const {
                data
            } = await axios.get(`/admin/berkas/ukur?${q}`, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            })
            context.commit('PAGINATOR', data);
            progress.hide()
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    proses: (context, payload) => {

        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;


        return new Promise((resolve) => {

            progress.show({
                title: 'Update status',
                text: 'Sedang update status ukur ke server',
            })
            const token = utils.session();

            axios.put(`/admin/bidang/ukur/${payload.id}`, payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            }).then((data) => {

                const bidang = data.data.bidang;

                const item = {
                    id: bidang.berkas_id,
                    bidang_id: bidang.id,
                    status_ukur: bidang.status_ukur,
                }
                context.commit('CHANGE_STATUS', item);
                progress.hide()
                toastr.s("Data bidang berhasil di tambahkan")
                resolve('success')
            }).catch(ex => {
                progress.hide()
                utils.error(ex, swal, router)
            });
        });
    },
    pilih: (context, payload) => {

        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;


        return new Promise((resolve) => {

            progress.show({
                title: 'Update status',
                text: 'Sedang update status berkas ke server',
            })
            const token = utils.session();

            axios.put(`/admin/berkas/gambar`, payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            }).then(() => {
                progress.hide()
                toastr.s("Data bidang berhasil di tambahkan")
                resolve('success')
            }).catch(ex => {
                progress.hide()
                utils.error(ex, swal, router)

            });


        });



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