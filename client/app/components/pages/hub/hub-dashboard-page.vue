<template>
  <div>
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <h1 class="text-3xl font-bold">Hub Dashboard</h1>
        <UButton @click="isModalOpen = true" icon="i-heroicons-plus">
          Create Organization
        </UButton>
      </div>
      <p class="text-gray-600">Manage organizations and users</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid md:grid-cols-3 gap-4 mb-8">
      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Total Organizations</p>
            <p class="text-3xl font-bold mt-2">
              {{ organizations?.metadata.total || 0 }}
            </p>
          </div>
          <UIcon name="i-heroicons-building-office" class="text-4xl text-primary" />
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Active Conversations</p>
            <p class="text-3xl font-bold mt-2">
              {{ conversations?.metadata.total || 0 }}
            </p>
          </div>
          <UIcon name="i-heroicons-chat-bubble-left-right" class="text-4xl text-blue-500" />
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Total Tickets</p>
            <p class="text-3xl font-bold mt-2">
              {{ tickets?.metadata.total || 0 }}
            </p>
          </div>
          <UIcon name="i-heroicons-ticket" class="text-4xl text-green-500" />
        </div>
      </UCard>
    </div>

    <!-- Organizations List -->
    <UCard>
      <template #header>
        <h2 class="text-xl font-bold">Organizations</h2>
      </template>

      <div v-if="pending" class="text-center py-8">
        <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
      </div>

      <div v-else-if="organizations?.organizations.length" class="space-y-4">
        <div
          v-for="org in organizations.organizations"
          :key="org.id"
          class="flex items-center justify-between p-4 border rounded-lg hover:bg-gray-50"
        >
          <div>
            <h3 class="font-semibold text-lg">{{ org.name }}</h3>
            <p class="text-sm text-gray-600">
              Owner: {{ org.owner?.name || 'N/A' }}
            </p>
            <p class="text-xs text-gray-500 mt-1">
              Created: {{ formatDate(org.created_at) }}
            </p>
          </div>
          <UButton variant="ghost" icon="i-heroicons-eye">
            View
          </UButton>
        </div>
      </div>

      <div v-else class="text-center py-8 text-gray-500">
        No organizations yet. Create one to get started!
      </div>
    </UCard>

    <!-- Create Organization Modal -->
    <UModal v-model="isModalOpen">
      <UCard>
        <template #header>
          <h3 class="text-xl font-bold">Create Organization</h3>
        </template>

        <UForm :state="formState" @submit="handleSubmit">
          <UFormGroup label="Organization Name" name="name" class="mb-4">
            <UInput
              v-model="formState.name"
              placeholder="Enter organization name"
            />
          </UFormGroup>

          <UFormGroup label="Owner User ID" name="owner_id" class="mb-4">
            <UInput
              v-model.number="formState.owner_id"
              type="number"
              placeholder="Enter owner user ID"
            />
          </UFormGroup>

          <UAlert
            v-if="error"
            color="error"
            variant="soft"
            :title="error"
            class="mb-4"
          />

          <div class="flex justify-end gap-2">
            <UButton variant="outline" @click="isModalOpen = false">
              Cancel
            </UButton>
            <UButton type="submit" :loading="submitting">
              Create Organization
            </UButton>
          </div>
        </UForm>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { OrganizationListResponse, ConversationListResponse, TicketListResponse } from '~/types';

const isModalOpen = ref(false);
const error = ref('');
const submitting = ref(false);

const formState = reactive({
  name: '',
  owner_id: null as number | null,
});

const { data: organizations, pending, refresh } = await useLazyAsyncData<OrganizationListResponse>(
  'hub-organizations',
  () => $fetch('/api/organizations')
);

const { data: conversations } = await useLazyAsyncData<ConversationListResponse>(
  'hub-conversations',
  () => $fetch('/api/conversations')
);

const { data: tickets } = await useLazyAsyncData<TicketListResponse>(
  'hub-tickets',
  () => $fetch('/api/tickets')
);

const handleSubmit = async () => {
  if (!formState.name || !formState.owner_id) {
    error.value = 'Please fill in all fields';
    return;
  }

  submitting.value = true;
  error.value = '';

  try {
    await $fetch('/api/organizations', {
      method: 'POST',
      body: formState,
    });

    isModalOpen.value = false;
    formState.name = '';
    formState.owner_id = null;
    await refresh();
  } catch (err: any) {
    error.value = err.data?.message || 'Failed to create organization';
  } finally {
    submitting.value = false;
  }
};

const formatDate = (date: string | Date) => {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  });
};
</script>
