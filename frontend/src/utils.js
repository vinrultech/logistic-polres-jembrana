import {
    DASHBOARD
} from "./breadcum";
import _ from 'lodash';

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
                if (code === 400 || code === 401 || code === 403 || code === 405) {
                    sessionStorage.removeItem("logistic-polres-jembrana:token")
                    sessionStorage.removeItem("logistic-polres-jembrana:username")
                    sessionStorage.removeItem("logistic-polres-jembrana:role");
                    sessionStorage.removeItem("logistic-polres-jembrana:nama");
                    sessionStorage.removeItem("logistic-polres-jembrana:foto");
                    router.push('/login');
                }

            }
        });
    },
    session() {
        return sessionStorage.getItem("logistic-polres-jembrana:token");
    },
    role() {
        return sessionStorage.getItem("logistic-polres-jembrana:role");
    },
    username() {
        return sessionStorage.getItem("logistic-polres-jembrana:username");
    },
    nama() {
        return sessionStorage.getItem("logistic-polres-jembrana:nama");
    },
    foto() {
        return sessionStorage.getItem("logistic-polres-jembrana:foto");
    },
    removeSession() {
        sessionStorage.removeItem("logistic-polres-jembrana:token");
        sessionStorage.removeItem("logistic-polres-jembrana:username");
        sessionStorage.removeItem("logistic-polres-jembrana:role");
        sessionStorage.removeItem("logistic-polres-jembrana:nama");
        sessionStorage.removeItem("logistic-polres-jembrana:foto");
    },
    addSession(token, username, role, nama, foto) {
        sessionStorage.setItem("logistic-polres-jembrana:token", token);
        sessionStorage.setItem("logistic-polres-jembrana:username", username);
        sessionStorage.setItem("logistic-polres-jembrana:role", role);
        sessionStorage.setItem("logistic-polres-jembrana:nama", nama);
        sessionStorage.setItem("logistic-polres-jembrana:foto", foto);
    },
    changeNama(nama) {
        sessionStorage.setItem("logistic-polres-jembrana:nama", nama);
    },
    changeFoto(foto) {
        sessionStorage.setItem("logistic-polres-jembrana:foto", foto);
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
    },
    arrayUnion (arr1, arr2, identifier) {
        const array = [...arr1, ...arr2]

        return _.uniqBy(array, identifier)
    }

}