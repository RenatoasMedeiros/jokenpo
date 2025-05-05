<template>
  <div class="game-room-container">
    <!-- Animated background particles -->
    <div class="particles-container">
      <div v-for="n in 15" :key="n" class="particle"
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

    <div class="game-room-content animate__animated animate__fadeIn">
      <div class="game-logo">
        <h1>Jokenpo Realm</h1>
      </div>

      <div class="room-header">
        <h2 class="room-title">⚡ Game Room ⚡</h2>
        <div class="room-id">
          <span>Room Code:</span>
          <div class="room-code">{{ roomId }}</div>
        </div>
      </div>

      <div class="moves-section">
        <h3 class="section-title">Choose Your Move</h3>
        <div class="moves-container">
          <button @click="sendMove('rock')" class="move-btn animate__animated animate__fadeInUp" style="animation-delay: 0.1s">
            <div class="move-icon">🪨</div>
            <div class="move-name">Rock</div>
          </button>
          <button @click="sendMove('paper')" class="move-btn animate__animated animate__fadeInUp" style="animation-delay: 0.3s">
            <div class="move-icon">📄</div>
            <div class="move-name">Paper</div>
          </button>
          <button @click="sendMove('scissors')" class="move-btn animate__animated animate__fadeInUp" style="animation-delay: 0.5s">
            <div class="move-icon">✂️</div>
            <div class="move-name">Scissors</div>
          </button>
        </div>
      </div>

      <div class="glowing-divider"></div>

      <div class="messages-section">
        <h3 class="section-title">Game Status</h3>
        <div class="messages-container">
          <div v-for="(msg, index) in messages" :key="index" class="message-item animate__animated animate__fadeIn">
            {{ msg }}
          </div>
          <div v-if="messages.length === 0" class="no-messages">
            Waiting for game to start...
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, onMounted, onBeforeUnmount, PropType } from 'vue'

interface WSMessage {
  type: string
  body: string
  sender: string
}

