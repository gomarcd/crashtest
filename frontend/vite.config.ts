import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': '/src',
    },
  },
  // Define options to optimize Monaco Editor bundling
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          // Create a separate chunk for Monaco Editor
          'monaco-editor': ['monaco-editor'],
        },
      },
    },
  }
})