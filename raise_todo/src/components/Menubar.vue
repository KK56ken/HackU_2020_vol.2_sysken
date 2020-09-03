<template>
  <v-container>
    <router-link to="/">
      <v-btn width="360" height="50">
        <v-icon>mdi-home</v-icon>
        <h2 font-weight="50">{{ btnname }}</h2>
      </v-btn>
    </router-link>
  </v-container>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname"],
};
</script>
