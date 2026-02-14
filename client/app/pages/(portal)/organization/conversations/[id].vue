<template>
  <div>
    <div v-if="pending" class="text-center py-8">
      <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
    </div>

    <div v-else-if="conversation" class="max-w-4xl mx-auto">
      <!-- Header -->
      <UCard class="mb-4">
        <div class="flex items-center justify-between">
          <div>
            <UButton
              variant="ghost"
              icon="i-heroicons-arrow-left"
              @click="navigateTo('/organization/dashboard')"
            >
              Back
            </UButton>
            <h1 class="text-2xl font-bold mt-2">
              Conversation with {{ conversation.guest?.name || 'Guest' }}
            </h1>
            <UBadge :color="getStatusColor(conversation.status)" class="mt-2">
              {{ conversation.status }}
            </UBadge>
          </div>
        </div>
      </UCard>

      <!-- Messages -->
      <UCard class="mb-4">
        <div class="h-[500px] overflow-y-auto mb-4 space-y-4" ref="messagesContainer">
          <div v-if="loadingMessages" class="text-center py-8">
            <UIcon name="i-heroicons-arrow-path" class="animate-spin text-2xl" />
          </div>

          <div v-else-if="messages?.messages.length">
            <div
              v-for="message in messages.messages"
              :key="message.id"
              :class="[
                'flex',
                message.created_by_id === user?.id ? 'justify-end' : 'justify-start'
              ]"
            >
              <div
                :class="[
                  'max-w-[70%] rounded-lg p-3',
                  message.created_by_id === user?.id
                    ? 'bg-primary text-white'
                    : 'bg-gray-100 text-gray-900'
                ]"
              >
                <p class="text-sm font-medium mb-1">
                  {{ message.created_by?.name }}
                </p>
                <p class="whitespace-pre-wrap">{{ message.message }}</p>
                <p class="text-xs mt-1 opacity-75">
                  {{ formatDate(message.created_at) }}
                </p>
              </div>
            </div>
          </div>

          <div v-else class="text-center py-8 text-gray-500">
            No messages yet
          </div>
        </div>

        <!-- Message Input -->
        <UForm :state="messageForm" @submit="sendMessage">
          <div class="flex gap-2">
            <UTextarea
              v-model="messageForm.message"
              placeholder="Type your message..."
              :rows="3"
              class="flex-1"
              :disabled="sending"
            />
            <UButton
              type="submit"
              icon="i-heroicons-paper-airplane"
              :loading="sending"
              :disabled="!messageForm.message.trim()"
            >
              Send
            </UButton>
          </div>
        </UForm>
      </UCard>
    </div>

    <UCard v-else class="text-center py-8">
      <p class="text-red-600">Conversation not found</p>
    </UCard>
  </div>
</template>

<script setup lang="ts">
import type { ConversationResponse, ConversationMessageListResponse } from '~/types';

definePageMeta({
  layout: 'portal',
  middleware: 'organization',
});

const route = useRoute();
const { user } = useAuth();
const conversationId = computed(() => Number(route.params.id));

const messageForm = reactive({
  message: '',
  conversation_id: conversationId.value,
});

const sending = ref(false);
const messagesContainer = ref<HTMLElement | null>(null);

const { data: conversation, pending } = await useLazyAsyncData<ConversationResponse>(
  `org-conversation-${conversationId.value}`,
  () => $fetch(`/api/conversations/${conversationId.value}`)
);

const { data: messages, pending: loadingMessages, refresh: refreshMessages } = 
  await useLazyAsyncData<ConversationMessageListResponse>(
    `org-messages-${conversationId.value}`,
    () => $fetch(`/api/conversations/${conversationId.value}/messages`)
  );

const sendMessage = async () => {
  if (!messageForm.message.trim()) return;

  sending.value = true;

  try {
    await $fetch('/api/messages', {
      method: 'POST',
      body: {
        conversation_id: conversationId.value,
        message: messageForm.message,
      },
    });

    messageForm.message = '';
    await refreshMessages();

    nextTick(() => {
      if (messagesContainer.value) {
        messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
      }
    });
  } catch (error: any) {
    console.error('Failed to send message:', error);
  } finally {
    sending.value = false;
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
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
};

onMounted(() => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
  }
});
</script>
