<template>
  <div class="game-container">
    <!-- Animated background particles -->
    <div class="particles-container">
      <div v-for="n in 20" :key="n" class="particle"
          :style="{
            left: `${Math.random() * 100}%`,
            top: `${Math.random() * 100}%`,
            width: `${Math.random() * 10 + 5}px`,
            height: `${Math.random() * 10 + 5}px`,
            animationDuration: `${Math.random() * 20 + 10}s`,
            animationDelay: `${Math.random() * 5}s`
          }">
      </div>
    </div>

    <div class="game-content animate__animated animate__fadeIn">
      <div class="game-header">
        <h1>Jokenpo Realm</h1>
        <h2 class="battle-title">⚔️ Battle Arena ⚔️</h2>
      </div>

      <!-- Countdown Timer -->
      <div v-if="countdown !== null" class="countdown-timer animate__animated animate__pulse">
        <div class="timer-icon">⏱️</div>
        <div class="timer-value">{{ countdown }}</div>
      </div>

      <!-- Move Buttons -->
      <div class="moves-container" :class="{ 'disabled': buttonsDisabled }">
        <button
          v-for="move in moves"
          :key="move.name"
          @click="sendMove(move.name)"
          class="move-btn animate__animated animate__fadeInUp"
          :disabled="buttonsDisabled"
          :style="{ animationDelay: `${moves.indexOf(move) * 0.2}s` }"
        >
          <div class="move-icon">{{ move.icon }}</div>
          <div class="move-name">{{ move.name }}</div>
        </button>
      </div>

      <!-- Result Message -->
      <div v-if="resultMessage" class="result-message animate__animated animate__bounceIn">
        {{ resultMessage }}
      </div>

      <!-- Play Again Button -->
      <button
        v-if="gameOver"
        @click="createNewRoom"
        class="play-again-btn animate__animated animate__fadeInUp"
      >
        🔄 Play Again
      </button>

      <!-- Confetti Canvas -->
      <canvas v-if="showConfetti" ref="confettiCanvas" class="confetti-canvas"></canvas>
    </div>
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
      buttonsDisabled.value = true

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

      // Multiple bursts of confetti
      const duration = 3 * 1000;
      const end = Date.now() + duration;

      (function frame() {
        myConfetti({
          particleCount: 2,
          angle: 60,
          spread: 55,
          origin: { x: 0, y: 0.6 },
          colors: ['#4f6cff', '#832eff', '#22d3ee', '#4ade80']
        });

        myConfetti({
          particleCount: 2,
          angle: 120,
          spread: 55,
          origin: { x: 1, y: 0.6 },
          colors: ['#4f6cff', '#832eff', '#22d3ee', '#4ade80']
        });

        if (Date.now() < end) {
          requestAnimationFrame(frame);
        }
      }());
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

.game-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: radial-gradient(circle at center, #1a1b4b 0%, #0d0e24 70%);
  position: relative;
  overflow: hidden;
  padding: 20px;
}

/* Animated particles */
.particles-container {
  position: absolute;
  width: 100%;
  height: 100%;
  z-index: 0;
}

.particle {
  position: absolute;
  background: #4f6cff;
  border-radius: 50%;
  opacity: 0.3;
  animation: floatUp linear infinite;
}

@keyframes floatUp {
  from { transform: translateY(100vh) rotate(0deg); }
  to { transform: translateY(-100vh) rotate(360deg); }
}

.game-content {
  width: 100%;
  max-width: 800px;
  padding: 40px 30px;
  background: rgba(20, 21, 47, 0.85);
  backdrop-filter: blur(10px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.5),
              0 0 20px rgba(78, 84, 255, 0.2);
  border-radius: 16px;
  border: 1px solid rgba(78, 84, 255, 0.3);
  position: relative;
  z-index: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}

/* Glowing border effect */
.game-content::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, #4f6cff, #832eff, #4f6cff, #832eff);
  background-size: 400%;
  border-radius: 18px;
  z-index: -1;
  filter: blur(10px);
  opacity: 0.7;
  animation: glowing 20s linear infinite;
}

