<template>
  <OrganizationDashboardPage />
</template>

<script setup lang="ts">
import OrganizationDashboardPage from "~/components/pages/organization/organization-page.vue";
definePageMeta({
  middleware: 'organization',
  layout: 'portal',
});</script>
