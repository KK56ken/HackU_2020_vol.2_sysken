<template>
<v-container>
    <v-bottom-navigation
    horizontal
    >
    
      <router-link to="/">
        <v-btn > 
          <v-icon>mdi-home</v-icon> 
          {{btnname}}
        </v-btn>
      </router-link>
    </v-bottom-navigation>
    </v-container>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname"],
};
</script>