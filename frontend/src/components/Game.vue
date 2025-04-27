<template>
  <div class="h-screen flex flex-col justify-center items-center bg-gray-900 text-white animate__animated animate__fadeIn">
    <h1 class="text-4xl font-bold mb-8">Choose Your Move!</h1>

    <div class="flex space-x-8">
      <button
        v-for="move in moves"
        :key="move.name"
        @click="sendMove(move.name)"
        class="move-btn"
      >
        <div class="text-6xl">{{ move.icon }}</div>
        <div class="mt-2 text-xl">{{ move.name }}</div>
      </button>
    </div>

    <div v-if="result" class="mt-10 text-3xl font-bold animate__animated animate__bounceIn">
      {{ resultMessage }}
    </div>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'

export default {
  props: {
    webSocket: {
      type: Object as () => WebSocket,
      required: true
    },
    roomId: {
      type: String,
      required: true
    },
    userId: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const moves = [
      { name: "rock", icon: "🪨" },
      { name: "paper", icon: "📄" },
      { name: "scissors", icon: "✂️" }
    ]

    const result = ref<string | null>(null)

    function sendMove(move: string) {
      const message = {
        type: "move",
        body: JSON.stringify({
          game_id: props.roomId,
          move: move
        }),
        sender: props.userId
      }
      props.webSocket.send(JSON.stringify(message))
    }

    props.webSocket.onmessage = (e) => {
      const message = JSON.parse(e.data)
      console.log('WebSocket Game Message:', message)

      if (message.type === 'result') {
        if (message.body === 'player1') {
          result.value = '🏆 You won!'
        } else if (message.body === 'player2') {
          result.value = '😔 You lost!'
        } else if (message.body === 'draw') {
          result.value = '🤝 It\'s a draw!'
        }
      }
    }

    return { moves, sendMove, result, resultMessage: result }
  }
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

.move-btn {
  background: linear-gradient(135deg, #4ade80, #22d3ee);
  border: 2px solid #22d3ee;
  border-radius: 1rem;
  padding: 2rem;
  transition: all 0.3s;
  width: 160px;
  height: 160px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-weight: bold;
}

.move-btn:hover {
  transform: scale(1.1);
  background: linear-gradient(135deg, #22d3ee, #4ade80);
  cursor: pointer;
}
</style>
