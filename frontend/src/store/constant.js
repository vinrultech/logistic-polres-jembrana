const defaultState = {
    limits: [
        {
            value: 5,
            text: '5'
        },
        {
            value: 10,
            text: '10'
        },
        {
            value: 25,
            text: '25'
        },
        {
            value: 100,
            text: '100'
        }
    ],
    menu: 0,
    error: false,
    errorMessage: ''
}

const actions = {
    menu: (context, payload) => {
        if (payload === undefined) {
            const item = context.getters.menu
            context.commit('MENU', payload);
            setTimeout(function(){
                context.commit('MENU', item);
            }, 1);
            
        } else {
            context.commit('MENU', payload);
        }
        
    },
    error: (context, payload) => {
        context.commit('ERROR', payload)
    },
    set_error: (context, payload) => {
        context.commit('SET_ERROR', {
            error: payload
        })
    },
    remove_error: (context) => {
        context.commit('REMOVE_ERROR')
    }
}

const mutations = {
    MENU: (state, payload) => {
        state.menu = payload;
    },
    REMOVE_ERROR: (state) => {
        state.error = false
        state.errorMessage = ''
    },
    ERROR: (state, payload) => {
        state.error = true
        state.errorMessage = payload
    },
    SET_ERROR: (state, payload) => {
        state.error = payload.error
        state.errorMessage = ''
    }
}

const getters = {
    limits: (state) => {
        return state.limits;
    },
    menu: (state) => {
        return state.menu;
    },
    error: (state) => {
        return state.error;
    },
    errorMessage: (state) => {
        return state.errorMessage;
    },
}

export default {
    namespaced: true,
    state: defaultState,
    getters,
    actions,
    mutations,
}