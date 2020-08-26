<template>
  <v-row>
    <v-col cols="12" sm="12" md="4" lg="3">
      <p>あと3日</p>
      <v-card class="mx-auto" height="80" width="350">
        <v-card-text>
          <span class="display-1 text--primary">{{todo.title}}</span>
          <v-btn class="text-right" text color="deep-purple accent-4">終了</v-btn>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["todo"],
};
</script>
