import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    todos: [],
    user: { name: "", password: "", token: "" },
    now_feed: 1,
    ate_feed: 0,
    hp: 7,
    died: false,
  },
  mutations: {
    add_todo(state, todo) {
      // 状態を変更する
      state.todos.push(todo);
    },
    set_user(state, post_data) {
      state.user.name = post_data.name;
      state.user.password = post_data.password;
    },
  },
  actions: {
    post_new_todo(context, todo) {
      axios
        .post("http://127.0.0.1:3001/todo/register", {
          name: todo.title,
          date: todo.date,
          subject: todo.subject,
        })
        .then((res) => {
          console.log("送れた");
          if (res.data) {
            console.log(res.date);
          }
        })
        .catch((err) => {
          if (err.response) {
            console.log("送れなかった");
          }
        });
    },
    post_new_user(context, user) {
      axios
        .post("http://127.0.0.1:3001/user/signup", {
          user_name: user.name,
          password: user.password,
        })
        .then((res) => {
          console.log(res.data);
          if (res.data) {
            context.commit("set_user", res.data);
          }
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
