<template>
  <UButton
    type="submit"
    :variant="variant"
    :block="block"
    :size="size"
    :loading="loading"
    :disabled="disabled || loading"
  >
    <slot>{{ label }}</slot>
  </UButton>
</template>

<script setup lang="ts">
interface Props {
  label?: string
  color?: string
  variant?: 'solid' | 'outline' | 'soft' | 'ghost' | 'link'
  block?: boolean
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
  loading?: boolean
  disabled?: boolean
}

withDefaults(defineProps<Props>(), {
  label: 'Submit',
  color: 'primary',
  variant: 'solid',
  size: 'xl',
  block: true,
  loading: false,
  disabled: false
})
</script>
