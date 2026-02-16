import { defineEventHandler } from "h3";
import { apiClient } from "$shared/lib/api-client";
import type { TicketListResponse, Filters } from "$shared/types";
import type { AxiosResponse } from "axios";
import { ORG_TICKET } from "$shared/constants/api-path";

export default defineEventHandler(async (event) => {
  const token = getCookie(event, "auth_token");
  const query = getQuery(event) as Partial<Filters>;
  try {
    const { data } = await apiClient.get<
      AxiosResponse<TicketListResponse>
    >(ORG_TICKET, {
      headers: {
        Authorization: token,
      },
      params: query,
    });

    return data;
  } catch (err: any) {
    return {
      error: true,
      message: err.response?.data?.message || "Login failed",
    };
  }
});
