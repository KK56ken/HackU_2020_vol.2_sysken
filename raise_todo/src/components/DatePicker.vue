<template>
  <v-menu v-model="menu" offset-y :close-on-content-click="false">
    <template v-slot:activator="{ on }">
      <v-btn icon color="primary" dark elevation="0" v-on="on">
        <v-icon>mdi-calendar</v-icon>
      </v-btn>
    </template>
    <v-date-picker v-model="picker" @click="menu = false" 
    locale="ja-jp" :day-format="date => new Date(date).getDate()" />
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
};
</script>

<style>
.v-date-picker-table.v-date-picker-table--date > table > tbody tr td:nth-child(7) .v-btn__content {
    color:blue
}
.v-date-picker-table.v-date-picker-table--date > table > tbody tr td:nth-child(1) .v-btn__content {
    color:red
}
</style>
