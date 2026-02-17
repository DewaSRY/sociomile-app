<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { breakpointsTailwind, useBreakpoints } from "@vueuse/core";
import GuestConversationList from "~/components/conversations/guest-conversation-list.vue";
import GuestConversationBubble from "~/components/conversations/guest-conversation-bubble.vue";
import { useGuestConversation } from "~/composables/guest/useGuestConversation";
const { conversation, fetchGuestConversation, isLoading } =
  useGuestConversation();

const conversationList = computed(() => conversation.value?.data ?? []);
// import type { Mail } from '~/types'

// const selectedTab = ref('all')

// const { data: mails } = await useFetch<Mail[]>('/api/mails', { default: () => [] })

// Filter mails based on the selected tab
// const filteredMails = computed(() => {
//   if (selectedTab.value === 'unread') {
//     return mails.value.filter(mail => !!mail.unread)
//   }

//   return mails.value
// })

const selectedConversation = ref<ConversationResponse | null>();

// const isMailPanelOpen = computed({
//   get() {
//     return !!selectedMail.value
//   },
//   set(value: boolean) {
//     if (!value) {
//       selectedMail.value = null
//     }
//   }
// })

// Reset selected mail if it's not in the filtered mails
// watch(filteredMails, () => {
//   if (!filteredMails.value.find(mail => mail.id === selectedMail.value?.id)) {
//     selectedMail.value = null
//   }
// })

const breakpoints = useBreakpoints(breakpointsTailwind);
const isMobile = breakpoints.smaller("lg");

await fetchGuestConversation();
</script>

<template>
  <UDashboardPanel
    id="inbox-1"
    :default-size="25"
    :min-size="20"
    :max-size="30"
    resizable
  >
    <UDashboardNavbar title="Inbox">
      <template #right> <UDashboardSidebarCollapse /> </template>
    </UDashboardNavbar>

    <GuestConversationList
      v-model="selectedConversation"
      v-if="!isLoading"
      :conversation="conversationList"
    />
  </UDashboardPanel>

  <template v-if="selectedConversation">
    <GuestConversationBubble :id="selectedConversation.id" />
  </template>

  <!-- <InboxMail v-if="selectedMail" :mail="selectedMail" @close="selectedMail = null" />

  <div v-else class="hidden lg:flex flex-1 items-center justify-center">
    <UIcon name="i-lucide-inbox" class="size-32 text-dimmed" />
  </div> -->

  <!-- v-model:open="true" -->
  <ClientOnly>
    <USlideover v-if="isMobile">
      <template #content>
        <template v-if="selectedConversation">
          <GuestConversationBubble :id="selectedConversation.id" />
        </template>
        <!-- <InboxMail v-if="selectedMail" :mail="selectedMail" @close="selectedMail = null" /> -->
      </template>
    </USlideover>
  </ClientOnly>
</template>
