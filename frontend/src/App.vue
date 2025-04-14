<template>
  <div
    class="h-full flex flex-col bg-gray-900 text-gray-200"
    style="--wails-draggable:drag"
  >
    <main
      class="flex-grow h-full flex flex-col px-2 gap-2"
      :class="isNonMac ? 'pt-3 pb-2' : 'pt-8 pb-2'"
    >
      <div class="flex items-center h-8">
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
              placeholder="Enter request URL"
              class="w-full bg-gray-800 text-white border-0 rounded-none h-full px-3 py-1 text-sm focus:ring-0"
              :style="{
                '--wails-draggable': 'none',
                'color': 'transparent', 
                'caretColor': 'white'
              }"
              @focus="storePreviousUrl"
              @keyup.enter="sendRequest"
              @keydown.escape.prevent="restorePreviousUrl"
            >
            
            <div class="absolute inset-0 pointer-events-none flex items-center px-3 text-sm">
              <template v-if="url">
                <span class="text-gray-500">{{ urlProtocol }}</span>
                <span class="text-white">{{ urlWithoutProtocol }}</span>
              </template>
              <span v-else class="text-gray-500">Enter request URL</span>
            </div>
          </div>
          
          <button
            style="--wails-draggable:none;"
            class="h-full px-3 text-gray-400 hover:text-indigo-400 transition-colors focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="!url"
            title="Send Request"
            @click="sendRequest"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor"><path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11h2v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" /></svg>
          </button>
        </div>
      </div>

      <div class="flex-grow grid grid-cols-1 md:grid-cols-2 gap-2">
        <div class="border border-gray-700 rounded-md overflow-hidden shadow-sm bg-gray-800 flex flex-col">
          <div class="h-[calc(100%-40px)]">
            <div class="flex border-b border-gray-700 min-h-[41px]">
              <button
                v-for="tab in ['Body', 'Headers', 'Params']"
                :key="tab"
                class="px-4 pt-2 text-sm font-medium transition-none border-b-2"
                :class="activeTab === tab.toLowerCase()
                  ? 'pb-[5px] border-indigo-500 text-indigo-400'
                  : 'pb-[5px] border-transparent text-gray-400 hover:text-white'"
                @click="activeTab = tab.toLowerCase()"
              >
                {{ tab }}
              </button>
            </div>
            <div class="h-full overflow-y-auto">
              <div v-if="activeTab === 'headers'" class="p-4">
                <div v-for="(header, index) in headersList" :key="index" class="flex items-center gap-2 mb-2 w-full">
                  <input v-model="header.enabled" type="checkbox" class="rounded text-indigo-500 bg-gray-700 border-gray-600 focus:ring-indigo-500">
                  <input v-model="header.key" type="text" placeholder="Header" class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500">
                  <input v-model="header.value" type="text" placeholder="Value" class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500">
                  <button class="text-red-500 hover:text-red-400" @click="removeHeader(index)">✕</button>
                </div>
                <button class="mt-2 text-indigo-400 hover:text-indigo-300 text-sm" @click="addHeader">+ Add Header</button>
              </div>
              <div v-if="activeTab === 'params'" class="p-4">
                <div v-for="(param, index) in params" :key="index" class="flex items-center gap-2 mb-2 w-full">
                  <input v-model="param.enabled" type="checkbox" class="rounded text-indigo-500 bg-gray-700 border-gray-600 focus:ring-indigo-500">
                  <input v-model="param.key" type="text" placeholder="Parameter" class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500">
                  <input v-model="param.value" type="text" placeholder="Value" class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500">
                  <button class="text-red-500 hover:text-red-400" @click="removeParam(index)">✕</button>
                </div>
                <button class="mt-2 text-indigo-400 hover:text-indigo-300 text-sm" @click="addParam">+ Add Parameter</button>
              </div>
              <div v-if="activeTab === 'body'" class="h-full">
                <textarea v-model="requestBody" placeholder="Request body (GraphQL, JSON, etc)" class="w-full h-full p-4 border-none bg-gray-800 text-gray-200 focus:ring-0 font-mono text-sm resize-none"></textarea>
              </div>
            </div>
          </div>
        </div>

        <div
          class="border border-gray-700 rounded-md overflow-hidden shadow-sm bg-gray-800 flex flex-col transition-shadow duration-300 ease-in-out"
          :class="{ 'response-glow': isResponseGlowing }"
        >
          <div class="h-[calc(100%-40px)]">
            <div :class="['flex justify-between border-b border-gray-700 px-4 min-h-[41px]', !response ? 'items-center' : '']">
              <div class="text-gray-500" v-if="!response">Response</div>
              <div class="flex" v-if="response">
                <button
                  v-for="tab in ['Body', 'Headers']"
                  :key="tab"
                  class="px-4 pt-2 text-sm font-medium transition-none border-b-2"
                  :class="activeResponseTab === tab.toLowerCase()
                    ? 'pb-[5px] border-indigo-500 text-indigo-400'
                    : 'pb-[5px] border-transparent text-gray-400 hover:text-white'"
                  @click="activeResponseTab = tab.toLowerCase()"
                >
                  {{ tab }}
                </button>

              </div>
              <div v-else class="flex-1"></div>
              <div v-if="response" class="flex items-center gap-4">
                <div class="text-gray-400 pt-[2px] text-xs">{{ response.timeMs }}ms</div>
                <div class="px-2 py-1 mt-[2px] rounded-md" :class="statusColorClass">{{ response.statusCode }}</div>
              </div>
              <div v-else class="h-[32px]"></div>
            </div>
            <div class="h-[calc(100%-41px)] overflow-y-auto">
              <div v-if="activeResponseTab === 'body' && response" style="--wails-draggable:none;">
                <div v-if="formattedResponse" class="relative p-4">
                  <button class="absolute top-2 right-2 text-gray-400 hover:text-indigo-400 transition-colors" @click="copyToClipboard" title="Copy to Clipboard">
                    <svg v-if="!copied" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="h-4 w-4 size-6"><path stroke-linecap="round" stroke-linejoin="round" d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75" /></svg>
                    <svg v-if="copied" class="h-4 w-4 size-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#34c759"><path d="M0 0h24v24H0z" fill="none"/><path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/></svg>
                  </button>
                  <pre class="whitespace-pre-wrap text-gray-200 font-mono text-sm">{{ formattedResponse }}</pre>
                </div>
                <div v-else class="p-4 text-gray-500">No response body</div>
              </div>
              <div v-if="activeResponseTab === 'headers' && response" class="p-4">
                <table class="w-full text-sm">
                  <thead><tr class="text-left"><th class="pb-2 font-medium text-gray-400">Header</th><th class="pb-2 font-medium text-gray-400">Value</th></tr></thead>
                  <tbody><tr v-for="(value, key) in response.headers" :key="key" class="border-t border-gray-700"><td class="py-2 pr-4 font-medium text-gray-300 break-words">{{ key }}</td><td class="py-2 text-gray-400 break-words">{{ value }}</td></tr></tbody>
                </table>
              </div>
              <div v-if="!response" class="flex items-center justify-center h-full text-gray-500">Send a request to see the response here</div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, nextTick } from 'vue';
