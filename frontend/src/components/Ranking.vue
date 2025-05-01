<template>
  <div class="w-full max-w-md bg-gray-800 rounded-lg shadow-lg p-6">
    <h2 class="text-2xl font-bold mb-4 text-center">🏆 Top Players</h2>
    <ul class="divide-y divide-gray-700">
      <li v-for="(player, index) in ranking" :key="index" class="py-2 flex justify-between">
        <span>{{ index + 1 }}. {{ player.username }}</span>
        <span class="font-semibold">{{ player.wins }} wins</span>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { getRanking } from '../services/api'

export default defineComponent({
  name: 'Ranking',
  setup() {
    const ranking = ref<{ username: string; wins: number }[]>([])

    onMounted(async () => {
      try {
        const res = await getRanking()
        ranking.value = res.data
      } catch (err) {
        console.error('Failed to fetch ranking:', err)
      }
    })

    return { ranking }
  }
})
</script>

<style scoped>
</style>
