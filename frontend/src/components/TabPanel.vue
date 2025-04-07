<template>
  <div class="w-full h-full flex flex-col">
    <div class="flex bg-gray-100 border-b">
      <button 
        v-for="tab in tabs" 
        :key="tab.id"
        class="px-4 py-2 text-sm font-medium transition-colors"
        :class="[
          activeTab === tab.id 
            ? 'border-b-2 border-primary text-primary' 
            : 'text-gray-600 hover:text-gray-900'
        ]"
        @click="selectTab(tab.id)"
      >
        {{ tab.name }}
      </button>
    </div>
      
    <div class="flex-grow overflow-auto">
      <div
        v-show="activeTab === 'headers'"
        class="p-4"
      >
        <div
          v-for="(header, index) in headers"
          :key="index"
          class="flex items-center gap-2 mb-2 w-full"
        >
          <input 
            v-model="header.enabled" 
            type="checkbox" 
            class="rounded text-primary focus:ring-primary flex-shrink-0"
          >
          <input 
            v-model="header.key" 
            type="text" 
            placeholder="Header" 
            class="flex-1 min-w-0 rounded-md border-gray-300 focus:border-primary focus:ring-primary"
          >
          <input 
            v-model="header.value" 
            type="text" 
            placeholder="Value" 
            class="flex-1 min-w-0 rounded-md border-gray-300 focus:border-primary focus:ring-primary"
          >
          <button 
            class="text-red-500 hover:text-red-700 flex-shrink-0" 
            @click="removeHeader(index)"
          >
            ✕
          </button>
        </div>
        <button 
          class="mt-2 text-primary hover:text-primary-700 text-sm" 
          @click="addHeader"
        >
          + Add Header
        </button>
      </div>
        
      <div
        v-show="activeTab === 'params'"
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
            class="rounded text-primary focus:ring-primary flex-shrink-0"
          >
          <input 
            v-model="param.key" 
            type="text" 
            placeholder="Parameter" 
            class="flex-1 min-w-0 rounded-md border-gray-300 focus:border-primary focus:ring-primary"
          >
          <input 
            v-model="param.value" 
            type="text" 
            placeholder="Value" 
            class="flex-1 min-w-0 rounded-md border-gray-300 focus:border-primary focus:ring-primary"
          >
          <button 
            class="text-red-500 hover:text-red-700 flex-shrink-0" 
            @click="removeParam(index)"
          >
            ✕
          </button>
        </div>
        <button 
          class="mt-2 text-primary hover:text-primary-700 text-sm" 
          @click="addParam"
        >
          + Add Parameter
        </button>
      </div>
        
      <div
        v-show="activeTab === 'body'"
        class="h-full"
      >
        <textarea 
          v-model="requestBody" 
          placeholder="Request body (JSON, XML, etc.)" 
          class="w-full h-full p-4 border-none focus:ring-0 font-mono text-sm"
        />
      </div>
    </div>
  </div>
</template>
  
  <script setup lang="ts">
  import { ref, defineEmits, watch } from 'vue';
  import { Header, QueryParam } from '../types';
  
  const emit = defineEmits(['update-headers', 'update-params', 'update-body']);
  
  const tabs = [
    { id: 'headers', name: 'Headers' },
    { id: 'params', name: 'Params' },
    { id: 'body', name: 'Body' }
  ];
  
  const activeTab = ref('headers');
  const headers = ref<Header[]>([{ key: '', value: '', enabled: true }]);
  const params = ref<QueryParam[]>([{ key: '', value: '', enabled: true }]);
  const requestBody = ref('');
  
  function selectTab(tabId: string) {
    activeTab.value = tabId;
  }
  
  function addHeader() {
    headers.value.push({ key: '', value: '', enabled: true });
  }
  
  function removeHeader(index: number) {
    headers.value.splice(index, 1);
    if (headers.value.length === 0) {
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
  
  watch(headers, (newHeaders) => {
    const enabledHeaders = newHeaders
      .filter(h => h.enabled && h.key.trim() !== '')
      .reduce((acc, h) => {
        acc[h.key.trim()] = h.value;
        return acc;
      }, {} as Record<string, string>);
    
    emit('update-headers', enabledHeaders);
  }, { deep: true });
  
  watch(params, (newParams) => {
    const enabledParams = newParams
      .filter(p => p.enabled && p.key.trim() !== '')
      .reduce((acc, p) => {
        acc[p.key.trim()] = p.value;
        return acc;
      }, {} as Record<string, string>);
    
    emit('update-params', enabledParams);
  }, { deep: true });
  
  watch(requestBody, (newBody) => {
    emit('update-body', newBody);
  });
  </script>