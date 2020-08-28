<template>
  <div class="SingUp">
    <v-container>
      <v-card width="400px" class="mx-auto mt-5">
        <v-card-title>
          <h1 class="display-1">登録</h1>
        </v-card-title>
        <v-card-text>
          <v-form>
            <v-text-field
              v-model="user_name"
              prepend-icon="mdi-account-circle"
              label="ユーザ名"
            />
            <v-text-field
              v-model="password"
              v-bind:type="showPassword ? 'text' : 'password'"
              prepend-icon="mdi-lock"
              v-bind:append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              label="パスワード"
              @click:append="showPassword = !showPassword"
            />
            <v-card-actions>
              <v-btn v-on:click="signup()">登録</v-btn>
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
    user_name: "",
    password: "",
    showPassword: false,
  }),
  methods: {
    signup() {
      this.$store.dispatch("post_new_user", {
        name: this.user_name,
        password: this.password,
      });
      // if (this.$store.state.user.logined == false) {
      //   setTimeout(this.enable, 1500);
      // }
    },
  },
};
</script>
