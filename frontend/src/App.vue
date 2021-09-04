<template>
  <div id="app">
    <GoodDay />
  </div>
</template>

<script>
import GoodDay from './components/GoodDay.vue'

export default {
  name: 'App',
  components: {
    GoodDay
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
