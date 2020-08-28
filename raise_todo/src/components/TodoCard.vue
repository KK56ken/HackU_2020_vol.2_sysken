<template>
  <v-row v-if="show"> 
    <v-col cols="12" sm="12" md="4" lg="3">
      <!-- <p>あと3日</p> -->
      <v-card class="mx-auto" height="80" width="350">
        <v-card-text>
          <span class="display-1 text--primary">{{todo.name}}</span>
          <span class="h4 text--primary">科目：{{todo.subject}}</span>
          <v-btn class="text-right" text color="deep-purple accent-4" @click.stop="dialog = true">終了</v-btn>
          <v-dialog
            v-model="dialog"
            max-width="290"
          >
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
                  @click="dialog = false; show=!show"
                >
                  いいよ ○
                </v-btn>
                <v-btn
                  color="red darken-1"
                  text
                  @click="dialog = false"
                >
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
  props: ["todo"],
};
</script>
