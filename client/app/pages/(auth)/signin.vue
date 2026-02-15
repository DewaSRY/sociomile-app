<template>
  <SigninPage />
</template>

<script setup lang="ts">
import SigninPage from "~/components/pages/auth/signin-page.vue";

definePageMeta({
  layout: 'auth',
  middleware: 'guest',
});
</script>
