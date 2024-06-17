import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      "/graph": {
        target: "http://192.168.1.5:8080",
        changeOrigin: true,
      },
    },
  },
  plugins: [react()],
});
