// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from "pathe";
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },

  modules: [
    "@nuxt/a11y",
    "@nuxt/eslint",
    "@nuxt/hints",
    "@nuxt/image",
    "@nuxt/ui",
    "@artmizu/nuxt-prometheus",
    "@formkit/nuxt",
  ],

  runtimeConfig: {
    public: {
      apiBaseUrl:
        process.env.NUXT_PUBLIC_API_BASE_URL || "http://localhost:8080/api",
    },
  },

  alias: {
    $types: resolve(__dirname, "types"),
    $components: resolve(__dirname, "app/components"),
    $utils: resolve(__dirname, "utils"),
    $libs: resolve(__dirname, "libs"),
    $assets: resolve(__dirname, "assets"),
    $constants: resolve(__dirname, "constants"),
    $shared: resolve(__dirname, "shared"),
  },

  // Add color mode configuration
  colorMode: {
    preference: "light",
  },
});