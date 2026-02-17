import type { ConversationMessagePaginate, Filters } from "$shared/types";
import { API_GUEST_CONVERSATION } from "$shared/constants/api-path";

export function useGuestMessages() {
  const messages = ref<ConversationMessagePaginate | null>(null);
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  async function fetchGuestMessages(
    id: number,
    filters: Partial<Filters> = {
      limit: 20,
      page: 1,
    },
  ) {
    
    isLoading.value = true;
    error.value = null;

    try {
      const data = await $fetch<ConversationMessagePaginate>(
        `${API_GUEST_CONVERSATION}/${id}/messages`,
        {
          method: "GET",
          query: filters,
        },
      );

      messages.value = data;
      return data;
    } catch (err: any) {
      error.value = err?.data?.message || "Failed to fetch GuestConversation";
      throw err;
    } finally {
      isLoading.value = false;
    }
  }

  return {
    messages,
    isLoading,
    error,
    fetchGuestMessages,
  };
}
