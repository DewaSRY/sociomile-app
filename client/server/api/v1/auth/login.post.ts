import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { LoginRequest, AuthResponse } from "$shared/types";
import type { AxiosResponse, AxiosError } from "axios";
import { API_AUTH_SIGNIN } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const body = await readBody<LoginRequest>(event);

  try {
    const { data } = await apiClient.post<
      LoginRequest,
      AxiosResponse<AuthResponse>
    >(API_AUTH_SIGNIN, body);

    setCookie(event, "auth_token", `Bearer ${data.token}`, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      path: "/",
      maxAge: 60 * 60 * 24 * 7,
    });

    return data.user;
  } catch (error: any) {
    const err = error as AxiosError<any>;
    throw createError({
      statusCode: err.response?.status || 500,
      statusMessage: err.response?.data?.message || "Login failed",
    });
  }
});
