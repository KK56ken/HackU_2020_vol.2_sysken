<template>
  <div class="raising">
    <v-container>
      <v-row class="text-center">
        <v-col cols="6" sm="6" md="4" lg="3">
          
          <h1>育成</h1>
          <v-divider>  
          </v-divider>

           
        </v-col>
      </v-row>
            <v-system-bar color="#8CF37C" ></v-system-bar>
            
            <v-sheet color="#D4D4D4" height="300"></v-sheet>
            <p>えさの数０個</p>
            <v-avatar color="#7C5736" size="120">
             <span class="white--text headline">エサ </span> 
            </v-avatar> 

            <v-avatar color="#0094FF" size="78">
             <span class="white--text headline caption">コレクション </span> 
             
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
