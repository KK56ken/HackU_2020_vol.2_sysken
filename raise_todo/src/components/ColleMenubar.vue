<template>
<v-container>
    <v-bottom-navigation
    horizontal
    >
      <router-link to="/">
        <v-btn dark > 
          <span>{{btnname}}</span>
        </v-btn>
      </router-link>
       <router-link to="/Raising">
        <v-btn dark > 
          <span>{{Rname}}</span>
        </v-btn>
      </router-link>
    </v-bottom-navigation>
    </v-container>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname","Rname"],
};
</script>