<template>
  <SignupPage />
</template>

<script setup lang="ts">
import SignupPage from "~/components/pages/auth/signup-page.vue";

definePageMeta({
  layout: 'auth',
  middleware: 'guest',
});
</script>
