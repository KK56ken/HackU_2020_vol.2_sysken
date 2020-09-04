<template>

  <v-footer
          absolute 
          height="100"
          width="100%"
          color="brown lighten-1"
          dark
          >
      <v-row>
        <v-col cols="6">
            <router-link to="/">
                <v-icon large>mdi-home</v-icon>
                <h3 font-weight="30">{{btnname}}</h3>
            <!-- </v-footer> -->
            </router-link>
          </v-col>
        </v-row>
        <v-row>
         <v-col cols="6"> 
           <router-link to="/Raising">
              <!--  <v-footer
                absolute 
                height="50"
                width="50%"
                color="#f8b739" 
                dark
                > -->
              <v-icon large>mdi-tree</v-icon>
              <h3 font-weight="30">{{Rname}}</h3>
          </router-link>
     </v-col>
  </v-row>
  </v-footer>
 
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname","Rname"],
};
</script>