<template>
  <div class="Todo">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <Title titlename="やるべきこと" size="12" />
        <!-- {{ this.$store.state.todos }} -->
        <v-col cols="5" sm="6" md="4" lg="3">
          <v-dialog v-model="dialog" max-width="600px">
            <template v-slot:activator="{ on, attrs }">
              <v-btn color="#F29993" dark v-bind="attrs" v-on="on" fab x-large
                >➕</v-btn
              >
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">Todoリスト</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12">
                      <v-text-field
                        :rules="toRules"
                        label="やること"
                        required
                        v-model="title"
                      ></v-text-field>
                    </v-col>

                    <v-col cols="12">
                      <!-- <v-text-field label="いつまで" required v-model="date"></v-text-field> -->

                      日付： {{ date }}
                      <DatePicker v-model="date" />
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-select
                        :items="[
                          '国語',
                          '算数',
                          '理科',
                          '社会',
                          '英語',
                          'その他',
                        ]"
                        label="科目"
                        required
                        v-model="subject"
                        :rules="subjetRules"
                      ></v-select>
                    </v-col>
                  </v-row>
                </v-container>
                <!-- <small>*indicates required field</small> -->
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="dialog = false"
                  >閉じる</v-btn
                >
                <v-btn
                  color="blue darken-1"
                  text
                  v-on:click="store_add_todo()"
                  :disabled="!valid"
                  >保存する</v-btn
                >
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <v-sheet color="grey lighten-3" height="424">
        <v-row class="#F29993" justify="center" align-content="center">
          <v-col
            cols="12"
            sm="10"
            md="8"
            lg="4"
            xl="3"
            v-for="(one_todo, i) in this.$store.state.todos"
            :key="i"
          >
            <TodoCard :todo="one_todo" />
          </v-col>
        </v-row>
        <!-- <v-row
          class="blue lighten-4"
          style="height: 450px;"
          justify="center"
          align-content="center"
        >
          <v-col v-for="n in 3" :key="n" cols="12" sm="10" md="8" lg="4" xl="3">
            <v-card color="blue" outlined tile height="150"></v-card>
          </v-col>
        </v-row> -->
      </v-sheet>
    </v-container>
    <v-container>
       <!-- <v-bottom-navigation
      horizontal
    >
      <router-link to="/">
        <v-btn dark > 
          <span>ホーム</span>
          
        </v-btn>
      </router-link>
       </v-bottom-navigation> -->
       <Menubar btnname="ホーム"></Menubar>
    </v-container>
    {{ subject_num }}
  </div>
</template>

<script>
// @ is an alias to /src
import Title from "@/components/Title.vue";
import TodoCard from "@/components/TodoCard.vue";
import DatePicker from "@/components/DatePicker";
import Menubar from "@/components/Menubar.vue";
export default {
  components: {
    Title,
    TodoCard,
    DatePicker,
    Menubar
  },
  data: () => ({
    valid: true,
    title: "",
    subject_num: 0,
    toRules: [(v) => !!v || "やることを入力してね"],
    select: null,
    items: ["国語", "算数", "理科", "社会", "英語", "その他"],
    date: "",
    dataRules: [(v) => !!v || "日付を選択してね"],
    subject: "",
    subjetRules: [(v) => !!v || "科目を選択してね"],
    dialog: false,
    // テスト
  }),
  methods: {
    store_add_todo() {
      switch (this.subject) {
        case "国語":
          this.subject_num = 1;
          break;
        case "数学":
          this.subject_num = 2;
          break;
        case "英語":
          this.subject_num = 3;
          break;
        case "理科":
          this.subject_num = 4;
          break;
        case "社会":
          this.subject_num = 5;
          break;
        default:
          this.subject_num = 6;
      }
      if (this.title && this.date && this.subject) {
        this.$store.commit("add_todo", {
          name: this.title,
          date: this.date,
          subject: this.subject_num,
        });
        this.$store.dispatch("post_new_todo", {
          title: this.title,
          date: this.date,
          subject: this.subject_num,
        });
        this.title = "";
        this.date = "";
        this.subject = "";
        this.subject_num = 0;
        this.dialog = false;
      }
    },
    validate() {
      this.$refs.form.validate();
    },
  },
};
</script>
<style>
h1 {
  color: #6b421d !important;
}
</style>
