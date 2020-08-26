<template>
  <div class="Todo">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <Title titlename="やるべきこと" size="12" />
        <!-- {{this.$store.state.todos}} -->
        <v-col cols="5" sm="6" md="4" lg="3">
          <v-dialog v-model="dialog" max-width="600px">
            <template v-slot:activator="{ on, attrs }">
              <v-btn color="#F29993" dark v-bind="attrs" v-on="on" fab x-large>➕</v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">Todoリスト</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <v-text-field label="やること" required v-model="title"></v-text-field>
                    </v-col>
                    <v-col cols="12">
                      <v-text-field label="いつまで" required v-model="date"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-select
                        :items="['国語', '算数', '理科', '社会', '英語', 'その他']"
                        label="科目"
                        required
                        v-model="subject"
                      ></v-select>
                    </v-col>
                  </v-row>
                </v-container>
                <!-- <small>*indicates required field</small> -->
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="dialog = false">閉じる</v-btn>
                <v-btn color="blue darken-1" text v-on:click="store_add_todo()">保存する</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <v-sheet class="d-flex" color="grey lighten-3" height="424">
        <v-col v-for="(one_todo,i) in this.$store.state.todos" :key="i">
          <TodoCard :todo="one_todo" />
        </v-col>
      </v-sheet>
    </v-container>
    <v-container>
      <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Calendar">
        <v-btn color="green" dark>カレンダー</v-btn>
      </router-link>
      <router-link to="/Raising">
        <v-btn color="green" dark>育成</v-btn>
      </router-link>
    </v-container>
  </div>
</template>


<script>
// @ is an alias to /src
import Title from "@/components/Title.vue";
import TodoCard from "@/components/TodoCard.vue";

export default {
  components: {
    Title,
    TodoCard,
  },
  data: () => ({
    title: "",
    date: "",
    subject: "",
    dialog: false,
  }),
  methods: {
    store_add_todo() {
      this.$store.commit("add_todo", {
        title: this.title,
        date: this.date,
        subject: this.subject,
      });
      this.dialog = false;
    },
  },
};
</script>
<style>
h1 {
  color: #6b421d !important;
}
</style>
