<template>
  <UFormField
    :label="label"
    :name="name"
    :required="required"
    :help="help"
  >
    <UInput
      :model-value="modelValue"
      type="email"
      :placeholder="placeholder"
      :size="size"
      :icon="icon"
      :disabled="disabled"
      @update:model-value="$emit('update:modelValue', $event)"
      class="w-full"
    />
  </UFormField>
</template>

<script setup lang="ts">
interface Props {
  modelValue?: string
  label?: string
  name: string
  placeholder?: string
  required?: boolean
  disabled?: boolean
  size?: 'sm' | 'md' | 'lg' | 'xl'
  icon?: string
  help?: string
}

withDefaults(defineProps<Props>(), {
  label: 'Email',
  placeholder: 'Enter your email',
  required: false,
  disabled: false,
  size: 'xl',
  icon: 'i-heroicons-envelope',
  help: '',
})

defineEmits<{
  'update:modelValue': [value: string]
}>()
</script>
