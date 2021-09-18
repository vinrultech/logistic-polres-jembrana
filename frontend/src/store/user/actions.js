import utils from '../../utils'
import injector from 'vue-inject'

export default {
    login: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        axios.post("user/login", {
            username: payload.username,
            password: payload.password,
        }).then((response) => {
            const data = response.data
            utils.addSession(data.token, data.username, data.role, data.nama, data.foto);
            router.push("/admin")
        }).catch((error) => {
            const data = error.response.data;
            context.dispatch('constant/error', data.message, {
                root: true
            })
        })

    }),
    fetch: injector.encase(['axios'], (axios) => (context, payload) => {
        const username = payload.username;
        axios.get(`/admin/user/account/${username}`)
            .then((response) => {
                const data = response.data
                context.commit('NAMA', data.nama);
                context.commit('FOTO', data.foto);
                context.commit('HP', data.hp);
            })
    }),
    upload_image: injector.encase(['axios'], (axios) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        return new Promise((resolve) => {
            axios.post(`/admin/user/upload_image/${payload.username}`, payload.formData, {
                    message: "Upload image berhasil"
                })
                .then((response) => {
                    const data = response.data
                    utils.changeFoto(data.filename);
                    context.dispatch('set_foto', data.filename);
                    resolve(data)
                }).
            catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
        });
    }),
    remove_image: injector.encase(['axios'], (axios) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        return new Promise((resolve) => {
            axios.put(`/admin/user/remove_image`, payload.item, {
                message: "Remove image berhasil"
            }).then(() => {
                utils.changeFoto('');
                context.dispatch('set_foto', '');
                resolve(true)
            }).
            catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
        });
    }),
    change: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        });
        axios.put("/admin/user/change", payload.item, {
            message: "Password berhasil di rubah"
        }).then(() => {
            router.push("/admin")
        }).
        catch((error) => {
            const data = error.response.data;
            context.dispatch('constant/error', data.message, {
                root: true
            });
        })
    }),
    account: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        })
        axios.put("/admin/user/account", payload.item, {
            message: "Account berhasil di rubah"
        }).then((response) => {
            const data = response.data
            utils.changeNama(data.nama);
            context.commit('ACCOUNT', data.nama);
            router.push("/admin")
        }).
        catch((error) => {
            const data = error.response.data;
            context.dispatch('constant/error', data.message, {
                root: true
            });
        });
    }),
    nama: (context) => {
        context.commit('ACCOUNT', utils.nama());
    },
    foto: (context) => {
        context.commit('FOTO', utils.foto());
    },
    set_nama: (context, payload) => {
        context.commit('NAMA', payload);
    },
    set_foto: (context, payload) => {
        context.commit('FOTO', payload);
    },
    set_hp: (context, payload) => {
        context.commit('HP', payload);
    },
}