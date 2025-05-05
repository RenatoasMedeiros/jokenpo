<!-- <template>
  <div class="space-y-4 max-w-md mx-auto p-6 rounded-lg bg-gray-800 shadow-lg animate__animated animate__fadeIn">
    <h1 class="text-2xl font-bold text-white text-center mb-4">Create a Game Room</h1>
    <button
      @click="create"
      class="w-full p-3 rounded bg-purple-600 hover:bg-purple-700 text-white font-bold transition-transform transform hover:scale-105 focus:outline-none"
    >
      Create Room
    </button>
    <div v-if="roomId" class="mt-4 p-4 bg-gray-700 text-white rounded-lg text-center animate__animated animate__fadeIn">
      🎮 Room Created! <br> Room ID: <span class="font-mono text-green-400">{{ roomId }}</span>
    </div>
  </div>
</template> -->

<template>
  <div class="room-creator-container space-y-4 animate__animated animate__fadeIn">
    <h2 class="lobby-section-title">Create a Game Room</h2>
    <button
      @click="create"
      class="lobby-button create" >
      Create Room
    </button>
    <div v-if="roomId" class="room-id-display animate__animated animate__fadeIn">
      🎮 Room Created! <br> Room ID: <span>{{ roomId }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import { createRoom } from '../services/api'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

export default {
  emits: ['room-created'],
  setup(_, { emit }) {
    const roomId = ref('')

    async function create() {
      try {
        const res = await createRoom()
        roomId.value = res.data.toString()
        emit('room-created', roomId.value)
        toast.success('Room created successfully! 🎉', {
          position: 'top-center',
          autoClose: 2000,
        })
      } catch (error: any) {
        toast.error(error.response?.data?.error || 'Failed to create room.', {
          position: 'top-center',
        })
      }
    }

    return { roomId, create }
  },
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';
</style>
