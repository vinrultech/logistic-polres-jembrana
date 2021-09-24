import utils from '../../utils'
import injector from 'vue-inject'
import _ from 'lodash';

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
    reset_password: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        context.dispatch('constant/remove_error', {}, {
            root: true
        });
        axios.put("/admin/user/reset_password", payload.item, {
            message: "Password berhasil di rubah"
        }).then(() => {
            router.push("/admin/user")
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
    gets: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, payload) => {
        let q = `limit=${payload.limit}&last_id=${payload.last_id}`
        axios.get(`/admin/user?${q}`)
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
        axios.get(`/admin/user/search?${q}`)
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
        axios.post(`/admin/user/create`, item, {
                message: "User berhasil di tambah"
            })
            .then(() => {
                router.push("/admin/user")
            }).catch((error) => {
                const data = error.response.data;
                context.dispatch('constant/error', data.message, {
                    root: true
                })
            })
    }),
    remove: injector.encase(['axios'], (axios) => (_, id) => {
        return new Promise((resolve) => {
            axios.delete(`/admin/user/remove/${id}`, {
                    message: "User berhasil di hapus"
                })
                .then(() => {
                    resolve(true);
                })
        });
    }),
    update: injector.encase(['axios', 'router'], (axios, router) => (context, payload) => {
        const id = payload.id;
        const item = payload.item;
        axios.put(`/admin/user/update/${id}`, item, {
                message: "User berhasil di update"
            })
            .then(() => {
                router.push("/admin/user")
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