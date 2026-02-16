<script setup lang="ts">
import SignupForms from "~/components/forms/signup-forms.vue";
</script>

<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900 px-4 py-12"
  >
    <div class="w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
          Create Account
        </h1>
        <p class="text-gray-600 dark:text-gray-400">
          Sign up to get started with your free account
        </p>
      </div>

      <UCard>
        <SignupForms />
      </UCard>

      <!-- Sign In Link -->
      <p class="mt-6 text-center text-sm text-gray-600 dark:text-gray-400">
        Already have an account?
        <NuxtLink
          to="/auth/signin"
          class="font-medium text-primary-600 hover:text-primary-500 dark:text-primary-400"
        >
          Sign in
        </NuxtLink>
      </p>
    </div>
  </div>
</template>
