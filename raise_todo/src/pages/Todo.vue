<template>
  <div class="Todo">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <v-col cols="7" sm="6" md="4" lg="3">
          <h1>やるべきこと</h1>
          <v-divider></v-divider>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <v-sheet class="d-flex" color="grey lighten-3" height="424"></v-sheet>
    </v-container>
    <v-container>
      <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Calendar">
        <v-btn color="green" dark>カレンダー</v-btn>
      </router-link>
    </v-container>
  </div>
</template>


<script>
// @ is an alias to /src
// import HomeIcon from "@/components/HomeIcon.vue";

export default {
  name: "",
  components: {
    // HomeIcon,
  },
};
</script>
<style>
h1 {
  color: #6b421d !important;
}
</style>
