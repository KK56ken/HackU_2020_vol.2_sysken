<template>
  <div class="raising">
    <v-container>
      <Title titlename="育成" size="6" />
      <v-row class="flex-child">
        <v-col class="d-flex" cols="12" md="4">
          <div v-for="i in this.$store.state.hp" :key="i">
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
  
      <v-alert 
        width="350"
        color="red darken-1"
        dark
        transition="scale-transition"
        :value="alert"
      >
        うぅ...おなかいっぱいだよ〜
      </v-alert>
     
      <v-sheet color="#D4D4D4" height="300" width="350">
        <template v-slot:activator="{ on }">
          <v-sheet color="#D4D4D4" height="300" v-on="on"></v-sheet>
        </template>
        <div>
          <template v-if="!this.$store.state.died">
            <img
              class="img"
              src="../assets/character/beechild.png"
              width="300"
            />
          </template>
          <template v-else>
            <img class="img" src="../assets/character/beedie.png" width="300"
          /></template>
        </div>
      </v-sheet>
      <p>エサの数{{ this.$store.state.now_feed }}個</p>
      <v-row>
        <v-col offset="3">
      <v-avatar class="feed" v-on:click="feed_consume()" color="#7C5736" size="120">
        <span class="white--text headline">エサ</span>
      </v-avatar>
        </v-col>
        <v-col>
      <router-link to="/Collection">
        <v-avatar color="#0094FF" size="78">
          <span class="white--text headline caption">コレクション</span>
        </v-avatar>
      </router-link>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <!-- <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Calendar">
        <v-btn color="green" dark>カレンダー</v-btn>
      </router-link>
      <router-link to="/Todo">
        <v-btn color="green" dark>ToDoリスト</v-btn> 
        
      </router-link> -->
      <Menubar btnname="ホーム"></Menubar>
    </v-container>
    {{ this.startTime }}
    {{ this.nowTime }}
    {{ this.$store.state.hp }}
  </div>
</template>

<script>
import Title from "@/components/Title.vue";
import Menubar from "@/components/Menubar.vue";

export default {

  data() {
    
    return {
      // 現在時刻
      nowTime: 0,
      // スタートした時間
      startTime: 0,
      hp: 7,
      feed_num: 10,
      count: 0,
      alert:false,
    };
  },
  created() {
    var date = new Date();
    this.startTime =
      date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
    setInterval(() => {
      this.alert=false;
      var date = new Date();
      this.nowTime =
        date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
    }, 10000);
  },
  methods: {
    feed_consume() {
      if (this.$store.state.now_feed > 0 && this.$store.state.hp < 7 ){
        this.$store.state.now_feed--;
        this.$store.state.hp++;
      }else if(this.$store.state.hp===7){
        this.alert=true;
      }
    },
  },
  watch: {
    nowTime: function(newValue, oldValue) {
      console.log(newValue, oldValue);
      if (this.$store.state.hp > 0) {
        this.$store.state.hp--;
      }
    },
    // this.$store.state.now_feed:function(){

    // }
  },
  mounted() {
    this.$store.watch(
      () => {
        return this.$store.state.hp;
      },
      (newValue, oldValue) => {
        console.log(oldValue, newValue);
        if (newValue == 0) {
          this.$store.state.died = true;
        }
      },
      {
        deep: true,
      }
    );
  },
  components: {
    Title,
    Menubar
  },
};
</script>
