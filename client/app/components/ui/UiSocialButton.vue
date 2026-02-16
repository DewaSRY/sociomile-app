<template>
  <UButton
    :variant="variant"
    :block="block"
    :size="size"
    :disabled="disabled"
    :loading="loading"
    class="justify-center"
    @click="$emit('click')"
  >
    <template #leading>
      <UIcon :name="icon" class="w-5 h-5" />
    </template>
    <slot>{{ label }}</slot>
  </UButton>
</template>

<script setup lang="ts">
interface Props {
  icon: string
  label?: string
  color?: string
  variant?: 'solid' | 'outline' | 'soft' | 'ghost' | 'link'
  block?: boolean
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
  disabled?: boolean
  loading?: boolean
}

withDefaults(defineProps<Props>(), {
  label: 'Continue',
  color: 'white',
  variant: 'solid',
  block: true,
  size: 'lg',
  disabled: false,
  loading: false
})

defineEmits<{
  click: []
}>()
</script>
