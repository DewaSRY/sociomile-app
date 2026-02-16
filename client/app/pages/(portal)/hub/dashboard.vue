<template>
  <UDashboardPanel id="hub-dashboard">
    <template #header>
      <UDashboardNavbar title="hub-dashboard">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
        <template #right> </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <HubDashboardPage />
    </template>
  </UDashboardPanel>
</template>

<script setup lang="ts">
import HubDashboardPage from "~/components/pages/hub/hub-dashboard-page.vue";

definePageMeta({
  layout: "hub-layout",
});
</script>
