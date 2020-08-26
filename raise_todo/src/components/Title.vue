<template>
  <div class="home">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <v-col :cols="this.size" sm="6" md="4" lg="3">
          <h1>{{titlename}}</h1>
          <v-divider></v-divider>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["titlename", "size"],
};
</script>
