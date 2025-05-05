<template>
  <div class="game-lobby-container">
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

    <div class="game-lobby-content animate__animated animate__fadeIn">
      <div class="game-logo">
        <h1>Jokenpo Realm</h1>
      </div>

      <h2 class="lobby-title">📜 Lobby 🪨</h2>

      <!-- Create Room Section -->
      <div class="lobby-section">
        <RoomCreator @room-created="emitRoomCreated" />
      </div>

      <!-- Divider with glow effect -->
      <div class="glowing-divider"></div>

      <!-- Join Room Section -->
      <div class="lobby-section">
        <RoomJoin @room-joined="emitRoomJoined" />
      </div>

      <!-- Divider with glow effect -->
      <div class="glowing-divider"></div>

      <!-- 🏆 Ranking Section -->
      <div class="lobby-section">
        <Ranking />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import RoomCreator from './RoomCreator.vue'
import RoomJoin from './RoomJoin.vue'
import Ranking from './Ranking.vue'
import { defineComponent } from 'vue'

export default defineComponent({
  components: { RoomCreator, RoomJoin, Ranking },
  emits: ['room-created', 'room-joined'],
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

.game-lobby-container {
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

.game-lobby-content {
  width: 100%;
  max-width: 700px;
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
.game-lobby-content::before {
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

.lobby-title {
  font-size: 1.8rem;
  text-align: center;
  margin-bottom: 30px;
  color: #ffffff;
  position: relative;
}

.lobby-title::after {
  content: '';
  display: block;
  width: 80px;
  height: 3px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  margin: 10px auto 0;
  border-radius: 2px;
}

.lobby-section {
  margin-bottom: 15px;
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

@media (max-width: 768px) {
  .game-lobby-content {
    padding: 30px 20px;
  }
}

/* General styles for form elements within the lobby */
:deep(.lobby-section-title) { /* Use :deep() to target child components */
  font-size: 1.5rem; /* Slightly smaller than main lobby title */
  font-weight: 700;
  text-align: center;
  margin-bottom: 1rem;
  color: #e0e0ff; /* Lighter color */
  text-shadow: 0 0 8px rgba(78, 84, 255, 0.3);
}

:deep(.lobby-input) {
  width: 100%;
  padding: 0.8rem 1rem;
  border-radius: 8px;
  background-color: rgba(30, 32, 70, 0.7); /* Match ranking rows */
  border: 1px solid rgba(78, 84, 255, 0.4);
  color: #ffffff;
  font-size: 1rem;
  transition: all 0.3s ease;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.3);
}

:deep(.lobby-input::placeholder) {
  color: #6e7cac;
  opacity: 0.8;
}

:deep(.lobby-input:focus) {
  outline: none;
  border-color: #832eff;
  background-color: rgba(40, 42, 90, 0.8);
  box-shadow: 0 0 15px rgba(131, 46, 255, 0.4),
              inset 0 1px 3px rgba(0, 0, 0, 0.3);
}

:deep(.lobby-button) {
  width: 100%;
  padding: 0.8rem 1rem;
  border-radius: 8px;
  border: none;
  background: linear-gradient(to right, #4f6cff, #5f7cff); /* Blue gradient */
  color: white;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 10px rgba(79, 108, 255, 0.3);
  text-transform: uppercase;
  letter-spacing: 1px;
}

:deep(.lobby-button:hover) {
  background: linear-gradient(to right, #5f7cff, #6f8cff);
  box-shadow: 0 6px 15px rgba(95, 124, 255, 0.4);
  transform: translateY(-2px);
}

:deep(.lobby-button:active) {
  transform: translateY(0);
  box-shadow: 0 2px 5px rgba(79, 108, 255, 0.3);
}

/* Specific button variations if needed */
:deep(.lobby-button.create) {
  background: linear-gradient(to right, #832eff, #9340ff); /* Purple gradient */
  box-shadow: 0 4px 10px rgba(131, 46, 255, 0.3);
}
:deep(.lobby-button.create:hover) {
  background: linear-gradient(to right, #9340ff, #a350ff);
   box-shadow: 0 6px 15px rgba(147, 64, 255, 0.4);
}
:deep(.lobby-button.create:active) {
  box-shadow: 0 2px 5px rgba(131, 46, 255, 0.3);
}

:deep(.room-id-display) {
  margin-top: 1rem;
  padding: 0.8rem 1rem;
  background-color: rgba(30, 32, 70, 0.7);
  color: #e0e0ff;
  border-radius: 8px;
  text-align: center;
  border: 1px dashed rgba(78, 84, 255, 0.4);
  font-size: 0.9rem;
}

:deep(.room-id-display span) { /* Style the actual ID differently */
  font-family: 'Courier New', Courier, monospace;
  font-weight: bold;
  color: #7efca2; /* A vibrant green */
  margin-left: 0.5rem;
  word-break: break-all;
}

/* Remove default component backgrounds if they should blend */
:deep(.room-creator-container),
:deep(.room-join-container) {
  background-color: transparent !important; /* Override gray background */
  box-shadow: none !important; /* Override default shadow */
  padding: 0 !important; /* Reset padding if lobby-section handles it */
  max-width: 100% !important; /* Allow it to fill lobby-section */
  margin: 0 !important;
}
</style>
