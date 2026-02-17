<script setup lang="ts">
import { useGuestMessages } from "~/composables/guest/useGuestMessages";
import MessageBubble from "../ui/message-bubble.vue";
const { fetchGuestMessages, messages } = useGuestMessages();

const props = defineProps<{
  id: number;
}>();

const emits = defineEmits(["close"]);


await fetchGuestMessages(props.id);
</script>

<template>
  <UDashboardPanel id="inbox-2" grow>
    <UDashboardNavbar title="hallo" :toggle="false">
      <template #leading>
        <UButton
          icon="i-lucide-x"
          color="neutral"
          variant="ghost"
          class="-ms-1.5"
          @click="emits('close')"
        />
      </template>
    </UDashboardNavbar>
    <div class="h-full relative">
      <div class="h-4" />
      <template v-if="messages">
        <MessageBubble
          v-for="m of messages.data"
          :sender-name="m.createdBy?.name ?? ''"
          :message="m.message"
          :created-at="m.createdAt"
          side="right"
        />
      </template>

      <div class="pb-4 px-4 sm:px-6 shrink-0"></div>
    </div>
  </UDashboardPanel>
</template>
