import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import Hello from "./Views/Hello.vue";
import {createRouter, createWebHistory} from "vue-router";
import Setting from "./Views/Setting.vue";
import Chat from "./Views/Chat.vue";
const routes = [
    { path: '/', component:  Hello},
    {path: '/setting',component: Setting},
    {path: '/chat',component: Chat}
]
const router = createRouter({
    history: createWebHistory(),
    routes
})
const app=createApp(App)
app.use(router)
app.mount('#app')