@keyframes glowing {
  0% { background-position: 0 0; }
  50% { background-position: 400% 0; }
  100% { background-position: 0 0; }
}

.game-header {
  text-align: center;
  margin-bottom: 30px;
  width: 100%;
}

.game-header h1 {
  font-size: 2.5rem;
  font-weight: 800;
  margin: 0;
  background: linear-gradient(to right, #4f6cff, #832eff);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: 2px;
  text-shadow: 0 0 10px rgba(78, 84, 255, 0.3);
}

.battle-title {
  font-size: 1.8rem;
  text-align: center;
  margin: 10px 0 20px;
  color: #ffffff;
  position: relative;
}

.battle-title::after {
  content: '';
  display: block;
  width: 80px;
  height: 3px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  margin: 10px auto 0;
  border-radius: 2px;
}

.countdown-timer {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: bold;
  margin-bottom: 30px;
  padding: 15px 30px;
  background: rgba(20, 21, 47, 0.7);
  border-radius: 12px;
  border: 1px solid rgba(78, 84, 255, 0.4);
  box-shadow: 0 0 15px rgba(78, 84, 255, 0.3);
  color: #fff;
}

.timer-icon {
  margin-right: 10px;
  font-size: 2.2rem;
}

.moves-container {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin: 20px 0;
  width: 100%;
}

.moves-container.disabled {
  opacity: 0.6;
  pointer-events: none;
}

.move-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 180px;
  height: 180px;
  border: none;
  border-radius: 16px;
  background: linear-gradient(135deg, #2d3073, #191b3f);
  border: 2px solid rgba(78, 84, 255, 0.4);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.3),
              0 0 15px rgba(78, 84, 255, 0.2),
              inset 0 0 10px rgba(78, 84, 255, 0.1);
  transition: all 0.3s ease;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.move-btn::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, #4f6cff, #832eff, #4f6cff, #832eff);
  background-size: 400%;
  border-radius: 18px;
  z-index: -1;
  filter: blur(5px);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.move-btn:hover {
  transform: translateY(-10px) scale(1.05);
}

.move-btn:hover::before {
  opacity: 0.5;
  animation: glowing 20s linear infinite;
}

.move-icon {
  font-size: 4rem;
  margin-bottom: 10px;
}

.move-name {
  font-size: 1.6rem;
  font-weight: bold;
  color: #fff;
  text-shadow: 0 0 5px rgba(78, 84, 255, 0.8);
}

.result-message {
  margin-top: 40px;
  font-size: 2.5rem;
  font-weight: bold;
  padding: 15px 40px;
  border-radius: 12px;
  background: rgba(20, 21, 47, 0.8);
  border: 1px solid rgba(78, 84, 255, 0.4);
  box-shadow: 0 0 20px rgba(78, 84, 255, 0.3);
  color: #fff;
  text-shadow: 0 0 5px rgba(78, 84, 255, 0.8);
}

.play-again-btn {
  margin-top: 30px;
  padding: 15px 30px;
  font-size: 1.5rem;
  font-weight: bold;
  background: linear-gradient(135deg, #4ade80, #22d3ee);
  border: none;
  border-radius: 12px;
  color: #fff;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2),
              0 0 15px rgba(34, 211, 238, 0.4);
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.play-again-btn:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 25px rgba(0, 0, 0, 0.3),
              0 0 20px rgba(34, 211, 238, 0.5);
  background: linear-gradient(135deg, #22d3ee, #4ade80);
}

.confetti-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 1000;
}

@media (max-width: 768px) {
  .moves-container {
    flex-direction: column;
    align-items: center;
  }

  .move-btn {
    width: 150px;
    height: 150px;
  }

  .move-icon {
    font-size: 3rem;
  }

  .move-name {
    font-size: 1.3rem;
  }

  .countdown-timer,
  .result-message {
    font-size: 1.8rem;
    padding: 12px 25px;
  }
}
</style>
