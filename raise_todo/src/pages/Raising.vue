<template>
  <div class="raising">
    <v-container>
      <Title titlename="育成" size="6" />
      <v-row class="flex-child">
        <v-col class="d-flex" cols="12" md="4">
          <div v-for="i in hp" :key="i">
            <v-card
              outlined
              tile
              color="#8CF37C"
              width="50"
              height="25"
            ></v-card>
          </div>
        </v-col>
      </v-row>

      <v-sheet color="#D4D4D4" height="300" width="350">
        <template v-slot:activator="{ on }">
          <v-sheet color="#D4D4D4" height="300" v-on="on"></v-sheet>
        </template>
        <div>
          <img class="img" src="../assets/character/beechild.png" width="300" />
        </div>
      </v-sheet>
      <p>エサの数{{ this.$store.state.food }}個</p>
      <v-avatar v-on:click="feed_consume()" color="#7C5736" size="120">
        <span class="white--text headline">エサ</span>
      </v-avatar>
      <router-link to="/Collection">
        <v-avatar color="#0094FF" size="78">
          <span class="white--text headline caption">コレクション</span>
        </v-avatar>
      </router-link>
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
    {{ this.startTime }}
    {{ this.nowTime }}
  </div>
</template>

<script>
import Title from "@/components/Title.vue";

export default {
  data() {
    return {
      nowTime: 0,
      // 現在時刻
      diffTime: 0,
      // スタートボタンを押した時刻
      startTime: 0,
      hp: 7,
      feed_num: 10,
      count: 0,
    };
  },
  created() {
    var date = new Date();
    this.startTime =
      date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
    setInterval(() => {
      var date = new Date();
      this.nowTime =
        date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
    }, 1000);
  },
  methods: {
    feed_consume() {
      if (this.feed_num > 0) {
        this.$store.state.food--;
      }
    },
  },
  watch: {
    nowTime: function(newValue, oldValue) {
      console.log(newValue, oldValue);
      if (this.hp > 0) {
        this.hp = this.hp - 1;
      }
    },
  },
  components: {
    Title,
  },
};
</script>
