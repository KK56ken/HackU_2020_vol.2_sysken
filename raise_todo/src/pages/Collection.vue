
<template>
  <div class="collection">
    <v-container>
      <v-row class="text-center">
        <v-col cols="8" sm="6" md="4" lg="3">
          <h1>コレクション</h1>
          <v-divider> </v-divider>
        </v-col>
      </v-row>
 
      <v-sheet color="#D4D4D4" height="500" >
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

             <v-col cols="3" class="ma-2"> 
            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/mausuchild.png"
              />
            </v-sheet>
            </v-col>

            <v-col cols="3" class="ma-2"> 
            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/mausuadult.png"
              />
            </v-sheet>
            </v-col>

             <v-col cols="3" class="ma-2"> 
            <v-sheet  width="100" height="100">
              <img
                width="100"
                height="100"
                src="../assets/character/mausudie.png"
              />
            </v-sheet>
          </v-col> 

        </v-row>
      </v-sheet>
    </v-container>
    <v-container>
     <ColleMenubar btnname="ホーム" Rname="育成"></ColleMenubar>
    </v-container>
  </div>
</template>

<script>
import ColleMenubar from "@/components/ColleMenubar.vue";
export default {
  components:{
  ColleMenubar
  },
}
</script>