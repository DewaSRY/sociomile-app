// import type { CreateOrganizationRequest, OrganizationResponse } from '~/types';
// import { API_ENDPOINTS } from '~/config/api.config';
// import { createServerApi } from '~/composables/useApi';

// export default defineEventHandler(async (event) => {
//   const body = await readBody<CreateOrganizationRequest>(event);
//   const api = createServerApi(event);

//   try {
//     const response = await api.post<OrganizationResponse>(
//       API_ENDPOINTS.organizations.create,
//       body
//     );
//     return response;
//   } catch (error: any) {
//     throw createError({
//       statusCode: error.response?.status || 500,
//       message: error.response?.data?.message || 'Failed to create organization',
//     });
//   }
// });
