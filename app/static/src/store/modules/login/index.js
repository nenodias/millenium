import * as types from '../../mutation-types'
import * as api from '../../../api/login'

// state
const state = {
  token: null,
  processing: false
}

// getters
const getters = {
  token: state => state.token,
  processing: state => state.processing
}

// actions
const actions = {
  async doLogin ({ commit, state }, dados) {
    const token = 'lsaldosadoasddsa'
    commit(types.LOGIN_REQUEST, { dados })
    console.log(api)
    let res = await api.login(dados)
    console.log(res)
    if (res && res.token) {
      const token = res.token
      commit(types.LOGIN_SUCCESS, { token })
    } else {
      commit(types.LOGIN_FAILURE)
    }
  }
}

// mutations
const mutations = {
  [types.LOGIN_REQUEST] (state, { dados }) {
    state.processing = true
  },
  [types.LOGIN_FAILURE] (state, { dados }) {
    state.processing = false
  },
  [types.LOGIN_SUCCESS] (state, { token }) {
    state.processing = false
    state.token = token
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
