// API configuration
export const API_CONFIG = {
  baseURL: "http://localhost:8080/api", //process.env.NUXT_PUBLIC_API_BASE_URL ||
  timeout: 30000,
};

export const API_ENDPOINTS = {
  // Auth
  auth: {
    login: '/auth/login',
    register: '/auth/register',
    refresh: '/auth/refresh',
    me: '/auth/me',
  },
  // Organizations
  organizations: {
    list: '/organizations',
    create: '/organizations',
    get: (id: number) => `/organizations/${id}`,
    update: (id: number) => `/organizations/${id}`,
    delete: (id: number) => `/organizations/${id}`,
  },
  // Conversations
  conversations: {
    list: '/conversations',
    create: '/conversations',
    get: (id: number) => `/conversations/${id}`,
    update: (id: number) => `/conversations/${id}`,
    assign: (id: number) => `/conversations/${id}/assign`,
  },
  // Messages
  messages: {
    list: (conversationId: number) => `/conversations/${conversationId}/messages`,
    create: '/messages',
  },
  // Tickets
  tickets: {
    list: '/tickets',
    create: '/tickets',
    get: (id: number) => `/tickets/${id}`,
    update: (id: number) => `/tickets/${id}`,
  },
};
