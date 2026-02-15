<template>
  <GuestDashboardPage />
</template>

<script setup lang="ts">
import GuestDashboardPage from "~/components/pages/guest/guest-dashboard-page.vue";

definePageMeta({
  middleware: 'guest-user',
  layout: 'portal',
});
</script>