export default {
  props: {
    roomId: { type: String as PropType<string>, required: true },
    playerId: { type: String as PropType<string>, required: true },
  },
  setup(props) {
    const messages = ref<string[]>([])
    let socket: WebSocket

    onMounted(() => {
      let BACKEND_URL = import.meta.env.VITE_BACKEND_URL

      BACKEND_URL = BACKEND_URL.replace("http://", "")
      console.log("BACKEND_URL:", BACKEND_URL)
      socket = new WebSocket(`wss://${BACKEND_URL}/join/${props.roomId}`);

      socket.onopen = () => {
        messages.value.push('🎮 Connected to game room')
        // send join message
        console.log("props.playerId", props)
        socket.send(JSON.stringify({
          type: 'join',
          body: JSON.stringify({ game_id: props.roomId }),
          sender: props.playerId
        }))
      }

      socket.onmessage = (ev) => {
        // Format and clean up the message
        let msgData = ev.data
        try {
          const parsed = JSON.parse(msgData)
          if (parsed.type === 'join') {
            msgData = '👤 Player joined the room'
          } else if (parsed.type === 'countdown') {
            msgData = `⏱️ Countdown: ${parsed.body}`
          } else if (parsed.type === 'move') {
            msgData = '🎯 Move received'
          } else if (parsed.type === 'gameover') {
            if (parsed.body === 'draw') {
              msgData = '🤝 Game over: It\'s a draw!'
            } else {
              msgData = `🏆 Game over: ${parsed.body}`
            }
          }
        } catch (e) {
          // If it's not valid JSON, use the raw message
        }

        messages.value.push(msgData)

        // Limit to last 10 messages for better UI
        if (messages.value.length > 10) {
          messages.value = messages.value.slice(-10)
        }
      }

      socket.onerror = (e) => {
        messages.value.push('❌ WebSocket error')
        console.error(e)
      }

      socket.onclose = () => {
        messages.value.push('🔌 WebSocket closed')
      }
    })

    onBeforeUnmount(() => {
      socket.close()
    })

    function sendMove(move: string) {
      const msg: WSMessage = {
        type: 'move',
        body: JSON.stringify({ game_id: props.roomId, move }),
        sender: props.playerId
      }
      socket.send(JSON.stringify(msg))
      messages.value.push(`✅ You chose: ${move}`)
    }

    return { messages, sendMove }
  },
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

.game-room-container {
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

.game-room-content {
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
}

/* Glowing border effect */
.game-room-content::before {
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

.room-header {
  text-align: center;
  margin-bottom: 30px;
}

.room-title {
  font-size: 1.8rem;
  color: #ffffff;
  position: relative;
  margin-bottom: 15px;
}

.room-title::after {
  content: '';
  display: block;
  width: 80px;
  height: 3px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  margin: 10px auto 0;
  border-radius: 2px;
}

.room-id {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 15px;
}

.room-id span {
  font-size: 1rem;
  color: #a0a0cf;
  margin-bottom: 5px;
}

.room-code {
  font-size: 1.4rem;
  font-weight: bold;
  background: rgba(79, 108, 255, 0.2);
  padding: 8px 20px;
  border-radius: 8px;
  border: 1px solid rgba(79, 108, 255, 0.4);
  color: #fff;
  letter-spacing: 2px;
}

.section-title {
  font-size: 1.4rem;
  color: #fff;
  margin-bottom: 20px;
  text-align: center;
  position: relative;
  display: inline-block;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 100%;
  height: 2px;
  background: linear-gradient(to right, rgba(79, 108, 255, 0.1), rgba(79, 108, 255, 0.8), rgba(79, 108, 255, 0.1));
}

.moves-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 30px;
}

.moves-container {
  display: flex;
  justify-content: center;
  gap: 20px;
  flex-wrap: wrap;
}

.move-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 140px;
  height: 140px;
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
  transform: translateY(-5px) scale(1.05);
}

.move-btn:hover::before {
  opacity: 0.5;
  animation: glowing 20s linear infinite;
}

.move-icon {
  font-size: 3rem;
  margin-bottom: 10px;
}

.move-name {
  font-size: 1.2rem;
  font-weight: bold;
  color: #fff;
  text-shadow: 0 0 5px rgba(78, 84, 255, 0.8);
}

.glowing-divider {
  height: 1px;
  width: 100%;
  background: linear-gradient(to right, rgba(78, 84, 255, 0.1), rgba(131, 46, 255, 0.6), rgba(78, 84, 255, 0.1));
  margin: 30px 0;
  position: relative;
}

.glowing-divider::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(to right, rgba(78, 84, 255, 0), rgba(131, 46, 255, 0.8), rgba(78, 84, 255, 0));
  filter: blur(3px);
}

.messages-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 20px;
}

.messages-container {
  width: 100%;
  max-height: 250px;
  overflow-y: auto;
  background: rgba(20, 21, 47, 0.5);
  border-radius: 12px;
  padding: 15px;
  border: 1px solid rgba(78, 84, 255, 0.3);
  box-shadow: inset 0 0 15px rgba(0, 0, 0, 0.3);
  scrollbar-width: thin;
  scrollbar-color: rgba(78, 84, 255, 0.5) rgba(20, 21, 47, 0.3);
}

.messages-container::-webkit-scrollbar {
  width: 8px;
}

.messages-container::-webkit-scrollbar-track {
  background: rgba(20, 21, 47, 0.3);
  border-radius: 4px;
}

.messages-container::-webkit-scrollbar-thumb {
  background: rgba(78, 84, 255, 0.5);
  border-radius: 4px;
}

.message-item {
  padding: 10px 12px;
  margin-bottom: 8px;
  border-radius: 8px;
  background: rgba(79, 108, 255, 0.15);
  border-left: 3px solid rgba(79, 108, 255, 0.6);
  color: #d0d0f0;
  font-size: 0.95rem;
  text-shadow: 0 1px 1px rgba(0, 0, 0, 0.2);
  animation-duration: 0.5s;
}

.message-item:last-child {
  margin-bottom: 0;
}

.no-messages {
  padding: 20px;
  text-align: center;
  color: #a0a0cf;
  font-style: italic;
}

@media (max-width: 768px) {
  .game-room-content {
    padding: 30px 20px;
  }

  .moves-container {
    flex-direction: column;
  }

  .move-btn {
    width: 120px;
    height: 120px;
    margin-bottom: 15px;
  }

  .move-icon {
    font-size: 2.5rem;
  }

  .move-name {
    font-size: 1rem;
  }

  .messages-container {
    max-height: 200px;
  }
}
</style>
