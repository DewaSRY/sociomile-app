import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type {
  CommonResponse,
  CreateConversationRequest,
} from "$shared/types";
import type { AxiosResponse } from "axios";
import { API_GUEST_CONVERSATION } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const body = await readBody<CreateConversationRequest>(event);
  const token = getCookie(event, "auth_token");
  try {
    const { data } = await apiClient.post<
      CreateConversationRequest,
      AxiosResponse<CommonResponse>
    >(API_GUEST_CONVERSATION, body, {
      headers: {
        Authorization: token,
      },
    });

    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
