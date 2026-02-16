import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { ConversationResponse, Filters } from "$shared/types";
import type { AxiosResponse } from "axios";
import { API_GUEST_CONVERSATION } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const token = getCookie(event, "auth_token");
  const query = getQuery(event) as Partial<Filters>;
  try {
    const { data } = await apiClient.get<AxiosResponse<ConversationResponse>>(
      API_GUEST_CONVERSATION,
      {
        headers: {
          Authorization: token,
        },
        params: query,
      },
    );

    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
