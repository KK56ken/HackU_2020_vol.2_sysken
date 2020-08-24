<template>
  <div class="home">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <v-col cols="6" sm="6" md="4" lg="3">
          <h1>Calendar</h1>
          <v-divider></v-divider>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Todo">
        <v-btn color="green" dark>Todoリスト</v-btn>
      </router-link>
    </v-container>
  </div>
</template>


<script>
// @ is an alias to /src

export default {
  name: "",
  components: {},
};
</script>
<style>
.theme--light.v-divider {
  border-color: #6b421d !important;
  border-width: 5px !important;
}
h1 {
  color: #6b421d !important;
}
</style>
