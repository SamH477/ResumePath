<script setup lang="ts">
import { ref } from 'vue'
import ResumeForm from './ResumeForm.vue'
import Matches from './JobMatches.vue'

interface JobPosting {
  Title: string
  Company: string
  Location: string
  URL: string
}

const jobPostings = ref<JobPosting[]>([])

async function handleSubmit(fileFormData: FormData) {
  try {
    const response = await fetch('http://localhost:8080/api/matches', {
      method: 'POST',
      body: fileFormData
    })

    const jobs: JobPosting[] = await response.json()
    jobPostings.value = jobs // update reactive state
  } catch (error) {
    console.error('Error scraping jobs:', error)
  }
}
</script>

<template>
  <div class="job-matcher">
    <ResumeForm @submit="handleSubmit" />
    <Matches :jobs="jobPostings" />
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
