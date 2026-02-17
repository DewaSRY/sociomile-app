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
        </div>
        <div class="flex gap-2 ">
          <UBadge
            v-if="organizationName"
            color="neutral"
            variant="outline"
            size="sm"
          >
            {{ organizationName }}
          </UBadge>

            <UBadge color="primary" variant="soft" size="sm"  class="block xl:hidden">
            {{ userRole }}
          </UBadge>
        </div>
      </div>
    </div>
  </div>
</template>
