
<template>
  <div class="collection">
    <v-container>
      <v-row class="text-center">
        <v-col cols="8" sm="6" md="4" lg="3">
          <h1>コレクション</h1>
          <v-divider> </v-divider>
        </v-col>
      </v-row>
 
      <v-sheet color="#D4D4D4" height="600" >
        <v-row>
          <v-col cols="3" class="ma-2"> 
            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/beechild.png"
              />
            </v-sheet>
            </v-col>
            <v-col cols="3" class="ma-2">

            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/beeadult.png"
              />
            </v-sheet>
            </v-col >
            <v-col lg="2" class="ma-2">
            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/beedie.png"
              />
            </v-sheet>
          </v-col>
        </v-row>
      </v-sheet>
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

