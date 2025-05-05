<template>
  <div class="ranking-container">
    <h2 class="ranking-title">🏆 Leaderboard</h2>

    <div class="overflow-x-auto">
      <table class="ranking-table">
        <thead>
          <tr>
            <th class="rank-header">#</th>
            <th class="player-header">Player</th>
            <th class="stats-header">Wins</th>
            <th class="stats-header">Win Rate</th> <!-- TODO FIX THIS get this information from the db on statistics-->
            <th class="stats-header">Games Played</th> <!-- TODO FIX THIS get this information from the db on statistics-->
          </tr>
        </thead>
        <tbody>
          <tr v-for="(player, index) in ranking" :key="player.player_id" class="player-row">
            <td class="rank-cell" :class="rankColor(index)">
              {{ index + 1 }}
            </td>
            <td class="player-cell">
              <div class="player-name-container">
                <span class="player-name">{{ player.username }}</span>
                <span class="player-id">
                  ({{ player.player_id.slice(0, 6) }}…)
                </span>
              </div>
            </td>
            <td class="stats-cell">{{ player.wins }}</td>
            <td class="stats-cell">
              {{ calculateWinRate(player.wins, player.games_played) }}%
            </td>
            <td class="stats-cell">{{ player.games_played }}</td>
          </tr>

          <!-- Empty rows placeholder when no data -->
          <tr v-if="ranking.length === 0" v-for="i in 3" :key="`empty-${i}`" class="empty-row">
            <td colspan="5" class="empty-cell">Loading rankings...</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="move-stats-container">
      <h3 class="move-stats-title">📊 Move Usage</h3>
      <div class="move-stats-grid">
        <div v-for="(percent, move) in moveUsage" :key="move" class="move-stat">
          <div class="move-header">
            <span class="move-name">{{ move }}</span>
            <span class="move-percent">{{ percent.toFixed(2) }}%</span>
          </div>
          <div class="move-bar-bg">
            <div class="move-bar-fill" :style="{
              width: percent + '%',
              background: getMoveColor(move)
            }"></div>
          </div>
        </div>

        <!-- Empty bars placeholder when no data -->
        <div v-if="Object.keys(moveUsage).length === 0" v-for="move in ['rock', 'paper', 'scissors']" :key="`empty-${move}`" class="move-stat">
          <div class="move-header">
            <span class="move-name">{{ move }}</span>
            <span class="move-percent">--%</span>
          </div>
          <div class="move-bar-bg">
            <div class="move-bar-fill pulse-animation" :style="{ background: getMoveColor(move), width: '30%' }"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { getRanking } from '../services/api'

export default defineComponent({
  name: 'Ranking',
  setup() {
    const ranking = ref<
      { player_id: string; username: string; wins: number; games_played: number }[]
    >([])
    const moveUsage = ref<Record<string, number>>({})

    const rankColor = (index: number) => {
      return [
        'rank-gold',
        'rank-silver',
        'rank-bronze'
      ][index] || 'rank-other'
    }

    const getMoveColor = (move: string) => {
      const colors: Record<string, string> = {
        rock: 'linear-gradient(to right, #e44c65, #e41f3f)',
        paper: 'linear-gradient(to right, #4f6cff, #5f7cff)',
        scissors: 'linear-gradient(to right, #832eff, #9340ff)'
      }
      return colors[move] || 'linear-gradient(to right, #666, #888)'
    }

    const calculateWinRate = (wins: number, games: number) => {
      return games > 0 ? ((wins / games) * 100).toFixed(1) : '0.0'
    }

    onMounted(async () => {
      try {
        const res = await getRanking()
        ranking.value = res.data.top_players
        moveUsage.value = res.data.move_usage_percent
      } catch (err) {
        console.error('Failed to fetch ranking:', err)
      }
    })

    return { ranking, moveUsage, rankColor, getMoveColor, calculateWinRate }
  }
})
</script>

