<template>
  <UDashboardPanel id="org-dashboard">
    <template #header>
      <UDashboardNavbar title="org-dashboard">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
        <template #right> </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <OrganizationDashboardPage />
    </template>
  </UDashboardPanel>
</template>

<script setup lang="ts">
import OrganizationDashboardPage from "~/components/pages/organization/organization-page.vue";
definePageMeta({
  layout: "org-layout",
});
</script>
