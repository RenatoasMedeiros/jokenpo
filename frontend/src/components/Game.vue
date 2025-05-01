<template>
  <div class="h-screen flex flex-col justify-center items-center bg-gray-900 text-white animate__animated animate__fadeIn">
    <h1 class="text-4xl font-bold mb-8">Choose Your Move!</h1>

    <!-- Countdown Timer -->
    <div v-if="countdown !== null" class="text-2xl mb-4 animate__animated animate__fadeInDown">
      ⏱️ Time Left: {{ countdown }}
    </div>

    <!-- Move Buttons -->
    <div class="flex space-x-8" :class="{ 'opacity-50 pointer-events-none': buttonsDisabled }">
      <button
        v-for="move in moves"
        :key="move.name"
        @click="sendMove(move.name)"
        class="move-btn"
        :disabled="buttonsDisabled"
      >
        <div class="text-6xl">{{ move.icon }}</div>
        <div class="mt-2 text-xl">{{ move.name }}</div>
      </button>
    </div>

    <!-- Result Message -->
    <div v-if="resultMessage" class="mt-10 text-3xl font-bold animate__animated animate__bounceIn">
      {{ resultMessage }}
    </div>

    <!-- Play Again Button -->
    <button
      v-if="gameOver"
      @click="createNewRoom"
      class="mt-8 bg-green-500 hover:bg-green-600 text-white px-6 py-3 rounded-lg text-xl animate__animated animate__fadeInUp"
    >
      🔄 Play Again
    </button>

    <!-- Confetti Canvas -->
    <canvas v-if="showConfetti" ref="confettiCanvas" class="fixed top-0 left-0 w-full h-full pointer-events-none z-50"></canvas>
  </div>
</template>

<script lang="ts">
import { ref, onMounted, watch, nextTick } from 'vue'
import confetti from 'canvas-confetti'

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

    const countdown = ref<string | null>(null)
    const resultMessage = ref<string | null>(null)
    const buttonsDisabled = ref(false)
    const gameOver = ref(false)
    const showConfetti = ref(false)
    const confettiCanvas = ref<HTMLCanvasElement | null>(null)

    function sendMove(move: string) {
      if (buttonsDisabled.value) return
      const message = {
        type: "move",
        body: JSON.stringify({
          game_id: props.roomId,
          move: move
        }),
        sender: localStorage.getItem("playerId")
      }
      props.webSocket.send(JSON.stringify(message))
    }

    function launchConfetti() {
      if (!confettiCanvas.value) return
      const myConfetti = confetti.create(confettiCanvas.value, { resize: true, useWorker: true })
      myConfetti({
        particleCount: 150,
        spread: 80,
        origin: { y: 0.6 }
      })
    }

    function createNewRoom() {
      window.location.reload()
    }

    props.webSocket.onmessage = (e) => {
      const message = JSON.parse(e.data)
      console.log('WebSocket Game Message:', message)

      switch (message.type) {
        case 'countdown':
          countdown.value = message.body
          break
        case 'gameover':
          gameOver.value = true
          buttonsDisabled.value = true
          countdown.value = null

          const winnerId = message.body.replace("winner: ", "")
          const currentId = localStorage.getItem("playerId")

          if (message.body === "draw") {
            resultMessage.value = "🤝 It's a draw!"
          } else if (winnerId === currentId) {
            resultMessage.value = "🏆 You won!"
            showConfetti.value = true
            nextTick(() => launchConfetti())
          } else {
            resultMessage.value = "😔 You lost!"
          }
          break
      }
    }

    return {
      moves,
      countdown,
      resultMessage,
      sendMove,
      buttonsDisabled,
      gameOver,
      createNewRoom,
      showConfetti,
      confettiCanvas
    }
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
