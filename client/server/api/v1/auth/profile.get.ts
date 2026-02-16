import { defineEventHandler, setCookie, getCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { UserProfileData } from "$shared/types";
import type { AxiosError } from "axios";

import { API_AUTH_PROFILE } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  let token = "";
  try {
    token = getCookie(event, "auth_token") ?? "";

    const { data } = await apiClient.get<UserProfileData>(API_AUTH_PROFILE, {
      withCredentials: true,
      headers: {
        Authorization: token,
      },
    });

    return data;
  } catch (error: any) {
    const err = error as AxiosError<any>;
    throw createError({
      statusCode: err.response?.status || 500,
      statusMessage: err.response?.data?.message || "get profile failed",
    });
  }
});
