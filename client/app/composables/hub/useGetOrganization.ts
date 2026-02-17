import type { HubOrganizationPaginateResponse, Filters } from "$shared/types";
import { HUB_ORGANIZATION } from "$shared/constants/api-path";

export function useOrganizations() {
  const organizations = ref<HubOrganizationPaginateResponse | null>(null);
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  async function fetchOrganizations(
    filters: Partial<Filters> = {
      limit: 20,
      page: 1,
    },
  ) {
    isLoading.value = true;
    error.value = null;

    try {
      const data = await $fetch<HubOrganizationPaginateResponse>(
        HUB_ORGANIZATION,
        {
          method: "GET",
          query: filters,
        },
      );

      organizations.value = data;
      return data;
    } catch (err: any) {
      error.value = err?.data?.message || "Failed to fetch organizations";
      throw err;
    } finally {
      isLoading.value = false;
    }
  }

  return {
    organizations,
    isLoading,
    error,
    fetchOrganizations,
  };
}
