<template>
  <div v-if="!user">
    <LoginForm @logged-in="onLoggedIn"/>
    <RegisterForm @registered="onRegistered"/>
  </div>

  <div v-else-if="!roomId" class="min-h-screen flex flex-col items-center justify-center bg-gray-900 text-white space-y-10 p-8 animate__animated animate__fadeIn">
    <h1 class="text-4xl font-bold mb-6">🎮 Game Lobby</h1>
    <RoomCreator @room-created="onRoomCreated"/>
    <div class="border-t border-gray-700 w-full my-6"></div>
    <RoomJoin @room-joined="onRoomJoined" />
    <div class="border-t border-gray-700 w-full my-6"></div>
    <Ranking />
  </div>

  <div v-else-if="!inGame">
    <WaitingRoom :room-id="roomId" :user-id="user.id" :token="user.token" @game-start="onGameStart"/>
  </div>

  <div v-else>
    <Game :web-socket="webSocket" :room-id="roomId" :user-id="user.id"/>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import LoginForm from './components/LoginForm.vue'
import RegisterForm from './components/RegisterForm.vue'
import RoomCreator from './components/RoomCreator.vue'
import RoomJoin from './components/RoomJoin.vue'
import WaitingRoom from './components/WaitingRoom.vue'
import Game from './components/Game.vue'
import Ranking from './components/Ranking.vue'


export default {
  components: { LoginForm, RegisterForm, RoomCreator, RoomJoin, WaitingRoom, Game, Ranking },
  setup() {
    const user = ref<{ token: string, id: string } | null>(null)
    const roomId = ref<string | null>(null)
    const webSocket = ref<WebSocket | null>(null)
    const inGame = ref(false)

    function onLoggedIn(data: { token: string, userId: string }) {
      user.value = { token: data.token, id: data.userId }
    }

    function onRegistered() {
      // Maybe just show a toast "Registered successfully!" and go to login
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

    return { user, roomId, webSocket, inGame, onLoggedIn, onRegistered, onRoomCreated, onRoomJoined, onGameStart }
  }
}
</script>
