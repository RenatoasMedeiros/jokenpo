<template>
  <div class="space-y-6 max-w-md mx-auto p-6 rounded-lg bg-gray-800 shadow-lg animate__animated animate__fadeIn">
    <h2 class="text-2xl font-bold text-white text-center mb-4">Available Rooms</h2>

    <!-- List available rooms -->
    <div v-if="rooms.length > 0" class="space-y-2">
      <div
        v-for="room in rooms"
        :key="room.room_id"
        class="flex justify-between items-center p-3 bg-gray-700 rounded hover:bg-gray-600 transition">
        <span class="text-white">{{ room.room_id }}</span>
        <button @click="joinRoom(room.room_id)" class="bg-green-500 hover:bg-green-600 text-white px-3 py-1 rounded">
          Join
        </button>
      </div>
    </div>

    <div v-else class="text-gray-400 text-center">
      No rooms available.
    </div>

    <!-- Button to show manual join -->
    <div class="text-center">
      <button
        @click="showManualJoin = !showManualJoin"
        class="mt-4 bg-purple-500 hover:bg-purple-600 text-white px-4 py-2 rounded transition-transform transform hover:scale-105"
      >
        {{ showManualJoin ? 'Cancel' : 'Join with Room ID' }}
      </button>
    </div>

    <!-- Manual Room ID Input -->
    <div v-if="showManualJoin" class="mt-4 space-y-2 animate__animated animate__fadeIn">
      <input
        v-model="manualRoomId"
        placeholder="Enter Room ID..."
        class="w-full p-3 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-purple-500"
      />
      <button
        @click="joinRoom(manualRoomId)"
        class="w-full p-3 rounded bg-blue-500 hover:bg-blue-600 text-white font-bold transition-transform transform hover:scale-105"
      >
        Join Room
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'
import { toast } from 'vue3-toastify'
import { getRooms } from '../services/api'
import 'vue3-toastify/dist/index.css'

export default {
  emits: ['room-joined'],
  setup(_, { emit }) {
    const rooms = ref<{ room_id: string }[]>([])
    const showManualJoin = ref(false)
    const manualRoomId = ref('')

    async function fetchRooms() {
      try {
        const res = await getRooms()
        rooms.value = res.data
      } catch (error: any) {
        toast.error('Failed to load rooms', { position: 'top-center' })
      }
    }

    function joinRoom(id: string) {
      if (!id.trim()) {
        toast.error('Please enter a valid Room ID', { position: 'top-center' })
        return
      }
      emit('room-joined', id.trim())
      toast.success('Joining room...', { position: 'top-center', autoClose: 1000 })
    }

    onMounted(() => {
      // fetchRooms()
    })

    return { rooms, showManualJoin, manualRoomId, joinRoom }
  }
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';
</style>
