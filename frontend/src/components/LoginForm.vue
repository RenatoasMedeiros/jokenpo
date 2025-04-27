<template>
  <form @submit.prevent="submit" class="space-y-4 max-w-md mx-auto p-6 rounded-lg bg-gray-800 shadow-lg animate__animated animate__fadeIn">
    <h1 class="text-2xl font-bold text-white text-center mb-4">Welcome Back!</h1>
    <input v-model="username" placeholder="Username" required class="w-full p-3 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500" />
    <input v-model="password" type="password" placeholder="Password" required class="w-full p-3 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500" />
    <button type="submit" class="w-full p-3 rounded bg-blue-500 hover:bg-blue-600 text-white font-bold transition-transform transform hover:scale-105">
      Log In
    </button>
  </form>
</template>

<script lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login, setToken } from '../services/api'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

export default {
  emits: ['logged-in'],
  setup(_, { emit }) {
    const username = ref('')
    const password = ref('')
    const router = useRouter()

    async function submit() {
    try {
      const res = await login(username.value, password.value)

      setToken(res.data.token)
      localStorage.setItem('playerId', res.data.playerId)

      emit('logged-in', res.data)

      toast.success('Logged in successfully!', {
        position: "top-center",
        autoClose: 2000,
      })

      setTimeout(() => {
        router.push('/lobby')
      }, 2000)

    } catch (error: any) {
      toast.error(error.response?.data?.error || 'Login failed. Please check your credentials.', {
        position: "top-center",
      })
    }
  }


    return { username, password, submit }
  },
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';
</style>
