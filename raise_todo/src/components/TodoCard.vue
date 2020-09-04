<template>
  <v-row v-if="show">
    <v-col cols="12" sm="12" md="4" lg="3">
      <!-- <p>あと3日</p> -->
      <v-card class="mx-auto" height="100%" width="98%">
        <v-card-text>
          <div class="todo_name text--primary">
            {{ todo.name }}
          </div>
          <span class="end_date text--primary"> 提出日：{{ todo.date }} </span>
          <span class="h4 text--primary">科目：{{ todo.subject }}</span>
           <v-spacer></v-spacer>
          <v-btn
              dark
              class="text-right"
              color="green darken-1"
              @click.stop="dialog = true"
              >終了</v-btn
            >
          <v-dialog v-model="dialog" max-width="290">
            <v-card>
              <v-card-title class="headline">本当に消していいの？</v-card-title>

              <v-card-text>
                1回消したら戻せないよ！本当に終わった？
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>

                <v-btn
                  color="green darken-1"
                  text
                  @click="
                    card_destroy(), add_food(), (dialog = false);
                    show = !show;
                  "
                >
                  いいよ ○
                </v-btn>
                <v-btn color="red darken-1" text @click="dialog = false">
                  だめ ×
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "",
  data: () => ({
    show: true,
    dialog: false,
  }),
  methods: {
    add_food() {
      this.$store.state.now_feed++;
      
      this.$store.dispatch("post_new_food", {

        token:this.$store.state.user.token,
        flag:1,
        feednum:1
    });
    
    },
    
    card_destroy() {
      this.$destroy;
    },
  },
  props: ["todo"],
};
</script>

<style>
.todo_name {
  padding-top: 10px;
  padding-left: 0;
  padding-bottom: 10px;
  font-size: 30px;
}
.end_date {
  font-size: 18px;
  padding-right: 20px;
}
</style>
