<template>
<v-container>
      <router-link to="/">
        <v-btn  
        width="180"
        height="50"> 
            <v-icon>mdi-home</v-icon>
           <h3 font-weight="30">{{btnname}}</h3>
        </v-btn>
        </router-link>
        <router-link to="/Raising">
        <v-btn  
        width="180"
        height="50"
        color="#f8b739" 
        dark
        >
         <v-icon>mdi-tree</v-icon>
        
            <h3 font-weight="30">{{Rname}}</h3>
        </v-btn>
    </router-link>
    </v-container>
</template>

<script>
export default {
  name: "",
  data: () => ({}),
  props: ["btnname","Rname"],
};
</script>