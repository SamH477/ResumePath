<!-- frontend/src/App.vue -->
<script setup lang="ts">
import JobMatcher from './components/jobMatcher.vue'
</script>

<template>
  <div id="app">
    <h1>ResumePath</h1>
    <JobMatcher />
  </div>
</template>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 2rem;
}
</style>
