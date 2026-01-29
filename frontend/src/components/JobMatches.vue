<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps<{ fileUrl: string }>()
</script>

<template>
  <div v-if="props.fileUrl">
    <a :href="props.fileUrl" download="tailored_resume.pdf">Download Tailored Resume</a>
  </div>
  <div v-else>
    <p>Your tailored resume will appear here after generation.</p>
  </div>
</template>
