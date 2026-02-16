import { defineEventHandler, setCookie, getCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type {
  AuthResponse,
  RefreshTokenRequest,
} from "$shared/types";
import type { AxiosResponse } from "axios";

import { API_AUTH_REFRESH } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  try {
    const token = getCookie(event, "auth_token");

    const { data } = await apiClient.post<
      RefreshTokenRequest,
      AxiosResponse<AuthResponse>
    >(
      API_AUTH_REFRESH,
      {
        token: token,
      },
      {
        withCredentials: true,
      },
    );

    setCookie(event, "auth_token", `Bearer ${data.token}`, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      path: "/",
      maxAge: 60 * 60 * 24 * 7,
    });

    return { success: true };
  } catch (error) {
    throw createError({
      statusCode: 401,
      statusMessage: "Refresh failed",
    });
  }
});
