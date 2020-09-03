<template>
  <v-menu v-model="menu" offset-y :close-on-content-click="false">
    <template v-slot:activator="{ on }">
      <v-btn icon color="primary" dark elevation="0" v-on="on">
        <v-icon>mdi-calendar</v-icon>
      </v-btn>
    </template>
    <v-date-picker
      :min="created()"
      v-model="picker"
      @click="menu = false"
      locale="ja-jp"
      :day-format="(date) => new Date(date).getDate()"
    />
  </v-menu>
</template>
<script>
export default {
  props: {
    date: {
      type: String,
      default: new Date().toISOString().substr(0, 10),
    },
  },
  data() {
    return {
      menu: false,
      today: new Date(),
    };
  },
  computed: {
    picker: {
      get() {
        return this.date;
      },
      set(val) {
        this.menu = false;
        this.$emit("input", val);
      },
    },
  },
  created() {
    var a = 0;
    if (
      String(this.$store.state.month).length === 1 &&
      String(this.$store.state.today).length === 1
    ) {
      a =
        this.$store.state.year +
        "-" +
        "0" +
        this.$store.state.month +
        "-" +
        "0" +
        this.$store.state.today;
    } else if (String(this.$store.state.month).length === 1) {
      a =
        this.$store.state.year +
        "-" +
        "0" +
        this.$store.state.month +
        "-" +
        this.$store.state.today;
    } else if (String(this.$store.state.today).length === 1) {
      a =
        this.$store.state.year +
        "-" +
        this.$store.state.month +
        "-" +
        "0" +
        this.$store.state.today;
    } else {
      a =
        this.$store.state.year +
        "-" +
        this.$store.state.month +
        "-" +
        this.$store.state.today;
    }
    return a;
  },
};
</script>

<style>
.v-date-picker-table.v-date-picker-table--date
  > table
  > tbody
  tr
  td:nth-child(7)
  .v-btn__content {
  color: blue;
}
.v-date-picker-table.v-date-picker-table--date
  > table
  > tbody
  tr
  td:nth-child(1)
  .v-btn__content {
  color: red;
}
</style>
