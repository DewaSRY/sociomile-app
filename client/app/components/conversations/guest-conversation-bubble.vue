<script setup lang="ts">
import { format } from "date-fns";
import { useGuestMessages } from "~/composables/guest/useGuestMessages";
import MessageBubble from "../ui/message-bubble.vue";
const { fetchGuestMessages, messages } = useGuestMessages();
// import type { Mail } from '~/types'

const props = defineProps<{
  id: number;
}>();

const emits = defineEmits(["close"]);

const dropdownItems = [
  [
    {
      label: "Mark as unread",
      icon: "i-lucide-check-circle",
    },
    {
      label: "Mark as important",
      icon: "i-lucide-triangle-alert",
    },
  ],
  [
    {
      label: "Star thread",
      icon: "i-lucide-star",
    },
    {
      label: "Mute thread",
      icon: "i-lucide-circle-pause",
    },
  ],
];

const toast = useToast();

const reply = ref("");
const loading = ref(false);

function onSubmit() {
  loading.value = true;

  setTimeout(() => {
    reply.value = "";

    toast.add({
      title: "Email sent",
      description: "Your email has been sent successfully",
      icon: "i-lucide-check-circle",
      color: "success",
    });

    loading.value = false;
  }, 1000);
}
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
      <template #right>
        <UTooltip text="Archive">
          <UButton icon="i-lucide-inbox" color="neutral" variant="ghost" />
        </UTooltip>

        <UTooltip text="Reply">
          <UButton icon="i-lucide-reply" color="neutral" variant="ghost" />
        </UTooltip>

        <UDropdownMenu :items="dropdownItems">
          <UButton
            icon="i-lucide-ellipsis-vertical"
            color="neutral"
            variant="ghost"
          />
        </UDropdownMenu>
      </template>
    </UDashboardNavbar>
    <div class="h-full relative">
      <div class="h-4"/>
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
