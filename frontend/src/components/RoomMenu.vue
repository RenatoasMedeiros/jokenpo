<template>
  <div class="flex flex-col items-center space-y-6 p-8">
    <h2 class="text-3xl font-bold text-white">🎮 Join the Battle!</h2>

    <div class="flex space-x-10">
      <button @click="createRoom" class="menu-button bg-green-500 hover:bg-green-600">
        Create Room
      </button>

      <div class="flex flex-col items-center space-y-2">
        <input
          v-model="roomId"
          placeholder="Enter Room ID"
          class="border p-2 rounded-md w-64 text-black"
        />
        <button @click="joinRoom" class="menu-button bg-blue-500 hover:bg-blue-600">
          Join Room
        </button>
      </div>
    </div>

    <div v-if="createdRoomId" class="text-green-300 mt-6">
      Room created! ID: <span class="font-mono">{{ createdRoomId }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import { createRoom } from '../services/api'

export default {
  emits: ['room-created', 'room-joined'],
  setup(_, { emit }) {
    const roomId = ref('')
    const createdRoomId = ref('')

    async function createRoomHandler() {
      const res = await createRoom()
      createdRoomId.value = res.data.toString()
      emit('room-created', createdRoomId.value)
    }

    function joinRoom() {
      if (roomId.value.trim() !== '') {
        emit('room-joined', roomId.value.trim())
      }
    }

    return {
      roomId,
      createdRoomId,
      createRoom: createRoomHandler,
      joinRoom,
    }
  },
}
</script>

<style scoped>
.menu-button {
  @apply text-white font-bold py-3 px-6 rounded-lg transition-all duration-200 transform hover:scale-105;
}
</style>
