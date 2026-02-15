import type { CreateConversationMessageRequest, ConversationMessageResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const body = await readBody<CreateConversationMessageRequest>(event);
  const api = createServerApi(event);

  try {
    const response = await api.post<ConversationMessageResponse>(
      API_ENDPOINTS.messages.create,
      body
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to send message',
    });
  }
});
