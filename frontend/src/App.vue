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
          <!-- Method Selector -->
          <div class="border-r border-gray-700">
            <select 
              v-model="selectedMethod" 
              class="bg-gray-800 text-white border-0 rounded-none h-full px-2 py-1 text-sm focus:ring-0 appearance-none"
              style="min-width: 80px;"
            >
              <option
                v-for="method in REQUEST_METHODS"
                :key="method.name"
                :value="method.name"
              >
                {{ method.name }}
              </option>
            </select>
          </div>
          
          <input 
            v-model="url" 
            placeholder="Enter request URL" 
            class="flex-grow bg-gray-800 text-white border-0 rounded-none h-full px-3 py-1 text-sm focus:ring-0"
            style="--wails-draggable:none;"
          >
          
          <button 
            style="--wails-draggable:none;"
            class="h-full px-3 text-gray-400 hover:text-indigo-400 transition-colors focus:outline-none" 
            :disabled="!url"
            title="Send Request"
            @click="sendRequest"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11h2v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
            </svg>
          </button>
        </div>
      </div>
      
      <div class="flex-grow grid grid-cols-1 md:grid-cols-2 gap-4 overflow-hidden">
        <!-- Request Panel -->
        <div class="border border-gray-700 rounded-md overflow-hidden shadow-sm bg-gray-800">
          <div class="px-4 py-2 font-medium border-b border-gray-700">
            Request
          </div>
          <div class="h-[calc(100%-40px)]">
            <div class="flex border-b border-gray-700">
              <button 
                v-for="tab in ['Headers', 'Params', 'Body']" 
                :key="tab"
                class="px-4 py-2 text-sm font-medium transition-colors"
                :class="activeTab === tab.toLowerCase() ? 'border-b-2 border-indigo-500 text-indigo-400' : 'text-gray-400 hover:text-white'"
                @click="activeTab = tab.toLowerCase()"
              >
                {{ tab }}
              </button>
            </div>
            
            <div
              v-if="activeTab === 'headers'"
              class="p-4"
            >
              <div
                v-for="(header, index) in headersList"
                :key="index"
                class="flex items-center gap-2 mb-2 w-full"
              >
                <input 
                  v-model="header.enabled" 
                  type="checkbox" 
                  class="rounded text-indigo-500 bg-gray-700 border-gray-600 focus:ring-indigo-500"
                >
                <input 
                  v-model="header.key" 
                  type="text" 
                  placeholder="Header" 
                  class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500"
                >
                <input 
                  v-model="header.value" 
                  type="text" 
                  placeholder="Value" 
                  class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500"
                >
                <button 
                  class="text-red-500 hover:text-red-400" 
                  @click="removeHeader(index)"
                >
                  ✕
                </button>
              </div>
              <button 
                class="mt-2 text-indigo-400 hover:text-indigo-300 text-sm" 
                @click="addHeader"
              >
                + Add Header
              </button>
            </div>
            
            <div
              v-if="activeTab === 'params'"
              class="p-4"
            >
              <div
                v-for="(param, index) in params"
                :key="index"
                class="flex items-center gap-2 mb-2 w-full"
              >
                <input 
                  v-model="param.enabled" 
                  type="checkbox" 
                  class="rounded text-indigo-500 bg-gray-700 border-gray-600 focus:ring-indigo-500"
                >
                <input 
                  v-model="param.key" 
                  type="text" 
                  placeholder="Parameter" 
                  class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500"
                >
                <input 
                  v-model="param.value" 
                  type="text" 
                  placeholder="Value" 
                  class="flex-1 min-w-0 rounded-md bg-gray-700 border-gray-600 text-gray-200 focus:border-indigo-500 focus:ring-indigo-500"
                >
                <button 
                  class="text-red-500 hover:text-red-400" 
                  @click="removeParam(index)"
                >
                  ✕
                </button>
              </div>
              <button 
                class="mt-2 text-indigo-400 hover:text-indigo-300 text-sm" 
                @click="addParam"
              >
                + Add Parameter
              </button>
            </div>
            
            <div
              v-if="activeTab === 'body'"
              class="h-full"
            >
              <textarea 
                v-model="requestBody" 
                placeholder="Request body (JSON, XML, etc.)" 
                class="w-full h-full p-4 border-none bg-gray-800 text-gray-200 focus:ring-0 font-mono text-sm"
              />
            </div>
          </div>
        </div>
        
        <div class="border border-gray-700 rounded-md overflow-hidden shadow-sm bg-gray-800">
          <div class="flex items-center px-4 py-2 font-medium border-b border-gray-700">
            <span>Response</span>
            <div
              v-if="response"
              class="ml-4 flex items-center gap-2"
            >
              <div 
                class="px-2 py-1 rounded-md text-white font-medium text-xs"
                :class="statusColorClass"
              >
                {{ response.statusCode }}
              </div>
              <div class="text-gray-400 text-xs">
                {{ response.timeMs }}ms
              </div>
            </div>
          </div>
          
          <div class="h-[calc(100%-40px)]">
            <div
              v-if="response"
              class="flex border-b border-gray-700"
            >
              <button 
                v-for="tab in ['Body', 'Headers']" 
                :key="tab"
                class="px-4 py-2 text-sm font-medium transition-colors"
                :class="activeResponseTab === tab.toLowerCase() ? 'border-b-2 border-indigo-500 text-indigo-400' : 'text-gray-400 hover:text-white'"
                @click="activeResponseTab = tab.toLowerCase()"
              >
                {{ tab }}
              </button>
            </div>
            
            <div
              v-if="activeResponseTab === 'body' && response"
              class="h-[calc(100%-35px)] overflow-auto"
              style="--wails-draggable:none;"
            >
              <div v-if="formattedResponse" class="relative">
                <button
                  class="absolute top-2 right-2 text-gray-400 hover:text-indigo-400 transition-colors"
                  @click="copyToClipboard"
                  title="Copy to Clipboard"
                >
                  <svg
                    v-if="!copied"
                    xmlns="http://www.w3.org/2000/svg"
                    height="24"
                    viewBox="0 0 24 24"
                    width="24"
                    fill="currentColor"
                  >
                    <path d="M0 0h24v24H0z" fill="none"/>
                    <path d="M16 1H4c-1.1 0-2 .9-2 2v14h2V3h12V1zm3 4H8c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h11c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm0 16H8V7h11v14z"/>
                  </svg>
                  <svg
                    v-if="copied"
                    xmlns="http://www.w3.org/2000/svg"
                    height="24"
                    viewBox="0 0 24 24"
                    width="24"
                    fill="#34c759"
                  >
                    <path d="M0 0h24v24H0z" fill="none"/>
                    <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/>
                  </svg>
                </button>
                <pre
                  class="p-4 whitespace-pre-wrap text-gray-200 font-mono text-sm"
                >{{ formattedResponse }}</pre>
              </div>
              <div
                v-else
                class="p-4 text-gray-500"
              >
                No response body
              </div>
            </div>
            
            <div
              v-if="activeResponseTab === 'headers' && response"
              class="p-4 overflow-auto h-[calc(100%-35px)]"
            >
              <table class="w-full text-sm">
                <thead>
                  <tr class="text-left">
                    <th class="pb-2 font-medium text-gray-400">
                      Header
                    </th>
                    <th class="pb-2 font-medium text-gray-400">
                      Value
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(value, key) in response.headers"
                    :key="key"
                    class="border-t border-gray-700"
                  >
                    <td class="py-2 pr-4 font-medium text-gray-300">
                      {{ key }}
                    </td>
                    <td class="py-2 text-gray-400">
                      {{ value }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <div
              v-if="!response"
              class="flex items-center justify-center h-full text-gray-500"
            >
              Send a request to see the response here
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { Header, QueryParam, RequestConfig, APIResponse, REQUEST_METHODS } from './types';
import { Environment } from '../wailsjs/runtime/runtime';

declare global {
  interface Window {
    go: {
      main: {
        App: {
          Greet(name: string): Promise<string>;
        };
        APIService: {
          SendRequest(config: RequestConfig): Promise<APIResponse>;
        };
      };
    };
  }
}

const isNonMac = ref(false);
onMounted(async () => {
  const env = await Environment();
  isNonMac.value = env.platform.toLowerCase() !== 'darwin';
});

const selectedMethod = ref('GET');
const url = ref('https://jsonplaceholder.typicode.com/users');
const activeTab = ref('body');
const activeResponseTab = ref('body');
const headersList = ref<Header[]>([{ key: '', value: '', enabled: true }]);
const params = ref<QueryParam[]>([{ key: '', value: '', enabled: true }]);
const requestBody = ref('');
const response = ref<APIResponse | null>(null);
const copied = ref(false);

const headers = computed<Record<string, string>>(() => {
  return headersList.value
    .filter(h => h.enabled && h.key.trim() !== '')
    .reduce((acc, h) => {
      acc[h.key.trim()] = h.value;
      return acc;
    }, {} as Record<string, string>);
});

const queryParams = computed<Record<string, string>>(() => {
  return params.value
    .filter(p => p.enabled && p.key.trim() !== '')
    .reduce((acc, p) => {
      acc[p.key.trim()] = p.value;
      return acc;
    }, {} as Record<string, string>);
});

const statusColorClass = computed(() => {
  if (!response.value) return '';
  
  const status = response.value.statusCode;
  if (status >= 200 && status < 300) return 'bg-green-600';
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
  if (headersList.value.length === 0) {
    addHeader();
  }
}

function addParam() {
  params.value.push({ key: '', value: '', enabled: true });
}

function removeParam(index: number) {
  params.value.splice(index, 1);
  if (params.value.length === 0) {
    addParam();
  }
}

async function copyToClipboard() {
  if (formattedResponse.value) {
    await navigator.clipboard.writeText(formattedResponse.value);
    copied.value = true;
    setTimeout(() => {
      copied.value = false;
    }, 2000);
  }
}

async function sendRequest() {
  try {
    const config: RequestConfig = {
      method: selectedMethod.value,
      url: url.value,
      headers: headers.value,
      queryParams: queryParams.value,
      body: requestBody.value,
    };
    
    response.value = await window.go.main.APIService.SendRequest(config);
    activeResponseTab.value = 'body';
  } catch (error) {
    console.error('Error sending request:', error);
    response.value = {
      statusCode: 0,
      headers: {},
      body: `Error: ${error instanceof Error ? error.message : String(error)}`,
      timeMs: 0,
      error: error instanceof Error ? error.message : String(error),
    };
  }
}
</script>