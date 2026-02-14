<template>
  <div>
    <h1 class="text-2xl font-bold text-center mb-6">Sign Up</h1>
    
    <UForm :state="formState" :schema="registerRequestSchema" @submit="handleSubmit">
      <UFormGroup label="Full Name" name="name" class="mb-4">
        <UInput
          v-model="formState.name"
          type="text"
          placeholder="Enter your full name"
          icon="i-heroicons-user"
        />
      </UFormGroup>

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
          placeholder="Enter your password (min 6 characters)"
          icon="i-heroicons-lock-closed"
        />
      </UFormGroup>

      <UAlert
        v-if="error"
        color="error"
        variant="soft"
        :title="error"
        class="mb-4"
      />

      <UButton type="submit" block :loading="loading">
        Sign Up
      </UButton>
    </UForm>

    <div class="mt-6 text-center text-sm text-gray-600">
      Already have an account?
      <NuxtLink to="/signin" class="text-primary hover:underline font-medium">
        Sign In
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { RegisterRequest } from '~/types';
import { registerRequestSchema } from '~/types';

const { setAuth } = useAuth();

const formState = reactive<RegisterRequest>({
  name: '',
  email: '',
  password: '',
});

const loading = ref(false);
const error = ref('');

const handleSubmit = async () => {
  loading.value = true;
  error.value = '';

  try {
    const response = await $fetch('/api/auth/register', {
      method: 'POST',
      body: formState,
    });

    setAuth(response);
    navigateTo('/guest/dashboard');
  } catch (err: any) {
    error.value = err.data?.message || 'Registration failed. Please try again.';
  } finally {
    loading.value = false;
  }
};
</script>