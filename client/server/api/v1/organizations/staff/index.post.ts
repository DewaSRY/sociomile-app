import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type {  CommonResponse, RegisterRequest } from "$shared/types";
import type { AxiosResponse } from "axios";
import { ORG_STAFF } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const body = await readBody<RegisterRequest>(event);

  try {
    const { data } = await apiClient.post<
      RegisterRequest,
      AxiosResponse<CommonResponse>
    >(ORG_STAFF, body);


    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
