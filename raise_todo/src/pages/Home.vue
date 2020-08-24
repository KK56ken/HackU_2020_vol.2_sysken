<template>
  <div class="home">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <v-col cols="6" sm="6" md="4" lg="3">
          <h1>TitleMITEI</h1>
          <v-divider></v-divider>
        </v-col>
      </v-row>
      <HomeIcon />
    </v-container>
  </div>
</template>

<script>
// @ is an alias to /src
import HomeIcon from "@/components/HomeIcon.vue";

export default {
  name: "Home",
  components: {
    HomeIcon,
  },
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
