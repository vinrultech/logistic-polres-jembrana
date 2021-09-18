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
            } = await axios.get(`/admin/berkas?${q}`, {
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
    remove: (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;
        return new Promise((resolve, reject) => {

            progress.show({
                title: 'Remove data',
                text: 'Sedang remove data ke server',
            })
            const token = utils.session();
            axios.delete(`/admin/berkas/remove/${payload.id}`, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            }).then(() => {
                progress.hide()
                toastr.s("Data berkas berhasil di hapus")
                resolve('success')
            }).catch(ex => {
                progress.hide()
                utils.error(ex, swal, router)
                reject('error')
            })
        });
    },
    save_berkas: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;
        try {
            progress.show({
                title: 'Tambah data',
                text: 'Sedang create data ke server',
            })
            const token = utils.session();
            await axios.post("/admin/berkas/create", payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            });
            progress.hide()
            toastr.s("Data user berhasil di tambahkan")
            router.push('/admin/berkas')
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    edit_update_berkas: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;
        try {
            progress.show({
                title: 'Update data',
                text: 'Sedang update data ke server',
            })
            const token = utils.session();
            await axios.put(`/admin/berkas/update/${payload.id}`, payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            });
            progress.hide()
            toastr.s("Data user berhasil di update")
            router.push('/admin/berkas')
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }
    },
    edit_remove_bidang: async (context, payload) => {
        //console.log(payload)

        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;


        try {
            progress.show({
                title: 'Delete data',
                text: 'Sedang delete data ke server',
            })
            const token = utils.session();
            const {
                data
            } = await axios.delete(`/admin/bidang/remove/${payload.id}`, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            });
            progress.hide()
            toastr.s("Data user berhasil di update")
            context.commit('REMOVE_BIDANG_EDIT', data.id);
        } catch (ex) {
            progress.hide()
            utils.error(ex, swal, router);
        }



    },
    edit_update_bidang: (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;


        return new Promise((resolve) => {

            progress.show({
                title: 'Update data',
                text: 'Sedang update data ke server',
            })
            const token = utils.session();

            axios.put(`/admin/bidang/update/${payload.id}`, payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            }).then((data) => {

                context.commit('UPDATE_BIDANG_EDIT', data.data.bidang);
                progress.hide()
                toastr.s("Data bidang berhasil di update")
                resolve('success')
            }).catch(ex => {
                //console.log(ex);
                progress.hide()
                utils.error(ex, swal, router)

            });


        });
    },
    edit_add_bidang: (context, payload) => {

        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const progress = payload.vm.$progress;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;


        return new Promise((resolve) => {

            progress.show({
                title: 'Tambah data',
                text: 'Sedang tambah data ke server',
            })
            const token = utils.session();

            axios.post(`/admin/bidang/create`, payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            }).then((data) => {

                context.commit('ADD_BIDANG_EDIT', data.data.bidang);
                progress.hide()
                toastr.s("Data bidang berhasil di tambahkan")
                resolve('success')
            }).catch(ex => {
                progress.hide()
                utils.error(ex, swal, router)

            });


        });



    },
    edit_get_bidang: ( /* context, payload */ ) => {

    },
    add_berkas: (context, payload) => {
        return new Promise((resolve) => {

            context.commit('ITEM', payload.item);
            resolve('success')

        });

    },
    null_berkas: (context) => {
        return new Promise((resolve) => {

            context.commit('NULL_ITEM');
            resolve('success')

        });

    },
    null_items: (context) => {
        return new Promise((resolve) => {

            context.commit('NULL_ITEMS');
            resolve('success')

        });

    },
    add_bidang: (context, payload) => {
        return new Promise((resolve) => {

            context.commit('ADD_BIDANG', payload.item);
            resolve('success')

        });

    },
    add_item: (context, payload) => {
        return new Promise((resolve) => {

            context.commit('ADD_BIDANG', payload.item);
            resolve('success')

        });

    },
    add_bidangs: (context, payload) => {
        context.commit('ADD_BIDANGS', payload);

    },
    update_bidang: (context, payload) => {
        return new Promise((resolve) => {
            context.commit('UPDATE_BIDANG', payload);
            resolve('success')

        });
    },
    remove_bidang: (context, payload) => {
        return new Promise((resolve) => {

            context.commit('REMOVE_BIDANG', payload);
            resolve('success')

        });

    },
    send_wa: async (context, payload) => {
        const axios = payload.vm.$axios;
        const swal = payload.vm.$swal;
        const router = payload.vm.$router;
        const toastr = payload.vm.$toastr;
        try {
            const token = utils.session();
            await axios.post("/admin/wa/send", payload.item, {
                headers: {
                    'Content-Type': `application/json`,
                    'Authorization': `Bearer ${token}`,
                }
            });
            
            toastr.s("Message berhasil di kirim")
            
        } catch (ex) {
            
            utils.error(ex, swal, router);
        }
    },
    limit: (context, payload) => {
        context.commit('LIMIT', payload)
    },
}