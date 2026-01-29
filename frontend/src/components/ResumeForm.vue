<script setup lang="ts">
import { ref, defineEmits } from 'vue'

const emit = defineEmits<{
  (e: 'submit', payload: FormData): void
}>()

const form = ref<{
  file: File | null
  city: string
  state: string
  job: string
}>({
  file: null,
  city: '',
  state: '',
  job:''
})

function handleFileUpload(event: Event) {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    form.value.file = target.files[0]
  }
}

function handleSubmit() {
  if (!form.value.file) {
    alert('Please upload your resume.')
    return
  }

  if (!form.value.city || !form.value.state) {
    alert('Please enter both city and state.')
    return
  }

    if (!form.value.job) {
    alert('Please enter job title.')
    return
  }

  const formData = new FormData()
  formData.append('resume', form.value.file)
  formData.append('city', form.value.city)
  formData.append('state', form.value.state)
  formData.append('job', form.value.job)

  emit('submit', formData)
}
</script>

<template>
  <form @submit.prevent="handleSubmit">
    <div>
      <label for="resume">Upload Resume:</label>
      <input
        id="resume"
        type="file"
        accept=".pdf,.doc,.docx"
        @change="handleFileUpload"
      />
    </div>

    <div>
      <label for="city">City:</label>
      <input
        id="city"
        type="text"
        v-model="form.city"
        placeholder="e.g. Lancaster"
      />
    </div>

    <div>
      <label for="state">State:</label>
      <input
        id="state"
        type="text"
        v-model="form.state"
        placeholder="e.g. PA"
      />
    </div>

    <div>
      <label for="job">Job Title:</label>
      <input
        id="job"
        type="text"
        v-model="form.job"
        placeholder="e.g. Data Analyst"
      />
    </div>

    <button type="submit">Find Matches</button>
  </form>
</template>
