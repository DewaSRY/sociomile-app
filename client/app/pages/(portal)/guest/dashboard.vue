<template>
  <UDashboardPanel id="guest-dashboard">
    <template #header>
      <UDashboardNavbar title="guest-dashboard">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
        <template #right> </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <GuestDashboardPage />
    </template>
  </UDashboardPanel>
</template>

<script setup lang="ts">
import GuestDashboardPage from "~/components/pages/guest/guest-dashboard-page.vue";

definePageMeta({
  layout: "guest-layout",
});
</script>
