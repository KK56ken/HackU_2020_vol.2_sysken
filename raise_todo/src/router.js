import Vue from "vue";
import VueRouter from "vue-router";
import Home from "./pages/Home.vue";
import Todo from "./pages/Todo.vue";
import Calendar from "./pages/Calendar.vue";
import Raiging from "./pages/Raising.vue";
import Collection from "./pages/Collection.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  // {
  //   path: "/about",
  //   name: "About",
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () =>
  //     import(/* webpackChunkName: "about" */ "./pages/About.vue"),
  // },
  {
    path: "/todo",
    name: "Todo",
    component: Todo,
  },
  {
    path: "/calendar",
    name: "Calendar",
    component: Calendar,
  },
  {
    path: "/raising",
    name: "Raiging",
    component: Raiging,
  },
  {
    path: "/collection",
    name: "Collection",
    component: Collection,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
