import Axios from 'axios'

const state = {
  api_key: undefined
}

const mutations = {
  set_tmdb_api_key(state, apiKey) {
    state.api_key = apiKey
  }
}

const actions = {
  async fetchTMDBApi({ commit, rootState }) {
    const { status, data } = await Axios.get(`${rootState.baseURLApi}/tmdb`)
    if (status === 200) {
      commit('set_tmdb_api_key', data.key)
    } else {
      alert('Unable to get TMDB Api Key from the server!')
    }
  }
}

const getters = {
  tmdb_api_key: (state) => {
    return state.api_key
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}