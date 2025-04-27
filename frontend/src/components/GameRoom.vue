<template>
  <div>
    <h2>Game Room: {{ roomId }}</h2>
    <div class="flex space-x-2">
      <button @click="sendMove('rock')" class="bg-gray-300 p-2">🪨 Rock</button>
      <button @click="sendMove('paper')" class="bg-gray-300 p-2">📄 Paper</button>
      <button @click="sendMove('scissors')" class="bg-gray-300 p-2">✂️ Scissors</button>
    </div>
    <div class="mt-4">
      <strong>Messages:</strong>
      <ul class="list-disc list-inside">
        <li v-for="msg in messages" :key="msg">{{ msg }}</li>
      </ul>
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
      socket = new WebSocket(`ws://localhost:8080/join/${props.roomId}`)
      socket.onopen = () => {
        messages.value.push('Connected to room')
        // send join message
        console.log("props.playerId",props)
        socket.send(JSON.stringify({
          type: 'join',
          body: JSON.stringify({ game_id: props.roomId }),
          sender: props.playerId
        }))
      }
      socket.onmessage = (ev) => {
        messages.value.push(`> ${ev.data}`)
      }
      socket.onerror = (e) => {
        messages.value.push('WebSocket error')
        console.error(e)
      }
      socket.onclose = () => {
        messages.value.push('WebSocket closed')
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
    }

    return { messages, sendMove }
  },
}
</script>
