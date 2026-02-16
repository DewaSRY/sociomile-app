import { defineEventHandler, readBody, setCookie } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { CommonResponse, UpdateTicketRequest } from "$shared/types";
import type { AxiosResponse } from "axios";
import { ORG_TICKET } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
     const id = getRouterParam(event, "id");
  const body = await readBody<UpdateTicketRequest>(event);

  try {
    const { data } = await apiClient.put<
      UpdateTicketRequest,
      AxiosResponse<CommonResponse>
    >(`${ORG_TICKET}/${id}`, body);

    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
