<template>
  <router-view v-if="!user" @logged-in="onLoggedIn" @registered="onRegistered"/>

  <div v-else-if="!roomId" class="game-wrapper">
    <GameLobby @room-created="onRoomCreated" @room-joined="onRoomJoined" />
  </div>

  <div v-else-if="!inGame">
    <WaitingRoom :room-id="roomId" :user-id="user.id" :token="user.token" @game-start="onGameStart"/>
  </div>

  <div v-else>
    <Game :web-socket="webSocket" :room-id="roomId" :user-id="user.id"/>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import GameLobby from './components/Lobby.vue'
import WaitingRoom from './components/WaitingRoom.vue'
import Game from './components/Game.vue'

export default {
  components: { GameLobby, WaitingRoom, Game },
  setup() {
    const user = ref<{ token: string, id: string } | null>(null)
    const roomId = ref<string | null>(null)
    const webSocket = ref<WebSocket | null>(null)
    const inGame = ref(false)
    const router = useRouter()

    // Check for existing session on component mount
    onMounted(() => {
      const token = localStorage.getItem('authToken')
      const playerId = localStorage.getItem('playerId')

      if (token && playerId) {
        user.value = { token, id: playerId }
      }
    })

    function onLoggedIn(data: { token: string, userId: string }) {
      user.value = { token: data.token, id: data.userId }
      localStorage.setItem('authToken', data.token)
      localStorage.setItem('playerId', data.userId)
    }

    function onRegistered(data: { token: string, playerId: string }) {
      // Store token and redirect to login or lobby
      localStorage.setItem('authToken', data.token)
      localStorage.setItem('playerId', data.playerId)
      user.value = { token: data.token, id: data.playerId }
    }

    function onRoomCreated(id: string) {
      roomId.value = id
    }

    function onRoomJoined(id: string) {
      roomId.value = id
    }

    function onGameStart(payload: { ws: WebSocket }) {
      webSocket.value = payload.ws
      inGame.value = true
    }

    return {
      user,
      roomId,
      webSocket,
      inGame,
      onLoggedIn,
      onRegistered,
      onRoomCreated,
      onRoomJoined,
      onGameStart
    }
  }
}
</script>

<style>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

body {
  margin: 0;
  padding: 0;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background: #0d0e24;
}

.game-wrapper {
  width: 100%;
  min-height: 100vh;
}
</style>
