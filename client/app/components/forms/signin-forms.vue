<template>
  <div class="w-full max-w-md mx-auto">
    <FormHeaderUi
      title="Welcome Back"
      subtitle="Sign in to your account to continue"
    />
    <UCard
      :ui="{
        body: 'p-6 sm:p-8 space-y-6',
      }"
      as="div"
      variant="subtle"
    >
      <UForm
        :schema="LoginRequestSchema"
        :state="formState"
        class="space-y-4"
        @submit="onSubmit"
      >
        <div class="flex flex-col gap-2">
          <EmailInputUi
            v-model="formState.email"
            label="Email Address"
            placeholder="you@example.com"
            name="email"
            required
          />

          <PasswordInputUi
            v-model="formState.password"
            label="Password"
            placeholder="Enter your password"
            name="password"
            required
          />

          <SubmitButtonUi
            label="Sign In"
            :loading="isLoading"
          />
        </div>
      </UForm>
    </UCard>

    <div class="text-center mt-6">
      <span class="text-gray-600 dark:text-gray-400">
        Don't have an account?
      </span>
      <UButton
        :to="'/auth/signup'"
        variant="link"
        color="primary"
        :padded="false"
        class="ml-1"
      >
        Sign up for free
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import PasswordInputUi from "$components/ui/password-input-ui.vue";
import EmailInputUi from "$components/ui/email-input-ui.vue";
import SubmitButtonUi from "$components/ui/submit-button-ui.vue";
import FormHeaderUi from "$components/ui/form-header-ui.vue";
import type { FormSubmitEvent, FormError, Form } from "@nuxt/ui";

import type { LoginRequest } from "$shared/types";
import { LoginRequestSchema } from "$shared/types";
import { API_AUTH_SIGNIN } from "$shared/constants/api-path";

const defaultState: Partial<LoginRequest> = {};
const formState = reactive<Partial<LoginRequest>>({
  password: undefined,
  email: undefined,
});

const toast = useToast();
const isLoading = ref(false);
const formRef = ref<Form<LoginRequest> | null>(null);

const isFormInvalid = computed<boolean>(() => {
  if (!formRef.value?.dirty) {
    return false;
  }
  return (formRef.value?.errors?.length ?? 0) > 0;
});

async function onSubmit(event: FormSubmitEvent<LoginRequest>) {
  isLoading.value = true;

  try {
    const body: LoginRequest = {
      email: event.data.email,
      password: event.data.password,
    };

    await $fetch(API_AUTH_SIGNIN, {
      method: "POST",
      body,
    });

    toast.add({
      title: "Success!",
      description: "Welcome back to ArusKu!",
      color: "success",
      icon: "i-heroicons-check-circle",
    });

    await navigateTo("/dashboard");
  } catch (error) {
    toast.add({
      title: "Sign In Failed",
      description: "Invalid email or password. Please try again.",
      color: "error",
      icon: "i-heroicons-x-circle",
    });
  } finally {
    isLoading.value = false;
    Object.assign(formState, defaultState);
  }
}
</script>
