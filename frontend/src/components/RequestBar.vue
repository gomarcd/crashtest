<template>
  <div class="relative flex items-center h-8">
    <div class="relative flex-grow bg-gray-800 border border-gray-700 rounded-md h-8 flex items-center overflow-hidden">
      <div class="border-r border-gray-700">
        <select
          v-model="selectedMethod"
          class="bg-gray-800 text-gray-400 border-0 rounded-none h-full px-2 py-1 text-sm focus:ring-0 appearance-none"
          style="min-width: 80px; --wails-draggable: none;"
          title="Select HTTP Method"
        >
          <option v-for="method in REQUEST_METHODS" :key="method.name" :value="method.name">{{ method.name }}</option>
        </select>
      </div>
      
      <div class="relative flex-grow">
        <input
          ref="urlInputRef"
          v-model="url"
          class="w-full bg-gray-800 text-white border-0 rounded-none h-full px-3 py-1 text-sm focus:ring-0"
          :style="{
            '--wails-draggable': 'none',
            'color': 'transparent', 
            'caretColor': 'white'
          }"
          @focus="storePreviousUrl"
          @keyup.enter="emitSendRequest"
          @keydown.escape.prevent="restorePreviousUrl"
        >
        
        <div
          class="absolute inset-0 pointer-events-none flex items-center text-sm"
          :class="url ? 'justify-start px-3' : 'justify-center pr-[32px]'"
        >
          <template v-if="url">
            <span class="text-gray-500">{{ urlProtocol }}</span>
            <span class="text-white">{{ urlWithoutProtocol }}</span>
          </template>
          <template v-else>
            <span class="text-gray-500 text-lg">Endpoint {{ shortcutText }}</span>
          </template>
        </div>
      </div>
      
      <button
        style="--wails-draggable:none;"
        class="h-full px-3 text-gray-400 hover:text-indigo-400 transition-colors focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed"
        :disabled="!url"
        title="Send Request"
        @click="emitSendRequest"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor"><path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11h2v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" /></svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, nextTick } from 'vue';
import { REQUEST_METHODS } from '../types';
import type { RequestMethod } from '../types';

const props = defineProps<{
  isNonMac: boolean;
}>();

const emit = defineEmits(['send-request']);

const selectedMethod = ref<RequestMethod>('GET');
const url = ref('https://jsonplaceholder.typicode.com/users');
const urlInputRef = ref<HTMLInputElement | null>(null);
const previousUrl = ref('');

const urlProtocol = computed(() => {
  const match = url.value.match(/^(https?:\/\/)/i);
  return match ? match[1] : '';
});

const urlWithoutProtocol = computed(() => {
  return url.value.replace(/^(https?:\/\/)/i, '');
});

const shortcutText = computed(() => {
  return props.isNonMac ? 'Ctrl+L' : '⌘L';
});

function storePreviousUrl() {
  previousUrl.value = url.value;
}

async function restorePreviousUrl() {
  url.value = previousUrl.value;
  await nextTick();
  if (urlInputRef.value) {
    urlInputRef.value.focus();
    urlInputRef.value.select();
  }
}

function emitSendRequest() {
  if (!url.value.trim()) return;
  
  let processedUrl = url.value.trim();
  if (!processedUrl.match(/^https?:\/\//i)) {
    processedUrl = `http://${processedUrl}`;
    url.value = processedUrl;
  }
  
  emit('send-request', {
    method: selectedMethod.value,
    url: processedUrl
  });
}

function focusUrlInput() {
  if (urlInputRef.value) {
    urlInputRef.value.focus();
    urlInputRef.value.select();
  }
}

function handleUrlBarShortcut(e: KeyboardEvent) {
  const isMac = !props.isNonMac;
  const metaOrCtrl = isMac ? e.metaKey : e.ctrlKey;

  if (metaOrCtrl && e.key.toLowerCase() === 'l') {
    e.preventDefault();
    focusUrlInput();
  }
}

function updateUrl(newUrl: string) {
  url.value = newUrl;
  previousUrl.value = newUrl;
}

onMounted(() => {
  window.addEventListener('keydown', handleUrlBarShortcut);
});

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleUrlBarShortcut);
});

defineExpose({
  updateUrl,
  focusUrlInput,
  getUrl: () => url.value,
  getMethod: () => selectedMethod.value
});
</script>