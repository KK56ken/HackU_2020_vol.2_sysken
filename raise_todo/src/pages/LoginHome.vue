<template>
  <div class="LoginHome">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
     <v-card width="400px" class="mx-auto mt-5">
    <v-card-title>
      <h1 class="display-1">ログイン</h1>
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field prepend-icon="mdi-account-circle" label="ユーザ名" />
        <v-text-field v-bind:type="showPassword ? 'text' : 'password'" prepend-icon="mdi-lock" v-bind:append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" label="パスワード" @click:append="showPassword = !showPassword" />
          <v-card-actions>
            <v-btn>ログイン</v-btn>
         </v-card-actions>
      </v-form>
    </v-card-text>
    </v-card>

    </v-container>
  </div>
</template>

<script>
export default {
 
  
  data: () => ({
    showPassword : false
  })
};

</script>