import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from '../components/RegisterForm.vue'
import Lobby from '../components/Lobby.vue'
import GameRoom from '../components/GameRoom.vue'

const routes = [
  { path: '/', component: LoginForm },
  { path: '/register', component: RegisterForm },
  { path: '/lobby', component: Lobby },
  { path: '/game/:roomId', component: GameRoom },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
