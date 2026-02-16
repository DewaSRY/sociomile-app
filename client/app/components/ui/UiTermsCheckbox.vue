<template>
  <UCheckbox
    :model-value="modelValue"
    :disabled="disabled"
    @update:model-value="$emit('update:modelValue', Boolean($event))"
  >
    <template #label>
      <span class="text-sm">
        {{ preText }}
        <UButton
          v-if="linkText && linkUrl"
          :to="linkUrl"
          variant="link"
          color="primary"
          :padded="false"
          size="sm"
          class="text-sm"
        >
          {{ linkText }}
        </UButton>
      </span>
    </template>
  </UCheckbox>
</template>

<script setup lang="ts">
interface Props {
  modelValue?: boolean
  preText?: string
  linkText?: string
  linkUrl?: string
  disabled?: boolean
  color?: string
}

withDefaults(defineProps<Props>(), {
  modelValue: false,
  preText: 'I agree to the',
  linkText: 'Terms and Conditions',
  linkUrl: '/terms',
  disabled: false,
  color: 'primary'
})

defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>
