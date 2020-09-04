import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    todos: [],
    user: { name: "", password: "", token: "" },
    now_feed: 20,
    ate_feed: 0,
    hp: 7,
    died: false,
    today: 0,
    month: 0,
    year: 0,
    nodl: 0,
   
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
          token:"aaa",
          name: todo.title,
          time_limite: todo.date,
          subject_id: todo.subject,
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
    post_add_feed(context,feed){
      axios
        .post("http://127.0.0.1:3001/raise/feed",{
          token:feed.token,
          flag:feed.flag,
          feednum:feed.feed_num,
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
    post_new_food(context,feed){
      axios
       .post("http://127.0.0.1:3001/raise/feed",{
         token:feed.token,
         flag:feed.flag,
         feednum:feed.feednum,

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
    }
  },
  modules: {},
});
