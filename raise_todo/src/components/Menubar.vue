<template>
<router-link to="/">
 <v-footer
    absolute
    height="50"
    color="brown lighten-1"
    dark
    >
    <v-icon large>mdi-home</v-icon>
    <h2 font-weight="50">{{ btnname }}</h2>
 </v-footer>
 </router-link>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname"],
};
</script>
