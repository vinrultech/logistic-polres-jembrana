import { DASHBOARD } from "./breadcum";

export default {
    error(err, swal, router) {
        //const code = ex.toString().match(/\d+/g).map(Number);
        const code = parseInt(err.response && err.response.status)
        const message = err.response.data.message;
        swal({
            icon: "error",
            title: "Oops...",
            text: message !== undefined || message !== null ? message : err,
            willClose: () => {
                if (code === 400 || code === 401 || code === 403 ||  code === 405) {
                    sessionStorage.removeItem("pelayanan-desa-melaya:token")
                    sessionStorage.removeItem("pelayanan-desa-melaya:username")
                    sessionStorage.removeItem("pelayanan-desa-melaya:role");
                    sessionStorage.removeItem("pelayanan-desa-melaya:nama");
                    sessionStorage.removeItem("pelayanan-desa-melaya:foto");
                    router.push('/login');
                }

            }
        });
    },
    session() {
        return sessionStorage.getItem("pelayanan-desa-melaya:token");
    },
    role() {
        return sessionStorage.getItem("pelayanan-desa-melaya:role");
    },
    username() {
        return sessionStorage.getItem("pelayanan-desa-melaya:username");
    },
    nama() {
        return sessionStorage.getItem("pelayanan-desa-melaya:nama");
    },
    foto() {
        return sessionStorage.getItem("pelayanan-desa-melaya:foto");
    },
    removeSession() {
        sessionStorage.removeItem("pelayanan-desa-melaya:token");
        sessionStorage.removeItem("pelayanan-desa-melaya:username");
        sessionStorage.removeItem("pelayanan-desa-melaya:role");
        sessionStorage.removeItem("pelayanan-desa-melaya:nama");
        sessionStorage.removeItem("pelayanan-desa-melaya:foto");
    },
    addSession(token, username, role, nama, foto) {
        sessionStorage.setItem("pelayanan-desa-melaya:token", token);
        sessionStorage.setItem("pelayanan-desa-melaya:username", username);
        sessionStorage.setItem("pelayanan-desa-melaya:role", role);
        sessionStorage.setItem("pelayanan-desa-melaya:nama", nama);
        sessionStorage.setItem("pelayanan-desa-melaya:foto", foto);
    },
    changeNama(nama) {
        sessionStorage.setItem("pelayanan-desa-melaya:nama", nama);
    },
    changeFoto(foto) {
        sessionStorage.setItem("pelayanan-desa-melaya:foto", foto);
    },
    breadcumOne(one) {
        const indexs = [];
        indexs.push(DASHBOARD);
        indexs.push(one)
        return indexs;
    },
    breadcumTwo(index, two) {
        const indexs = [];
        indexs.push(DASHBOARD);
        indexs.push(index)
        indexs.push(two)
        return indexs;
    }
}