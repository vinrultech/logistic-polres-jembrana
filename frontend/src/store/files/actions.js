import injector from 'vue-inject'

export default {
    upload: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, item) => {
        
        axios.post(`/admin/files/upload`, item, {
                message: "Files berhasil di upload"
            })
            .then((response) => {
                const file = response.data;
                context.commit('ADD_FILE', file)
            }).catch((error) => {
                const data = error.response.data;
                toastr.s(data.message)
            })
    }),
    remove: injector.encase(['axios', 'toastr'], (axios, toastr) => (context, item) => {
        
        axios.post(`/admin/files/remove`, item, {
                message: "Files berhasil di remove"
            })
            .then((response) => {
                const file = response.data;
                //console.log(file)
                //context.commit('ADD_FILE', file)
                context.commit("REMOVE_FILE", file.file_id);
            }).catch((error) => {
                const data = error.response.data;
                toastr.s(data.message)
            })
    }),
    reset: (context) => {
        console.log("RESET FILES")
        context.commit('RESET');
    },
    set: (context, files) => {
        context.commit('SET', files);
    }

}