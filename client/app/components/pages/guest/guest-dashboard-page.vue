<template>
  <div>
    <div class="mb-8">
      <div class="flex items-center justify-between mb-4">
        <h1 class="text-3xl font-bold">My Conversations</h1>
        <UButton @click="isModalOpen = true" icon="i-heroicons-plus">
          New Conversation
        </UButton>
      </div>
      <p class="text-gray-600">Manage your conversations with organizations</p>
    </div>

    <UCard v-if="pending" class="text-center py-8">
      <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
    </UCard>

    <div v-else-if="conversations?.conversations.length" class="space-y-4">
      <UCard
        v-for="conversation in conversations.conversations"
        :key="conversation.id"
        class="hover:bg-gray-50 cursor-pointer transition-colors"
        @click="navigateTo(`/guest/conversations/${conversation.id}`)"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center gap-2 mb-2">
              <h3 class="font-semibold text-lg">
                {{ conversation.organization?.name || 'Organization' }}
              </h3>
              <UBadge :color="getStatusColor(conversation.status)">
                {{ conversation.status }}
              </UBadge>
            </div>
            <p class="text-sm text-gray-600">
              Started: {{ formatDate(conversation.created_at) }}
            </p>
            <p v-if="conversation.organization_staff" class="text-sm text-gray-500 mt-1">
              Assigned to: {{ conversation.organization_staff.name }}
            </p>
          </div>
          <UIcon name="i-heroicons-chevron-right" class="text-gray-400" />
        </div>
      </UCard>
    </div>

    <UCard v-else class="text-center py-8">
      <p class="text-gray-600">No conversations yet. Start a new one!</p>
    </UCard>

    <!-- New Conversation Modal -->
    <UModal v-model="isModalOpen">
      <UCard>
        <template #header>
          <h3 class="text-xl font-bold">Start New Conversation</h3>
        </template>

        <UForm :state="formState" @submit="handleSubmit">
          <UFormGroup label="Select Organization" name="organization_id" class="mb-4">
            <USelectMenu
              v-model="selectedOrg"
              :options="organizations?.organizations || []"
              option-attribute="name"
              value-attribute="id"
              placeholder="Choose an organization"
              :loading="loadingOrgs"
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
              Start Conversation
            </UButton>
          </div>
        </UForm>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { ConversationListResponse, OrganizationListResponse } from '~/types';

const isModalOpen = ref(false);
const selectedOrg = ref(null);
const error = ref('');
const submitting = ref(false);

const formState = computed(() => ({
  organization_id: selectedOrg.value,
}));

const { data: conversations, pending, refresh } = await useLazyAsyncData<ConversationListResponse>(
  'guest-conversations',
  () => $fetch('/api/conversations')
);

const { data: organizations, pending: loadingOrgs } = await useLazyAsyncData<OrganizationListResponse>(
  'organizations',
  () => $fetch('/api/organizations')
);

const handleSubmit = async () => {
  if (!selectedOrg.value) {
    error.value = 'Please select an organization';
    return;
  }

  submitting.value = true;
  error.value = '';

  try {
    await $fetch('/api/conversations', {
      method: 'POST',
      body: { organization_id: selectedOrg.value },
    });

    isModalOpen.value = false;
    selectedOrg.value = null;
    await refresh();
  } catch (err: any) {
    error.value = err.data?.message || 'Failed to create conversation';
  } finally {
    submitting.value = false;
  }
};

const getStatusColor = (status: string) => {
  switch (status) {
    case 'pending': return 'warning';
    case 'in_progress': return 'info';
    case 'done': return 'success';
    default: return 'neutral';
  }
};

const formatDate = (date: string | Date) => {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
};
</script>
