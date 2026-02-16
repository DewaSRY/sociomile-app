<template>
  <div class="w-full max-w-md mx-auto">
    <FormHeaderUi
      title="Create Account"
      subtitle="Sign up to get started with your free account"
    />


    <UCard
      :ui="{
        body: 'p-6 sm:p-8 space-y-6',
      }"
      as="div"
      variant="soft"
    >
      <UForm
        ref="formRef"
        :schema="SignupFormSchema"
        :state="formState"
        class="space-y-4"
        :validate="validate"
        :validate-on="['blur', 'change', 'input']"
        @submit="onSubmit"
      >
        <div class="flex flex-col gap-2">
          <TextInputUi
            v-model="formState.name"
            label="Full Name"
            placeholder="John Doe"
            name="username"
            required
          />

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
            placeholder="Create a strong password"
            name="password"
            help="Must be at least 6 characters"
            required
          />

          <PasswordInputUi
            v-model="formState.confirmPassword"
            label="Confirm Password"
            placeholder="Re-enter your password"
            name="confirmPassword"
            required
          />

          <!-- Sign Up Button -->
          <SubmitButtonUi
            label="Create Account"
            :loading="isLoading"
            :disabled="isFormInvalid || isLoading"
          />
        </div>
      </UForm>
    </UCard>

    <div class="text-center mt-6">
      <span class="text-gray-600 dark:text-gray-400">
        Already have an account?
      </span>
      <UButton
        :to="'/auth/signin'"
        variant="link"
        color="primary"
        :padded="false"
        class="ml-1"
      >
        Sign in
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import PasswordInputUi from "$components/ui/password-input-ui.vue";
import TextInputUi from "$components/ui/text-input-ui.vue";
import EmailInputUi from "$components/ui/email-input-ui.vue";
import SubmitButtonUi from "$components/ui/submit-button-ui.vue";
import FormHeaderUi from "$components/ui/form-header-ui.vue";
import type { FormSubmitEvent, FormError, Form  } from "@nuxt/ui";
import { API_AUTH_SIGNUP } from "$shared/constants/api-path";

import type {  SignupForm, RegisterRequest} from "$shared/types";
import { SignupFormSchema } from "$shared/types";

const formRef = ref<Form<SignupForm> | null>(null)
const defaultState: Partial<SignupForm> = {};
const isLoading = ref(false);
const formState = reactive<Partial<SignupForm> >(defaultState);
  
const isFormInvalid = computed<boolean>(() => {
  if(!formRef.value?.dirty){
    return false
  }
  return (formRef.value?.errors?.length ?? 0) > 0
})

const toast = useToast();


function validate(state: Partial<SignupForm>): FormError[] {
  const errors:FormError<string>[] = []
  if (!state.email) errors.push({ name: 'email', message: 'Required' })
  if (!state.password) errors.push({ name: 'password', message: 'Required' })
  if(state.password && state.confirmPassword){
    if(state.password !== state.confirmPassword){
      errors.push({name: 'confirmPassword', message: "Passwords do not match"})
    }
  } 
  return errors
}


async function onSubmit(event: FormSubmitEvent<SignupForm>) {
  isLoading.value = true;

  try {
    const body: RegisterRequest = {
      email: event.data.email,
      password: event.data.password,
      name: event.data.name,
    };

    await $fetch(API_AUTH_SIGNUP, {
      method: "POST",
      body,
    });

    toast.add({
      title: "Success!",
      description: "Your account has been created. Welcome to ArusKu!",
      color: "success",
      icon: "i-heroicons-check-circle",
    });

    await navigateTo("/dashboard");
  } catch (error) {
    toast.add({
      title: "Sign Up Failed",
      description: "Unable to create account. Please try again.",
      color: "error",
      icon: "i-heroicons-x-circle",
    });
  } finally {
    isLoading.value = false;
    Object.assign(formState, defaultState);
  }
}
</script>
