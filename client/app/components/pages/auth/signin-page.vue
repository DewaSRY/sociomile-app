<template>
  <div>
    <h1 class="text-2xl font-bold text-center mb-6">Sign In</h1>
    
    <UForm :state="formState" :schema="loginRequestSchema" @submit="handleSubmit">
      <UFormGroup label="Email" name="email" class="mb-4">
        <UInput
          v-model="formState.email"
          type="email"
          placeholder="Enter your email"
          icon="i-heroicons-envelope"
        />
      </UFormGroup>

      <UFormGroup label="Password" name="password" class="mb-6">
        <UInput
          v-model="formState.password"
          type="password"
          placeholder="Enter your password"
          icon="i-heroicons-lock-closed"
        />
      </UFormGroup>

      <UAlert
        v-if="error"
        variant="soft"
        :title="error"
        class="mb-4"
      />

      <UButton type="submit" block :loading="loading">
        Sign In
      </UButton>
    </UForm>

    <div class="mt-6 text-center text-sm text-gray-600">
      Don't have an account?
      <NuxtLink to="/signup" class="text-primary hover:underline font-medium">
        Sign Up
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { LoginRequest } from '~/types';
import { loginRequestSchema } from '~/types';

definePageMeta({
  layout: 'auth',
  middleware: 'guest',
});

const { setAuth, isSuperAdmin, isOrganizationOwner, isOrganizationSales, isGuest } = useAuth();

const formState = reactive<LoginRequest>({
  email: '',
  password: '',
});

const loading = ref(false);
const error = ref('');

const handleSubmit = async () => {
  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch('/api/auth/login', {
      method: 'POST',
      body: formState,
    });

    setAuth(response);

    if (isSuperAdmin()) {
      navigateTo('/hub/dashboard');
    } else if (isOrganizationOwner() || isOrganizationSales()) {
      navigateTo('/organization/dashboard');
    } else if (isGuest()) {
      navigateTo('/guest/dashboard');
    } else {
      navigateTo('/');
    }
  } catch (err: any) {
    error.value = err.data?.message || 'Login failed. Please try again.';
  } finally {
    loading.value = false;
  }
};
</script>
