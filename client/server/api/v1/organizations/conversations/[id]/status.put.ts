import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { CommonResponse, UpdateConversationRequest } from "$shared/types";
import type { AxiosResponse } from "axios";
import { ORG_CONVERSATION } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const body = await readBody<UpdateConversationRequest>(event);
  const token = getCookie(event, "auth_token");
  try {
    const { data } = await apiClient.put<
      UpdateConversationRequest,
      AxiosResponse<CommonResponse>
    >(`${ORG_CONVERSATION}/${id}/status`, body, {
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
