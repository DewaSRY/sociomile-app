<template>
  <UFormField
    :label="label"
    :name="name"
    :required="required"
    :help="help"
  >
    <UInput
      :model-value="modelValue"
      :type="showPassword ? 'text' : 'password'"
      :placeholder="placeholder"
      :size="size"
      :icon="icon"
      :disabled="disabled"
      @update:model-value="$emit('update:modelValue', $event)"
      class="w-full"
    >
      <template #trailing>
        <UButton
          :icon="showPassword ? 'i-heroicons-eye-slash' : 'i-heroicons-eye'"
          variant="link"
          :padded="false"
          @click="togglePassword"
        />
      </template>
    </UInput>
  </UFormField>
</template>

<script setup lang="ts">
interface Props {
  modelValue?: string
  label?: string
  name?: string
  placeholder?: string
  required?: boolean
  disabled?: boolean
  size?: 'sm' | 'md' | 'lg' | 'xl'
  icon?: string
  help?: string
}

withDefaults(defineProps<Props>(), {
  modelValue: '',
  label: 'Password',
  name: 'password',
  placeholder: 'Enter your password',
  required: false,
  disabled: false,
  size: 'xl',
  icon: 'i-heroicons-lock-closed',
  help: '',
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
}>()

const showPassword = ref(false)

const togglePassword = () => {
  showPassword.value = !showPassword.value
}
</script>
