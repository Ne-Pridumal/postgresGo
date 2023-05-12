import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import * as path from 'path'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@features": `${path.resolve(__dirname, "./src/features/")}`,
      "@shared": `${path.resolve(__dirname, "./src/shared/")}`,
      "@entities": `${path.resolve(__dirname, "./src/entities/")}`,
      "@widgets": `${path.resolve(__dirname, "./src/widgets/")}`,
      "@app": `${path.resolve(__dirname, "./src/app/")}`,
      "@processes": `${path.resolve(__dirname, "./src/processes/")}`,
      "@pages": `${path.resolve(__dirname, "./src/pages/")}`
    },
  },

})
