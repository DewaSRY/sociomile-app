<template>
  <div>
    <div class="mb-8">
      <h1 class="text-3xl font-bold mb-2">Organization Dashboard</h1>
      <p class="text-gray-600">Manage conversations and tickets</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid md:grid-cols-3 gap-4 mb-8">
      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Pending Conversations</p>
            <p class="text-3xl font-bold mt-2">
              {{ pendingCount }}
            </p>
          </div>
          <UIcon name="i-heroicons-clock" class="text-4xl text-yellow-500" />
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">In Progress</p>
            <p class="text-3xl font-bold mt-2">
              {{ inProgressCount }}
            </p>
          </div>
          <UIcon name="i-heroicons-arrow-path" class="text-4xl text-blue-500" />
        </div>
      </UCard>

      <UCard>
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm text-gray-600">Completed</p>
            <p class="text-3xl font-bold mt-2">
              {{ completedCount }}
            </p>
          </div>
          <UIcon name="i-heroicons-check-circle" class="text-4xl text-green-500" />
        </div>
      </UCard>
    </div>

    <!-- Tabs -->
    <UTabs :items="tabs" v-model="selectedTab">
      <template #conversations>
        <div v-if="loadingConversations" class="text-center py-8">
          <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
        </div>

        <div v-else-if="conversations?.conversations.length" class="space-y-4 mt-4">
          <UCard
            v-for="conversation in conversations.conversations"
            :key="conversation.id"
            class="hover:bg-gray-50 cursor-pointer transition-colors"
            @click="viewConversation(conversation.id)"
          >
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-2">
                  <h3 class="font-semibold text-lg">
                    {{ conversation.guest?.name || 'Guest' }}
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
                <p v-else class="text-sm text-orange-600 mt-1">
                  Unassigned
                </p>
              </div>
              <div class="flex gap-2">
                <UButton
                  v-if="isOrganizationOwner() && !conversation.organization_staff"
                  size="sm"
                  variant="outline"
                  @click.stop="openAssignModal(conversation)"
                >
                  Assign
                </UButton>
                <UIcon name="i-heroicons-chevron-right" class="text-gray-400" />
              </div>
            </div>
          </UCard>
        </div>

        <UCard v-else class="text-center py-8 mt-4">
          <p class="text-gray-600">No conversations found</p>
        </UCard>
      </template>

      <template #tickets>
        <div v-if="loadingTickets" class="text-center py-8">
          <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
        </div>

        <div v-else-if="tickets?.tickets.length" class="space-y-4 mt-4">
          <UCard
            v-for="ticket in tickets.tickets"
            :key="ticket.id"
            class="hover:bg-gray-50 cursor-pointer transition-colors"
          >
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-2">
                  <h3 class="font-semibold text-lg">{{ ticket.name }}</h3>
                  <UBadge :color="getStatusColor(ticket.status)">
                    {{ ticket.status }}
                  </UBadge>
                </div>
                <p class="text-sm text-gray-600">
                  Ticket #{{ ticket.ticket_number }}
                </p>
                <p class="text-sm text-gray-500 mt-1">
                  Created by: {{ ticket.created_by?.name }} on {{ formatDate(ticket.created_at) }}
                </p>
              </div>
              <UIcon name="i-heroicons-chevron-right" class="text-gray-400" />
            </div>
          </UCard>
        </div>

        <UCard v-else class="text-center py-8 mt-4">
          <p class="text-gray-600">No tickets found</p>
        </UCard>
      </template>
    </UTabs>

    <!-- Assign Modal -->
    <UModal v-model="isAssignModalOpen">
      <UCard>
        <template #header>
          <h3 class="text-xl font-bold">Assign Conversation</h3>
        </template>

        <UForm :state="assignForm" @submit="handleAssign">
          <UFormGroup label="Assign to Staff" name="organization_staff_id" class="mb-4">
            <UInput
              v-model.number="assignForm.organization_staff_id"
              type="number"
              placeholder="Enter staff user ID"
            />
          </UFormGroup>

          <UAlert
            v-if="assignError"
            color="error"
            variant="soft"
            :title="assignError"
            class="mb-4"
          />

          <div class="flex justify-end gap-2">
            <UButton variant="outline" @click="isAssignModalOpen = false">
              Cancel
            </UButton>
            <UButton type="submit" :loading="assigning">
              Assign
            </UButton>
          </div>
        </UForm>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth';
import type { ConversationListResponse, TicketListResponse, ConversationResponse } from '~/types';

const { isOrganizationOwner } = useAuth();

const selectedTab = ref(0);
const tabs = [
  { label: 'Conversations', slot: 'conversations' },
  { label: 'Tickets', slot: 'tickets' },
];

const isAssignModalOpen = ref(false);
const selectedConversation = ref<ConversationResponse | null>(null);
const assignError = ref('');
const assigning = ref(false);

const assignForm = reactive({
  organization_staff_id: null as number | null,
});

const { data: conversations, pending: loadingConversations, refresh: refreshConversations } = 
  await useLazyAsyncData<ConversationListResponse>(
    'org-conversations',
    () => $fetch('/api/conversations')
  );

const { data: tickets, pending: loadingTickets } = await useLazyAsyncData<TicketListResponse>(
  'org-tickets',
  () => $fetch('/api/tickets')
);

const pendingCount = computed(() => {
  return conversations.value?.conversations.filter(c => c.status === 'pending').length || 0;
});

const inProgressCount = computed(() => {
  return conversations.value?.conversations.filter(c => c.status === 'in_progress').length || 0;
});

const completedCount = computed(() => {
  return conversations.value?.conversations.filter(c => c.status === 'done').length || 0;
});

const openAssignModal = (conversation: ConversationResponse) => {
  selectedConversation.value = conversation;
  isAssignModalOpen.value = true;
};

const handleAssign = async () => {
  if (!selectedConversation.value || !assignForm.organization_staff_id) {
    assignError.value = 'Please enter a staff ID';
    return;
  }

  assigning.value = true;
  assignError.value = '';

  try {
    await $fetch(`/api/conversations/${selectedConversation.value.id}/assign`, {
      method: 'POST',
      body: { organization_staff_id: assignForm.organization_staff_id },
    });

    isAssignModalOpen.value = false;
    assignForm.organization_staff_id = null;
    await refreshConversations();
  } catch (err: any) {
    assignError.value = err.data?.message || 'Failed to assign conversation';
  } finally {
    assigning.value = false;
  }
};

const viewConversation = (id: number) => {
  navigateTo(`/organization/conversations/${id}`);
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
