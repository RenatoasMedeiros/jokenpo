<template>
  <div class="game-login-container">
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

    <form @submit.prevent="submit" class="game-login-form animate__animated animate__fadeIn">
      <div class="game-logo">
        <h1>Jokenpo Realm</h1>
      </div>

      <h2 class="login-title">Welcome Back!</h2>

      <div class="form-control-group">
        <div class="form-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
            <circle cx="12" cy="7" r="4"></circle>
          </svg>
        </div>
        <input v-model="username" placeholder="Username" required class="game-input" />
      </div>

      <div class="form-control-group">
        <div class="form-icon">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
        </div>
        <input v-model="password" type="password" placeholder="Password" required class="game-input" />
      </div>

      <button type="submit" class="game-button">
        <span>LET'S PLAY</span>
        <div class="button-glow"></div>
      </button>

      <div class="register-link">
        <span>New player?</span>
        <router-link to="/register">Create account</router-link>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login, setToken } from '../services/api'
import { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

export default {
  emits: ['logged-in'],
  setup(_, { emit }) {
    const username = ref('')
    const password = ref('')
    const router = useRouter()

    async function submit() {
      try {
        const res = await login(username.value, password.value)

        setToken(res.data.token)
        localStorage.setItem('playerId', res.data.playerId)

        emit('logged-in', {
          token: res.data.token,
          userId: res.data.playerId
        })

        toast.success('Logged in successfully!', {
          position: "top-center",
          autoClose: 2000,
        })

        setTimeout(() => {
          router.push('/lobby')
        }, 2000)

      } catch (error: any) {
        toast.error(error.response?.data?.error || 'Login failed. Please check your credentials.', {
          position: "top-center",
        })
      }
    }

    return { username, password, submit }
  },
}
</script>

<style scoped>
@import 'https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css';

.game-login-container {
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

.game-login-form {
  width: 100%;
  max-width: 420px;
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
.game-login-form::before {
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

.login-title {
  font-size: 1.4rem;
  text-align: center;
  margin-bottom: 30px;
  color: #ffffff;
  position: relative;
}

.login-title::after {
  content: '';
  display: block;
  width: 60px;
  height: 3px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  margin: 10px auto 0;
  border-radius: 2px;
}

.form-control-group {
  position: relative;
  margin-bottom: 20px;
}

.form-icon {
  position: absolute;
  left: 15px;
  top: 50%;
  transform: translateY(-50%);
  color: #6e7cac;
  z-index: 2;
}

.game-input {
  width: 100%;
  padding: 15px 15px 15px 45px;
  background: rgba(16, 17, 35, 0.7);
  border: 1px solid rgba(78, 84, 255, 0.3);
  border-radius: 8px;
  color: #ffffff;
  font-size: 16px;
  transition: all 0.3s ease;
}

.game-input:focus {
  outline: none;
  border-color: #4f6cff;
  box-shadow: 0 0 0 3px rgba(78, 84, 255, 0.25);
}

.game-input::placeholder {
  color: #6e7cac;
}

.game-button {
  width: 100%;
  padding: 15px;
  margin-top: 10px;
  background: linear-gradient(to right, #4f6cff, #832eff);
  border: none;
  border-radius: 8px;
  color: white;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 1px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.game-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 7px 15px rgba(78, 84, 255, 0.4);
}

.game-button:active {
  transform: translateY(-1px);
}

.button-glow {
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.3), transparent);
  transition: 0.5s;
}

.game-button:hover .button-glow {
  left: 100%;
}

.register-link {
  margin-top: 25px;
  text-align: center;
  color: #6e7cac;
  font-size: 14px;
}

.register-link a {
  color: #4f6cff;
  text-decoration: none;
  font-weight: 600;
  margin-left: 8px;
  transition: all 0.2s ease;
}

.register-link a:hover {
  color: #832eff;
  text-decoration: underline;
}

@media (max-width: 576px) {
  .game-login-form {
    padding: 30px 20px;
    margin: 0 20px;
  }
}
</style>
