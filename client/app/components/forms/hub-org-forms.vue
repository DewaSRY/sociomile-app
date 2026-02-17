<template>
  <div class="w-full max-w-md mx-auto">
    <FormHeaderUi
      title="Create New Organization"
      subtitle="Create New Organization"
    />
    <UForm
      ref="formRef"
      :schema="RegisterOrganizationRequestSchema"
      :state="formState"
      class="space-y-4"
      :validate-on="['blur', 'change', 'input']"
      @submit="onSubmit"
    >
      <div class="flex flex-col gap-2">
        <TextInputUi
          v-model="formState.name"
          label="Organization Name"
          placeholder="Super CORP"
          name="name"
          required
        />

        <TextInputUi
          v-model="formState.ownerName"
          label="Owner Name"
          placeholder="John Doe"
          name="ownerName"
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

        <UButton
          type="submit"
          variant="solid"
          size="xl"
          :loading="isLoading"
          :disabled="isFormInvalid || isLoading"
        >
          <slot>Submit</slot>
        </UButton>
      </div>
    </UForm>
  </div>
</template>

<script setup lang="ts">
import PasswordInputUi from "$components/ui/password-input-ui.vue";
import TextInputUi from "$components/ui/text-input-ui.vue";
import EmailInputUi from "$components/ui/email-input-ui.vue";
import FormHeaderUi from "$components/ui/form-header-ui.vue";
import type { FormSubmitEvent, Form } from "@nuxt/ui";
import { HUB_ORGANIZATION } from "$shared/constants/api-path";

import type { CommonResponse, RegisterOrganizationRequest } from "$shared/types";
import { RegisterOrganizationRequestSchema } from "$shared/types";

const formRef = ref<Form<RegisterOrganizationRequest> | null>(null);
const defaultState: Partial<RegisterOrganizationRequest> = {};
const isLoading = ref(false);
const formState = reactive<Partial<RegisterOrganizationRequest>>(defaultState);

const emit = defineEmits<{
  (e: "onSubmit", value: CommonResponse): void
}>()

const isFormInvalid = computed<boolean>(() => {
  if (!formRef.value?.dirty) {
    return false;
  }
  return (formRef.value?.errors?.length ?? 0) > 0;
});

const toast = useToast();

async function onSubmit(event: FormSubmitEvent<RegisterOrganizationRequest>) {
  isLoading.value = true;

  try {
    const body: RegisterOrganizationRequest = {
      email: event.data.email,
      password: event.data.password,
      name: event.data.name,
      ownerName: event.data.ownerName,
    };

    const data = await $fetch<CommonResponse>(HUB_ORGANIZATION, {
      method: "POST",
      body,
    });

    toast.add({
      title: "Success!",
      description: "Your account has been created. Welcome to ArusKu!",
      color: "success",
      icon: "i-heroicons-check-circle",
    });
    
    emit("onSubmit", data)

  } catch (error) {
    toast.add({
      title: "Failed to Create New Organization",
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
