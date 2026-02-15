<template>
  <GuestDashboardPage />
</template>

<script setup lang="ts">
import GuestDashboardPage from "~/components/pages/guest/guest-dashboard-page.vue";

definePageMeta({
  layout: 'portal',
  middleware: 'guest-user',
});
</script>
