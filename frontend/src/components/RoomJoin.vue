<!-- <template>
  <div class="space-y-6 max-w-md mx-auto p-6 rounded-lg bg-gray-800 shadow-lg animate__animated animate__fadeIn">
    <h2 class="text-2xl font-bold text-white text-center mb-4">Enter in a Room</h2>
    <div class="mt-4 space-y-2 animate__animated animate__fadeIn">
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
</template> -->

<template>
  <div class="room-join-container space-y-4 animate__animated animate__fadeIn">
    <h2 class="lobby-section-title">Enter a Room</h2>

    <div class="space-y-2 animate__animated animate__fadeIn">
      <input
        v-model="manualRoomId"
        placeholder="Enter Room ID..."
        class="lobby-input" />
      <button
        @click="joinRoom(manualRoomId)"
        class="lobby-button" >
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
