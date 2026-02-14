import type { ConversationMessageListResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const conversationId = getRouterParam(event, 'id');
  const api = createServerApi(event);
  const query = getQuery(event);

  try {
    const response = await api.get<ConversationMessageListResponse>(
      API_ENDPOINTS.messages.list(Number(conversationId)),
      { params: query }
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to fetch messages',
    });
  }
});
