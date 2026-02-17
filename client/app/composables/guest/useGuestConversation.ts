import type { ConversationListResponse, Filters } from "$shared/types";
import { API_GUEST_CONVERSATION } from "$shared/constants/api-path";

export function useGuestConversation() {
  const conversation = ref<ConversationListResponse | null>(null);
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  async function fetchGuestConversation(
    filters: Partial<Filters> = {
      limit: 20,
      page: 1,
    },
  ) {
    isLoading.value = true;
    error.value = null;

    try {
      const data = await $fetch<ConversationListResponse>(
        API_GUEST_CONVERSATION,
        {
          method: "GET",
          query: filters,
        },
      );

      conversation.value = data;
      return data;
    } catch (err: any) {
      error.value = err?.data?.message || "Failed to fetch GuestConversation";
      throw err;
    } finally {
      isLoading.value = false;
    }
  }

  return {
    conversation,
    isLoading,
    error,
    fetchGuestConversation,
  };
}
