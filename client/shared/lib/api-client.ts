import axios from "axios";

export const apiClient = axios.create({
  baseURL: process.env.NUXT_PUBLIC_API_BASE ?? "http://localhost:8080",
  headers: {
  },
  withCredentials: true,
});


