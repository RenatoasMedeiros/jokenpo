<template>
  <form @submit.prevent="submit" class="space-y-4 max-w-md mx-auto p-6 rounded-lg bg-gray-800 shadow-lg animate__animated animate__fadeIn">
    <h1 class="text-2xl font-bold text-white text-center mb-4">Create your account</h1>
    <input v-model="username" placeholder="Username" required class="w-full p-3 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-green-500" />
    <input v-model="password" type="password" placeholder="Password" required class="w-full p-3 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-green-500" />
    <button type="submit" class="w-full p-3 rounded bg-green-500 hover:bg-green-600 text-white font-bold transition-transform transform hover:scale-105">
      Register
    </button>
  </form>
</template>

<script lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { register } from '../services/api'
import { toast } from 'vue3-toastify' // if you're using vue3-toastify or your toaster
import 'vue3-toastify/dist/index.css' // important: toastify styles

export default {
  setup() {
    const username = ref('')
    const password = ref('')
    const router = useRouter()

    async function submit() {
      try {
        const res = await register(username.value, password.value)
        toast.success('Registration successful! Redirecting to login...', {
          position: "top-center",
          autoClose: 3000,
        })
        setTimeout(() => {
          router.push('/login')
        }, 3000) // Wait a bit so user sees success before moving
      } catch (error: any) {
        toast.error(error.response?.data?.error || 'Registration failed. Try again.', {
          position: "top-center",
        })
      }
    }

    return { username, password, submit }
  },
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css'; /* Animate.css CDN */
</style>
