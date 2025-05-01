import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL,
})

const statisticsApi = axios.create({
  baseURL: import.meta.env.VITE_STATISTICS_URL,
})

let token: string | null = null
console.log("VITE_BACKEND_URL:", import.meta.env.VITE_BACKEND_URL)
console.log("VITE_STATISTICS_URL:", import.meta.env.VITE_STATISTICS_URL)

// Check if there's a token in localStorage and set it in axios headers
const savedToken = localStorage.getItem('authToken')
if (savedToken) {
  setToken(savedToken)
}


export function setToken(t: string) {
  console.log("token",t)
  token = t
  // localStorage.setItem('authToken', t)
  api.defaults.headers.common.Authorization = `${t}`
}

export function register(username: string, password: string) {
  return api.post('/auth/register', { username, password })
}

export function login(username: string, password: string) {
  return api.post<{ token: string, playerId: string }>('/auth/login', { username, password })
}

export function createRoom() {
  return api.post<{ room_id: string }>('/room')
}

// TODO (on backend api)
export function getRooms() {
  return api.get<{ room_id: string }[]>('/rooms')
}


export function getRanking() {
  return statisticsApi.get<{ username: string; wins: number }[]>('/statistics/ranking')
}
