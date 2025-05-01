<template>
  <div class="flex flex-col items-center justify-center h-screen bg-gray-900 text-white animate__animated animate__fadeIn">
    <h1 class="text-4xl font-bold mb-4">Waiting for opponent...</h1>
    <div class="loader mb-4"></div>
    <p class="text-gray-400">Room ID: <span class="font-mono text-green-400">{{ roomId }}</span></p>
  </div>
</template>

<script lang="ts">
import { onMounted, ref } from 'vue'

export default {
  props: {
    roomId: {
      type: String,
      required: true,
    },
    playerId: {
      type: String,
      required: true,
    },
    token: {
      type: String,
      required: true,
    }
  },
  emits: ['game-start'],
  setup(props, { emit }) {
    const webSocket = ref<WebSocket>()

    onMounted(() => {
      const backendURL = import.meta.env.VITE_BACKEND_URL.replace(/^https?:\/\//, "")
      const wsProtocol = window.location.protocol === "https:" ? "wss" : "ws"
      const ws = new WebSocket(`${wsProtocol}://${backendURL}/join/${props.roomId}`)

      ws.onopen = () => {
        const joinMessage = {
          type: "join",
          body: JSON.stringify({ game_id: props.roomId }),
          sender: localStorage.getItem("playerId"),
        }
        console.log("joinMessage: ", joinMessage)
        ws.send(JSON.stringify(joinMessage))
      }

      ws.onmessage = (e) => {
      const message = JSON.parse(e.data)
      console.log('WebSocket Message:', message)

      if (message.type === 'start') {
        emit('game-start', { ws })
      }

      ws.onerror = (err) => {
        console.error('WebSocket error ❌', err)
      }

      ws.onclose = () => {
        console.log('WebSocket closed 🚪')
      }

      if (message.type === 'players_ready') {
        console.log('Two players are in the room, sending start signal...')

        const startMessage = {
          type: "start",
          body: "",
          sender: props.playerId,
        }
        ws.send(JSON.stringify(startMessage))
      }
    }


      ws.onerror = (err) => {
        console.error('WebSocket error', err)
      }

      ws.onclose = () => {
        console.log('WebSocket closed')
      }

      webSocket.value = ws
    })

    return { }
  }
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

.loader {
  border: 8px solid #f3f3f3;
  border-top: 8px solid #4ade80;
  border-radius: 50%;
  width: 80px;
  height: 80px;
  animation: spin 1s linear infinite;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
