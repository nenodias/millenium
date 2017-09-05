import * as types from '../../mutation-types'
import * as api from '../../../api/login'
import * as acts from './actions'

// state
const state = {
  token: null,
  login_processing: false,
  login_message: null,
  login_done: false
}

// getters
const getters = {
  token: state => state.token,
  login_processing: state => state.login_processing,
  login_message: state => state.login_message,
  login_done: state => state.login_done
}

// actions
const actions = {
  async [acts.LOGIN_DO_LOGIN] ({ commit, state }, dados) {
    const token = 'lsaldosadoasddsa'
    commit(types.LOGIN_REQUEST, { dados })
    let res = await api.login(dados)
    if (res && res.token) {
      const token = res.token
      commit(types.LOGIN_SUCCESS, { token })
      setTimeout( () => commit(types.LOGIN_DONE), 2000 )
    } else {
      commit(types.LOGIN_FAILURE, dados)
    }
  },
  async [acts.LOGIN_CLEAR_MESSAGE] ({ commit, state }) {
    commit(types.LOGIN_CLEAR_MESSAGE)
  },
  async [acts.LOGIN_VERIFY_COOKIE] ({ commit, state }, token) {
    commit(types.LOGIN_BY_COOKIE, token)
  }
}

// mutations
const mutations = {
  [types.LOGIN_REQUEST] (state, { dados }) {
    state.login_processing = true
    state.login_message = null
  },
  [types.LOGIN_FAILURE] (state, dados) {
    state.login_message = {
      type:'is-danger',
      text:'Falha ao efetuar login'
    }
    state.login_processing = false
  },
  [types.LOGIN_SUCCESS] (state, { token }) {
    state.login_processing = false
    state.token = token
    state.login_message = {
      type:'is-success',
      text:'Login confirmado com sucesso!'
    }
  },
  [types.LOGIN_DONE] (state) {
    state.login_done = true
  },
  [types.LOGIN_CLEAR_MESSAGE] (state) {
    state.login_message = null
  },
  [types.LOGIN_BY_COOKIE] (state, token) {
    state.login_message = null
    state.login_processing = false
    state.login_done = true
    state.token = token
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
