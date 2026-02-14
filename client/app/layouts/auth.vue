<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center p-4">
    <div class="w-full max-w-md">
      <div class="text-center mb-8">
        <NuxtLink to="/" class="text-3xl font-bold text-primary">
          SocioMile
        </NuxtLink>
      </div>
      <UCard>
        <slot />
      </UCard>
      <div class="text-center mt-4">
        <NuxtLink to="/" class="text-sm text-gray-600 hover:text-gray-900">
          Back to Home
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
