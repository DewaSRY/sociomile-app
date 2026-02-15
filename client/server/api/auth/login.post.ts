import type { LoginRequest, AuthResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const body = await readBody<LoginRequest>(event);
  const api = createServerApi(event);

  try {
    const response = await api.post<AuthResponse>(API_ENDPOINTS.auth.login, body);
    
    // Set cookie
    setCookie(event, 'auth_token', response.token, {
      maxAge: 60 * 60 * 24 * 7, // 7 days
      sameSite: 'lax',
      secure: process.env.NODE_ENV === 'production',
      httpOnly: true,
      path: '/',
    });

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Login failed',
    });
  }
});
