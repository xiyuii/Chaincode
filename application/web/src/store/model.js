const state = {
  models: []
}

const mutations = {
  setModels(state, models) {
    state.models = models
  }
}

const actions = {
  fetchModels({ commit }) {
    // 调用API获取模型列表，并提交到mutations
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
