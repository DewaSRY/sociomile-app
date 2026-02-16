<template>
  <UDashboardPanel id="guest-conversation-dashboard">
    <template #header>
      <UDashboardNavbar title="guest-conversation-dashboard">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
        <template #right> </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <guestChatRomePage />
    </template>
  </UDashboardPanel>
</template>

<script setup lang="ts">
import guestChatRomePage from "~/components/pages/guest/guest-chat-rome.page.vue";

definePageMeta({
  layout: "guest-layout",
});
</script>
