<script setup lang="ts">
import { defineProps } from 'vue'

interface JobPosting {
  Title: string
  Company: string
  Location: string
  URL: string
}

const props = defineProps<{ jobs: JobPosting[] }>()
</script>

<template>
  <div v-if="props.jobs.length">
    <h2>Job Matches:</h2>
    <ul>
      <li v-for="job in props.jobs" :key="job.URL">
        <a :href="job.URL" target="_blank">{{ job.Title }} @ {{ job.Company }} ({{ job.Location }})</a>
      </li>
    </ul>
  </div>
  <div v-else>
    <p>No jobs found for your search criteria.</p>
  </div>
</template>
