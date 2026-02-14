import type { ConversationResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id');
  const api = createServerApi(event);

  try {
    const response = await api.get<ConversationResponse>(
      API_ENDPOINTS.conversations.get(Number(id))
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to fetch conversation',
    });
  }
});