<style scoped>
.ranking-container {
  width: 100%;
  background: rgba(20, 21, 47, 0.7);
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(78, 84, 255, 0.2);
  position: relative;
  overflow: hidden;
}

/* Glowing effect for container */
.ranking-container::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: linear-gradient(45deg, #4f6cff80, #832eff80, #4f6cff80, #832eff80);
  background-size: 400%;
  border-radius: 14px;
  z-index: -1;
  filter: blur(8px);
  opacity: 0.5;
  animation: glowing 20s linear infinite;
}

.ranking-title {
  font-size: 1.8rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  text-align: center;
  background: linear-gradient(to right, #4f6cff, #832eff);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  text-shadow: 0 0 10px rgba(131, 46, 255, 0.3);
  letter-spacing: 1px;
}

.ranking-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 0.5rem;
  margin-bottom: 1.5rem;
}

.rank-header, .player-header, .stats-header {
  text-transform: uppercase;
  font-size: 0.75rem;
  color: #6e7cac;
  font-weight: 600;
  padding: 0.5rem 1rem;
  text-align: left;
  letter-spacing: 1px;
}

.player-row {
  background: rgba(30, 32, 70, 0.6);
  transition: all 0.3s ease;
  border-left: 3px solid transparent;
}

.player-row:hover {
  background: rgba(40, 42, 90, 0.8);
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
  border-left: 3px solid #4f6cff;
}

.rank-cell, .player-cell, .stats-cell {
  padding: 0.75rem 1rem;
}

.rank-cell {
  font-weight: 700;
  font-size: 1.1rem;
}

.rank-gold {
  color: #FFD700;
  text-shadow: 0 0 5px rgba(255, 215, 0, 0.5);
}

.rank-silver {
  color: #C0C0C0;
  text-shadow: 0 0 5px rgba(192, 192, 192, 0.5);
}

.rank-bronze {
  color: #CD7F32;
  text-shadow: 0 0 5px rgba(205, 127, 50, 0.5);
}

.rank-other {
  color: #4f6cff;
}

.player-name-container {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.player-name {
  font-weight: 600;
  color: #fff;
}

.player-id {
  font-size: 0.7rem;
  color: #6e7cac;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.player-row:hover .player-id {
  opacity: 1;
}

.stats-cell {
  color: #e0e0ff;
}

.empty-row {
  background: rgba(30, 32, 70, 0.3);
  height: 3rem;
}

.empty-cell {
  text-align: center;
  color: #6e7cac;
  font-style: italic;
}

.move-stats-container {
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid rgba(78, 84, 255, 0.2);
}

.move-stats-title {
  font-size: 1.4rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: #a2a8ff;
}

.move-stats-grid {
  display: grid;
  gap: 1rem;
}

.move-stat {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.move-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.85rem;
  color: #6e7cac;
}

.move-name {
  text-transform: capitalize;
  font-weight: 500;
}

.move-percent {
  font-weight: 700;
}

.move-bar-bg {
  height: 0.5rem;
  background: rgba(30, 32, 70, 0.6);
  border-radius: 0.25rem;
  overflow: hidden;
}

.move-bar-fill {
  height: 100%;
  border-radius: 0.25rem;
  transition: width 1s ease;
}

@keyframes glowing {
  0% { background-position: 0 0; }
  50% { background-position: 400% 0; }
  100% { background-position: 0 0; }
}

@keyframes pulse {
  0% { opacity: 0.5; }
  50% { opacity: 0.8; }
  100% { opacity: 0.5; }
}

.pulse-animation {
  animation: pulse 1.5s infinite;
}

@media (max-width: 640px) {
  .ranking-container {
    padding: 1rem;
  }

  .rank-header, .player-header, .stats-header,
  .rank-cell, .player-cell, .stats-cell {
    padding: 0.5rem;
    font-size: 0.9rem;
  }

  .ranking-title {
    font-size: 1.5rem;
  }
}
</style>
