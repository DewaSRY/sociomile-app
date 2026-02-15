import type { CreateConversationRequest, ConversationResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const body = await readBody<CreateConversationRequest>(event);
  const api = createServerApi(event);

  try {
    const response = await api.post<ConversationResponse>(
      API_ENDPOINTS.conversations.create,
      body
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to create conversation',
    });
  }
});