import { Header, QueryParam, RequestConfig, APIResponse, REQUEST_METHODS, RequestMethod } from './types';
import { Environment } from '../wailsjs/runtime/runtime';

declare global {
  interface Window {
    go: {
      main: {
        APIService: {
          SendRequest(config: RequestConfig): Promise<APIResponse>;
        };
      };
    };
    runtime?: { 
      Environment(): Promise<{ platform: string; [key: string]: unknown }>;
    }
  }
}

const isNonMac = ref(false);
const selectedMethod = ref<RequestMethod>('GET');
const url = ref('https://jsonplaceholder.typicode.com/users');
const activeTab = ref('body');
const activeResponseTab = ref('body');
const headersList = ref<Header[]>([{ key: '', value: '', enabled: true }]);
const params = ref<QueryParam[]>([{ key: '', value: '', enabled: true }]);
const requestBody = ref('');
const response = ref<APIResponse | null>(null);
const copied = ref(false);
const urlInputRef = ref<HTMLInputElement | null>(null);
const previousUrl = ref('');
const isResponseGlowing = ref(false);

const urlProtocol = computed(() => {
  const match = url.value.match(/^(https?:\/\/)/i);
  return match ? match[1] : '';
});

const urlWithoutProtocol = computed(() => {
  return url.value.replace(/^(https?:\/\/)/i, '');
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

const headers = computed<Record<string, string>>(() => {
  return headersList.value
    .filter(h => h.enabled && h.key.trim())
    .reduce((acc, h) => { acc[h.key.trim()] = h.value; return acc; }, {});
});

const queryParams = computed<Record<string, string>>(() => {
  return params.value
    .filter(p => p.enabled && p.key.trim())
    .reduce((acc, p) => { acc[p.key.trim()] = p.value; return acc; }, {});
});

const statusColorClass = computed(() => {
  if (!response.value) return '';
  const status = response.value.statusCode;
  if (status >= 200 && status < 300) return 'bg-green-500 text-green-100 font-bold text-xs';
  if (status >= 300 && status < 400) return 'bg-blue-600';
  if (status >= 400 && status < 500) return 'bg-yellow-600';
  if (status >= 500) return 'bg-red-600';
  return 'bg-gray-600';
});

const formattedResponse = computed(() => {
  if (!response.value || response.value.body === null) return '';
  const body = response.value.body;
  if (typeof body === 'string') {
    return body;
  } else if (typeof body === 'object') {
    return JSON.stringify(body, null, 2);
  }
  return '';
});

function addHeader() {
  headersList.value.push({ key: '', value: '', enabled: true });
}

function removeHeader(index: number) {
  headersList.value.splice(index, 1);
  if (headersList.value.length === 0) addHeader();
}

function addParam() {
  params.value.push({ key: '', value: '', enabled: true });
}

function removeParam(index: number) {
  params.value.splice(index, 1);
  if (params.value.length === 0) addParam();
}

async function copyToClipboard() {
  if (formattedResponse.value) {
    await navigator.clipboard.writeText(formattedResponse.value);
    copied.value = true;
    setTimeout(() => { copied.value = false; }, 2000);
  }
}

async function sendRequest() {
  isResponseGlowing.value = false;
  if (!url.value.trim()) return;

  const currentUrlInBar = url.value.trim();
  let processedUrl = currentUrlInBar;

  if (!processedUrl.match(/^https?:\/\//i)) {
    processedUrl = "http://" + processedUrl;
  }

  try {
    const config: RequestConfig = {
      method: selectedMethod.value,
      url: processedUrl,
      headers: headers.value,
      queryParams: queryParams.value,
      body: requestBody.value || null,
    };

    response.value = await window.go.main.APIService.SendRequest(config);
    activeResponseTab.value = 'body';

    if (response.value && response.value.usedURL) {
      url.value = response.value.usedURL;
      previousUrl.value = url.value;
    }
  } catch (error: unknown) {
    console.error('Error sending request:', error);
    const errorMessage = error instanceof Error 
      ? error.message 
      : typeof error === 'string' 
        ? error 
        : 'Unknown error';
        
    response.value = {
      statusCode: 0,
      headers: {},
      body: `Error: ${errorMessage}`,
      timeMs: 0,
      error: errorMessage,
    };
    activeResponseTab.value = 'body';
  } finally {
    isResponseGlowing.value = true;
    setTimeout(() => {
      isResponseGlowing.value = false;
    }, 2000);
  }
}

function handleUrlBarShortcut(e: KeyboardEvent) {
  const isMac = /Mac|iPod|iPhone|iPad/.test(navigator.platform);
  const metaOrCtrl = isMac ? e.metaKey : e.ctrlKey;

  if (metaOrCtrl && e.key.toLowerCase() === 'l') {
    e.preventDefault();
    if (urlInputRef.value) {
      urlInputRef.value.focus();
      urlInputRef.value.select();
    }
  }
}

interface EnvironmentInfo {
  platform: string;
  [key: string]: unknown;
}

onMounted(async () => {
  try {
    let envInfo: EnvironmentInfo;
    if (window.runtime) {
      const env = await window.runtime.Environment();
      envInfo = env as EnvironmentInfo;
    } else {
      const env = await Environment();
      envInfo = env as EnvironmentInfo;
    }
    isNonMac.value = envInfo.platform.toLowerCase() !== 'darwin';
  } catch (e) {
    console.warn("Could not determine platform:", e);
  }

  window.addEventListener('keydown', handleUrlBarShortcut);
});

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleUrlBarShortcut);
});
</script>


<style>
.response-glow {
  box-shadow: 0 0 15px 5px rgba(99, 102, 241, 0.55);
}

::-webkit-scrollbar { width: 8px; height: 8px; }
::-webkit-scrollbar-track { background: #1f2937; border-radius: 4px; }
::-webkit-scrollbar-thumb { background-color: #4b5563; border-radius: 4px; border: 2px solid #1f2937; }
::-webkit-scrollbar-thumb:hover { background-color: #6b7280; }
::-webkit-scrollbar-corner { background: #1f2937; }
* { scrollbar-width: thin; scrollbar-color: #4b5563 #1f2937; }
textarea { resize: none; }
*:focus-visible { outline: 2px solid #818cf8; outline-offset: 1px; }
*:focus:not(:focus-visible) { outline: none; }
select {
  background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%239ca3af' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
  background-position: right 0.5rem center; background-repeat: no-repeat; background-size: 1.5em 1.5em; padding-right: 2.5rem;
}
</style>