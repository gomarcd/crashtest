import { createApp } from 'vue';
import App from './App.vue';
import './style.css';

createApp(App as import('vue').DefineComponent).mount('#app');