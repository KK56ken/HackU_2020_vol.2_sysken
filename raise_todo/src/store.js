import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    todos: [],
  },
  mutations: {
    add_todo(state, todo) {
      // 状態を変更する
      state.todos.push(todo);
    },
  },
  actions: {},
  modules: {},
});
