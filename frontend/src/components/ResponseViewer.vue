<template>
  <div class="h-full flex flex-col">
    <div
      v-if="response"
      class="p-4 bg-gray-100 border-b flex flex-col gap-2"
    >
      <div class="flex items-center gap-3">
        <div 
          class="px-2 py-1 rounded-md text-white font-medium text-sm"
          :class="statusColorClass"
        >
          {{ response.statusCode }}
        </div>
          
        <div class="text-gray-600 text-sm">
          {{ response.timeMs }}ms
        </div>
      </div>
    </div>
      
    <div class="tabs flex bg-gray-100 border-b">
      <button 
        v-for="tab in tabs" 
        :key="tab.id"
        class="px-4 py-2 text-sm font-medium transition-colors"
        :class="[
          activeTab === tab.id 
            ? 'border-b-2 border-primary text-primary' 
            : 'text-gray-600 hover:text-gray-900'
        ]"
        @click="activeTab = tab.id"
      >
        {{ tab.name }}
      </button>
    </div>
      
    <div
      v-if="response"
      class="flex-grow overflow-auto"
    >
      <!-- Body Tab -->
      <div
        v-show="activeTab === 'body'"
        class="h-full font-mono text-sm"
      >
        <pre
          v-if="formattedResponse"
          class="p-4 whitespace-pre-wrap"
        >{{ formattedResponse }}</pre>
        <div
          v-else
          class="p-4 text-gray-500"
        >
          No response body
        </div>
      </div>
        
      <!-- Headers Tab -->
      <div
        v-show="activeTab === 'headers'"
        class="p-4"
      >
        <table class="w-full text-sm">
          <thead>
            <tr class="text-left">
              <th class="pb-2 font-semibold">
                Header
              </th>
              <th class="pb-2 font-semibold">
                Value
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(value, key) in response.headers"
              :key="key"
              class="border-t"
            >
              <td class="py-2 pr-4 font-medium">
                {{ key }}
              </td>
              <td class="py-2 text-gray-700">
                {{ value }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
      
    <div
      v-else
      class="flex-grow flex items-center justify-center text-gray-500"
    >
      Send a request to see the response here
    </div>
  </div>
</template>
  
  <script setup lang="ts">
  import { ref, computed, watch } from 'vue';
  import { APIResponse } from '../types';
  
  const props = defineProps<{
    response: APIResponse | null;
  }>();
  
  const activeTab = ref('body');
  const tabs = [
    { id: 'body', name: 'Body' },
    { id: 'headers', name: 'Headers' }
  ];
  
  const statusColorClass = computed(() => {
    if (!props.response) return '';
    
    const status = props.response.statusCode;
    if (status >= 200 && status < 300) return 'bg-green-500';
    if (status >= 300 && status < 400) return 'bg-blue-500';
    if (status >= 400 && status < 500) return 'bg-yellow-500';
    if (status >= 500) return 'bg-red-500';
    return 'bg-gray-500';
  });
  
  const formattedResponse = computed(() => {
    if (!props.response || !props.response.body) return '';
    
    if (typeof props.response.body === 'object') {
      return JSON.stringify(props.response.body, null, 2);
    }
    
    return props.response.body.toString();
  });
  
  // Reset to body tab when receiving a new response
  watch(() => props.response, () => {
    activeTab.value = 'body';
  });
  </script>