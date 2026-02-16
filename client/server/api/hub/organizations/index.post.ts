import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { CommonResponse, RegisterOrganizationRequest } from "$shared/types";
import type { AxiosResponse } from "axios";
import { HUB_ORGANIZATION } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const body = await readBody<RegisterOrganizationRequest>(event);
  const token = getCookie(event, "auth_token");
  try {
    const { data } = await apiClient.post<
      RegisterOrganizationRequest,
      AxiosResponse<CommonResponse>
    >(HUB_ORGANIZATION, body, {
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
