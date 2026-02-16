import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { CommonResponse, CreateTicketRequest } from "$shared/types";
import type { AxiosResponse } from "axios";
import { ORG_TICKET } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const body = await readBody<CreateTicketRequest>(event);

  try {
    const { data } = await apiClient.post<
      CreateTicketRequest,
      AxiosResponse<CommonResponse>
    >(ORG_TICKET, body);

    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
