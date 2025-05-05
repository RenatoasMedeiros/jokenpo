<template>
  <div class="waiting-room-container">
    <!-- Animated background particles -->
    <div class="particles-container">
      <div v-for="n in 10" :key="n" class="particle"
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

    <div class="waiting-content animate__animated animate__fadeIn">
      <div class="game-logo">
        <h1>GAME REALM</h1>
      </div>

      <h2 class="waiting-title">Waiting for opponent...</h2>

      <div class="loader-container">
        <div class="loader-ring">
          <div class="loader-ring-inner"></div>
        </div>
      </div>

      <div class="room-info">
        <div class="room-id-label">Room ID</div>
        <div class="room-id-value">{{ roomId }}</div>

        <div class="copy-button" @click="copyRoomId" title="Copy Room ID">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
            <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
          </svg>
          <span v-if="copied" class="copied-text">Copied!</span>
        </div>
      </div>

      <div class="waiting-message">
        Share this Room ID with a friend to start playing!
      </div>

      <div class="pulsar-ring"></div>
    </div>
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
    userId: {
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
    const copied = ref(false)

    const copyRoomId = () => {
      navigator.clipboard.writeText(props.roomId)
      copied.value = true
      setTimeout(() => {
        copied.value = false
      }, 2000)
    }

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

        if (message.type === 'players_ready') {
          console.log('Two players are in the room, sending start signal...')

          const startMessage = {
            type: "start",
            body: "",
            sender: props.userId,
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

    return { copied, copyRoomId }
  }
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

.waiting-room-container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: radial-gradient(circle at center, #1a1b4b 0%, #0d0e24 70%);
  position: relative;
  overflow: hidden;
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

.waiting-content {
  width: 100%;
  max-width: 500px;
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
  text-align: center;
}

/* Glowing border effect */
.waiting-content::before {
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

.game-logo {
  text-align: center;
  margin-bottom: 20px;
}

.game-logo h1 {
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

.waiting-title {
  font-size: 1.8rem;
  font-weight: 700;
  margin-bottom: 30px;
  color: #ffffff;
  position: relative;
}

.waiting-title::after {
  content: '';
  display: block;
  width: 60px;
  height: 3px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  margin: 10px auto 0;
  border-radius: 2px;
}

.loader-container {
  position: relative;
  width: 120px;
  height: 120px;
  margin: 10px 0 30px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.loader-ring {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  position: relative;
  animation: rotate 2s linear infinite;
}

.loader-ring::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 4px solid transparent;
  border-top: 4px solid #4f6cff;
  border-right: 4px solid #832eff;
  filter: drop-shadow(0 0 8px rgba(78, 84, 255, 0.7));
}

.loader-ring-inner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: radial-gradient(circle at center, rgba(78, 84, 255, 0.2) 0%, rgba(131, 46, 255, 0.1) 70%);
  animation: pulse 1.5s ease-in-out infinite alternate;
}

@keyframes rotate {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@keyframes pulse {
  0% { transform: translate(-50%, -50%) scale(0.8); opacity: 0.5; }
  100% { transform: translate(-50%, -50%) scale(1.1); opacity: 0.8; }
}

.room-info {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
  padding: 15px 25px;
  background: rgba(20, 21, 47, 0.7);
  border-radius: 8px;
  border: 1px solid rgba(78, 84, 255, 0.3);
}

.room-id-label {
  font-size: 0.8rem;
  color: #6e7cac;
  margin-bottom: 5px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.room-id-value {
  font-family: 'Courier New', monospace;
  font-weight: 600;
  font-size: 1.5rem;
  color: #4f6cff;
  text-shadow: 0 0 5px rgba(78, 84, 255, 0.5);
  letter-spacing: 1px;
}

.copy-button {
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 6px;
  color: #6e7cac;
  background: rgba(20, 21, 47, 0.7);
  cursor: pointer;
  transition: all 0.2s ease;
}

.copy-button:hover {
  background: rgba(78, 84, 255, 0.2);
  color: #4f6cff;
}

.copied-text {
  position: absolute;
  top: -25px;
  left: 50%;
  transform: translateX(-50%);
  background: rgba(20, 21, 47, 0.9);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.7rem;
  color: #4f6cff;
  white-space: nowrap;
  animation: fadeInOut 2s ease;
}

@keyframes fadeInOut {
  0% { opacity: 0; }
  15% { opacity: 1; }
  85% { opacity: 1; }
  100% { opacity: 0; }
}

.waiting-message {
  color: #a2a8ff;
  font-size: 0.9rem;
  max-width: 280px;
  margin: 5px auto 0;
}

.pulsar-ring {
  position: absolute;
  width: 150%;
  height: 150%;
  border-radius: 50%;
  border: 2px solid rgba(78, 84, 255, 0.1);
  animation: pulsar 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
  z-index: -1;
}

@keyframes pulsar {
  0% { transform: scale(0.5); opacity: 0; }
  50% { transform: scale(1); opacity: 0.3; }
  100% { transform: scale(1.5); opacity: 0; }
}

@media (max-width: 576px) {
  .waiting-content {
    padding: 30px 20px;
    margin: 0 20px;
  }

  .game-logo h1 {
    font-size: 2rem;
  }

  .waiting-title {
    font-size: 1.5rem;
  }

  .loader-container {
    width: 100px;
    height: 100px;
  }

  .loader-ring {
    width: 70px;
    height: 70px;
  }

  .loader-ring-inner {
    width: 50px;
    height: 50px;
  }
}
</style>
