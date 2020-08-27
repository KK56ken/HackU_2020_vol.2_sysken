import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

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
  actions: {
    post_new_todo(context, todo) {
      axios
        .post("/todo/register", {
          name: todo.title,
          date: todo.date,
          subject: todo.subject,
        })
        .then(() => {
          console.log("送れた");
        })
        .catch((err) => {
          if (err.response) {
            console.log("送れなかった");
          }
        });
    },
  },
  modules: {},
});
