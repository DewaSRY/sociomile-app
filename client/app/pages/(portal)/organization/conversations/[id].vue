<template>
  <UDashboardPanel id="org-conversation-dashboard">
    <template #header>
      <UDashboardNavbar title="org-conversation-dashboard">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
        <template #right> </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <organizationChatRoomPage />
    </template>
  </UDashboardPanel>
</template>

<script setup lang="ts">
import organizationChatRoomPage from "~/components/pages/organization/organization-chat-room-page.vue";
definePageMeta({
  layout: "org-layout",
});
</script>
