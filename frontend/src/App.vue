<template>
  <div v-if="!user">
    <LoginForm @logged-in="onLoggedIn"/>
    <RegisterForm @registered="onRegistered"/>
  </div>

  <div v-else-if="!roomId">
    <RoomCreator @room-created="onRoomCreated"/>
    <RoomJoin @room-joined="onRoomJoined" />
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

export default {
  components: { LoginForm, RegisterForm, RoomCreator, RoomJoin, WaitingRoom, Game },
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
