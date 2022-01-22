import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import AllAudits from '../views/Audit.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/audits/all',
    name: 'All Audits',
    component: AllAudits,
    props: { query: "all" }
  },
  {
    path: '/audits/player',
    name: 'Player Audits',
    component: AllAudits,
    props: { query: "player" }
  },
  {
    path: '/audits/chat',
    name: 'Chat Audits',
    component: AllAudits,
    props: { query: "chat" }
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
