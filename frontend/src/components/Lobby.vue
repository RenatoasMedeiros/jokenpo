<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gray-900 text-white space-y-10 p-8 animate__animated animate__fadeIn">
    <h1 class="text-4xl font-bold mb-6">🎮 Game Lobby</h1>

    <!-- Create Room -->
    <RoomCreator @room-created="emitRoomCreated" />

    <!-- Divider -->
    <div class="border-t border-gray-700 w-full my-6"></div>

    <!-- Join Room -->
    <RoomJoin @room-joined="emitRoomJoined" />
  </div>
</template>

<script lang="ts">
import RoomCreator from './RoomCreator.vue'
import RoomJoin from './RoomJoin.vue'
import { useRouter } from 'vue-router'
import { defineComponent } from 'vue'

import { inject } from 'vue'

export default defineComponent({
  components: { RoomCreator, RoomJoin },
  emits: ['room-created', 'room-joined'], // ADD THIS
  setup(_, { emit }) {
    function emitRoomCreated(roomId: string) {
      emit('room-created', roomId)
    }
    function emitRoomJoined(roomId: string) {
      emit('room-joined', roomId)
    }

    return { emitRoomCreated, emitRoomJoined }
  }
})

</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';
</style>
