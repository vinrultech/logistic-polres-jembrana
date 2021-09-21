
//import _ from 'lodash';

export default {
    ADD_FILE: (state, file) => {
        state.files.push(file)
    },
    REMOVE_FILE: (state, file_id) => {
       const files = state.files.filter(function(item){
           return item.file_id !== file_id
       })
       state.files = files;
    },
    SET: (state, files) => {
        state.files = files
    },
    RESET: (state) => {
        state.files = [];
    },
}