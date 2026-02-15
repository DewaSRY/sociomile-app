import type { AssignConversationRequest, ConversationResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id');
  const body = await readBody<AssignConversationRequest>(event);
  const api = createServerApi(event);

  try {
    const response = await api.post<ConversationResponse>(
      API_ENDPOINTS.conversations.assign(Number(id)),
      body
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to assign conversation',
    });
  }
});
