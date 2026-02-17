<script setup lang="ts">
import { computed } from "vue";

interface Props {
  senderName: string;
  message: string;
  createdAt: string | Date;
  side?: "left" | "right";
}

const props = withDefaults(defineProps<Props>(), {
  side: "left",
});

const isRight = computed(() => props.side === "right");

const formattedTime = computed(() => {
  const date = new Date(props.createdAt);
  return date.toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
  });
});
</script>

<template>
  <div
    class="flex w-full mb-3"
    :class="isRight ? 'justify-end' : 'justify-start'"
  >
    <div
      class="max-w-xs md:max-w-md px-4 py-2 rounded-2xl shadow-sm transition-all duration-200"
      :class="
        isRight
          ? 'bg-blue-500 text-white rounded-br-none'
          : 'bg-gray-200 text-gray-900 rounded-bl-none'
      "
    >
      <p
        v-if="!isRight"
        class="text-xs font-semibold mb-1 opacity-70"
      >
        {{ senderName }}
      </p>

      <p class="text-sm whitespace-pre-line">
        {{ message }}
      </p>

      <p class="text-[10px] mt-1 text-right opacity-60">
        {{ formattedTime }}
      </p>
    </div>
  </div>
</template>
