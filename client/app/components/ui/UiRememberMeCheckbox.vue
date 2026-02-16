<template>
  <UCheckbox
    :model-value="modelValue"
    :label="label"
    :disabled="disabled"
    @update:model-value="$emit('update:modelValue',Boolean( $event))"
  />
</template>

<script setup lang="ts">
interface Props {
  modelValue?: boolean
  label?: string
  disabled?: boolean
  color?: string
}

withDefaults(defineProps<Props>(), {
  modelValue: false,
  label: 'Remember me',
  disabled: false,
  color: 'primary'
})

defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>
