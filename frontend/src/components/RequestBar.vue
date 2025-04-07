<template>
  <div class="flex gap-2 w-full">
    <div class="flex-shrink-0">
      <select 
        v-model="selectedMethod" 
        class="h-10 rounded-md px-3 font-medium border-gray-300 focus:border-primary focus:ring-primary"
      >
        <option 
          v-for="method in REQUEST_METHODS" 
          :key="method.name" 
          :value="method.name"
          :class="method.color"
        >
          {{ method.name }}
        </option>
      </select>
    </div>
    
    <input 
      v-model="url" 
      placeholder="Enter request URL" 
      class="flex-grow h-10 rounded-md px-3 border-gray-300 focus:border-primary focus:ring-primary"
    >
    
    <button 
      class="bg-primary hover:bg-primary-700 text-white h-10 px-4 rounded-md font-medium transition-colors" 
      :disabled="!url"
      @click="sendRequest"
    >
      Send
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits } from 'vue';
import { REQUEST_METHODS, RequestConfig } from '../types';

const emit = defineEmits(['send-request']);

const selectedMethod = ref('GET');
const url = ref('');

function sendRequest() {
  if (!url.value) return;
  
  const config: RequestConfig = {
    method: selectedMethod.value,
    url: url.value,
    headers: {},
    queryParams: {},
    body: ''
  };
  
  emit('send-request', config);
}
</script>