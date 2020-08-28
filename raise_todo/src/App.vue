<template>
  <v-app>
    <v-sheet color="#FFF7E8">
    <v-main>
      <router-view></router-view>
    </v-main>
    </v-sheet>
  </v-app>
</template>

<script>
// import Home from "./pages/Home";

export default {
  name: "App",

  components: {
    // Home,
  },

  data: () => ({
    //
  }),
};
</script>
