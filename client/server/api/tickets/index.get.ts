import type { TicketListResponse } from '~/types';
import { API_ENDPOINTS } from '~/config/api.config';
import { createServerApi } from '~/composables/useApi';

export default defineEventHandler(async (event) => {
  const api = createServerApi(event);
  const query = getQuery(event);

  try {
    const response = await api.get<TicketListResponse>(
      API_ENDPOINTS.tickets.list,
      { params: query }
    );
    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.response?.status || 500,
      message: error.response?.data?.message || 'Failed to fetch tickets',
    });
  }
});
