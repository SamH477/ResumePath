<script setup lang="ts">
import { ref } from 'vue'
import ResumeForm from './ResumeForm.vue'
import ResumeDownload from './JobMatches.vue'

const downloadUrl = ref('')

async function handleSubmit(fileFormData: FormData) {
  try {
    const response = await fetch('http://localhost:8080/api/matches', {
      method: 'POST',
      body: fileFormData
    })

    const blob = await response.blob()
    downloadUrl.value = URL.createObjectURL(blob)
  } catch (error) {
    console.error('Error tailoring resume:', error)
  }
}
</script>

<template>
  <div class="job-matcher">
    <ResumeForm @submit="handleSubmit" />
    <ResumeDownload :fileUrl="downloadUrl" />
  </div>
</template>

<style scoped>
.job-matcher {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}
</style>
