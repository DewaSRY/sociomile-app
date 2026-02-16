<script setup lang="ts">
interface Props {
  name: string;
  userRole: string;
  organizationName?: string;
  collapsed?: boolean;
}
const props = defineProps<Props>();

const initials = computed(() => {
  if (!props.name) return "";
  return props.name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .toUpperCase();
});
</script>

<template>
  <div class="w-full">
    <div v-if="collapsed" class="flex justify-center py-2">
      <UTooltip :text="name">
        <UAvatar :alt="name" :text="initials" size="md" />
      </UTooltip>
    </div>

    <div v-else class="flex items-center gap-3">
      <UAvatar :alt="name" :text="initials" size="lg" />

      <div class="flex-1 min-w-0">
        <div class="flex justify-between">
          <p class="font-semibold truncate">
            {{ name }}
          </p>
          <UBadge color="primary" variant="soft" size="xs">
            {{ userRole }}
          </UBadge>
        </div>



        <div class="flex items-center gap-2 mt-1 flex-wrap">
          <UBadge
            v-if="organizationName"
            color="neutral"
            variant="outline"
            size="xs"
          >
            {{ organizationName }}
          </UBadge>
        </div>
      </div>
    </div>
  </div>
</template>
