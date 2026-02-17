<script setup lang="ts">
import { format, isToday } from "date-fns";
import type { ConversationResponse } from "$shared/types";

const props = defineProps<{
  conversation: ConversationResponse[];
}>();
const selectedConversation = defineModel<ConversationResponse | null>();
const mailsRefs = ref<Record<number, Element | null>>({});
</script>

<template>
  <div class="overflow-y-auto divide-y divide-default">
    <div
      v-for="con in conversation"
      :key="con.id"
      :ref="
        (el) => {
          mailsRefs[con.id] = el as Element | null;
        }
      "
    >
      <div
        class="p-4 sm:px-6 text-sm cursor-pointer border-l-2 transition-colors"
        @click="selectedConversation = con"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            {{ con.organization?.name }}
          </div>
          <span>{{
            isToday(new Date(con.updatedAt))
              ? format(new Date(con.updatedAt), "HH:mm")
              : format(new Date(con.updatedAt), "dd MMM")
          }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
