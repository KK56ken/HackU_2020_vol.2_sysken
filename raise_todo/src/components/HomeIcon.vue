<template>
  <v-container>
    <v-row justify="space-around">
      <v-col cols="12" sm="6" md="4" lg="3">
        <router-link to="/Todo">
          <v-avatar color="#F29993" size="150">
            <span class="white--text headline">Todoリスト</span>
          </v-avatar>
        </router-link>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
};
</script>
<style>
</style>
