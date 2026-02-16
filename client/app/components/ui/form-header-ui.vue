<template>
  <div class="text-center mb-8">
    <UIcon 
      :name="icon" 
      :class="`w-${iconSize} h-${iconSize} mx-auto mb-4 text-${iconColor}`" 
    />
    <h1 :class="`text-${titleSize} font-bold text-gray-900 dark:text-white mb-2`">
      {{ title }}
    </h1>
    <p class="text-gray-600 dark:text-gray-400">
      {{ subtitle }}
    </p>
  </div>
</template>

<script setup lang="ts">
interface Props {
  icon?: string
  iconSize?: string
  iconColor?: string
  title: string
  subtitle?: string
  titleSize?: string
}

withDefaults(defineProps<Props>(), {
  icon: 'i-heroicons-currency-dollar',
  iconSize: '16',
  iconColor: 'primary',
  subtitle: '',
  titleSize: '3xl'
})
</script>
