<template>
  <div class="Calender">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <v-row class="text-center">
        <v-col cols="7" sm="7" md="5" lg="4">
          <h1>カレンダー</h1>
          <v-divider></v-divider>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <v-row>
        <v-col>
          <v-sheet height="64">
            <v-toolbar flat color="white">
              <v-btn
                outlined
                class="mr-4"
                color="grey darken-2"
                @click="setToday"
                >今日</v-btn
              >
              <v-spacer></v-spacer>

              <v-btn
                outlined
                class="ma-4"
                fab
                text
                small
                color="grey darken-2"
                @click="prev"
              >
                <v-icon small>前月</v-icon>
              </v-btn>

              <v-btn
                outlined
                fab
                text
                small
                color="grey darken-2"
                @click="next"
              >
                <v-icon small>次月</v-icon>
              </v-btn>

              <!-- <v-toolbar-title v-if="$refs.calendar">
                {{ $refs.calendar.title }}
              </v-toolbar-title>
              <v-spacer></v-spacer>-->
              <!-- <v-menu bottom right>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    outlined
                    color="grey darken-2"
                    v-bind="attrs"
                    v-on="on"
                  >
                    <span>{{ typeToLabel[type] }}</span>
                    <v-icon right>mdi-menu-down</v-icon>
                  </v-btn>
                </template>
                <v-list>
                  <v-list-item @click="type = 'day'">
                    <v-list-item-title>1日</v-list-item-title>
                  </v-list-item>
                  <v-list-item @click="type = 'week'">
                    <v-list-item-title>1週間</v-list-item-title>
                  </v-list-item>
                  <v-list-item @click="type = 'month'">
                    <v-list-item-title>1ヶ月</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>-->
            </v-toolbar>
          </v-sheet>
          <v-sheet height="460">
            <v-calendar
              ref="calendar"
              v-model="focus"
              color="primary"
              :events="events"
              :event-color="getEventColor"
              @click:date="display"
              locale="ja-jp"
              :day-format="(timestamp) => new Date(timestamp.date).getDate()"
              :month-format="
                (timestamp) => new Date(timestamp.date).getMonth() + 1 + ' /'
              "
            ></v-calendar>
            <v-dialog v-model="dialog" width="500">
              <v-card>
                <v-col>
                  <v-row
                    class="#F29993"
                    justify="center"
                    align-content="center"
                  >
                    <v-col
                      cols="12"
                      sm="10"
                      md="8"
                      lg="4"
                      xl="3"
                      v-for="(todo_name, i) in today_list"
                      :key="i"
                    >
                      <v-card height="100" class="name">
                        {{ todo_name }}
                      </v-card>
                    </v-col>
                  </v-row>
                </v-col>

                <v-btn
                  height="50"
                  width="100%"
                  color="primary"
                  @click="dialog = false"
                  >戻る</v-btn
                >
              </v-card>
            </v-dialog>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <!-- <router-link to="/">
        <v-btn color="green" dark>ホーム</v-btn>
      </router-link>
      <router-link to="/Todo">
        <v-btn color="green" dark>Todoリスト</v-btn>
      </router-link>
      <router-link to="/Raising">
        <v-btn color="green" dark>育成</v-btn>
      </router-link> -->
      <Menubar btnname="ホーム"></Menubar>
    </v-container>
    {{ today_list }}
  </div>
</template>

<script>
// @ is an alias to /src
import Menubar from "@/components/Menubar.vue";

export default {
  name: "",
  components: {
    Menubar,
  },
  data: () => ({
    day_list: [],
    lists: [],
    focus: "",
    dialog: false,
    type: "month",
    typeToLabel: {
      month: "1ヶ月",
      week: "1週間",
      day: "1日",
    },
    today_list: [],
    selectedEvent: {},
    selectedElement: null,
    selectedOpen: false,
    events: [],
    colors: [
      "blue",
      "indigo",
      "deep-purple",
      "cyan",
      "green",
      "orange",
      "grey darken-1",
    ],
    names: [
      "Meeting",
      "Holiday",
      "PTO",
      "Travel",
      "Event",
      "Birthday",
      "Conference",
      "Party",
    ],
  }),
  mounted() {
    // this.$refs.calendar.checkChange();
  },
  created() {
    this.$store.state.todos.forEach((todo) => {
      this.day_list.push(todo.date);
    });
  },
  methods: {
    display({ date }) {
      this.today_list = [];
      this.$store.state.todos.forEach((todo) => {
        if (todo.date === date) {
          this.today_list.push(todo.name);
        }
      });
      if (this.today_list.length > 0) {
        this.dialog = true;
      }
    },
    getEventColor(event) {
      return event.color;
    },
    setToday() {
      this.focus = "";
    },
    prev() {
      this.$refs.calendar.prev();
    },
    next() {
      this.$refs.calendar.next();
    },
    showEvent({ nativeEvent, event }) {
      const open = () => {
        this.selectedEvent = event;
        this.selectedElement = nativeEvent.target;
        setTimeout(() => (this.selectedOpen = true), 10);
      };

      if (this.selectedOpen) {
        this.selectedOpen = false;
        setTimeout(open, 10);
      } else {
        open();
      }
      nativeEvent.stopPropagation();
    },

    // updateRange ({ start, end }) {
    //   const events = []

    //   const min = new Date(`${start.date}T00:00:00`)
    //   const max = new Date(`${end.date}T23:59:59`)
    //   const days = (max.getTime() - min.getTime()) / 86400000
    //   const eventCount = this.rnd(days, days + 20)

    //   for (let i = 0; i < eventCount; i++) {
    //     const allDay = this.rnd(0, 3) === 0
    //     const firstTimestamp = this.rnd(min.getTime(), max.getTime())
    //     const first = new Date(firstTimestamp - (firstTimestamp % 900000))
    //     const secondTimestamp = this.rnd(2, allDay ? 288 : 8) * 900000
    //     const second = new Date(first.getTime() + secondTimestamp)

    //     events.push({
    //       name: this.names[this.rnd(0, this.names.length - 1)],
    //       start: first,
    //       end: second,
    //       color: this.colors[this.rnd(0, this.colors.length - 1)],
    //       timed: !allDay,
    //     })
    //   }

    //   this.events = events
    // },
    // rnd (a, b) {
    //   return Math.floor((b - a + 1) * Math.random()) + a
    // },
  },
};
</script>
<style>
.theme--light.v-divider {
  border-color: #6b421d !important;
  border-width: 5px !important;
}
h1 {
  color: #6b421d !important;
}
.name {
  font-size: 30px;
}
</style>
