import type { UserData } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const api = createServerApi(event);

  try {
    const user = await api.get<UserData>(API_ENDPOINTS.auth.me);
    return user;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 401,
      message: error.response?.data?.message || 'Unauthorized',
    });
  }
});
