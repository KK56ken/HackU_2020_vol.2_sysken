<template>
  <div class="raising">
    <v-container>
      <Title titlename="育成" size="6"/>

      <v-system-bar color="#8CF37C"></v-system-bar>

      <v-sheet color="#D4D4D4" height="300">
        <template v-slot:activator="{ on }">
          <v-sheet color="#D4D4D4" height="300" v-on="on"></v-sheet>
        </template>
        <div>
          <img class="img" src="../assets/character/beeadult.png" width="300" />
        </div>
      </v-sheet>
      <p>エサの数０個</p>
      <v-avatar color="#7C5736" size="120">
        <span class="white--text headline">エサ</span>
      </v-avatar>

      <v-avatar color="#0094FF" size="78">
        <span class="white--text headline caption">コレクション</span>
      </v-avatar>
    </v-container>
    <v-container>
      <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Calendar">
        <v-btn color="green" dark>カレンダー</v-btn>
      </router-link>
      <router-link to="/Todo">
        <v-btn color="green" dark>ToDoリスト</v-btn>
      </router-link>
    </v-container>
  </div>
</template>




<script>
import Title from "@/components/Title.vue";

export default {
  data() {
    return {};
  },
  components: {
    Title,
  },
};
</script>
