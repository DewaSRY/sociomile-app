<template>
  <div class="min-h-screen bg-gray-50">
    <header class="bg-white shadow-sm border-b">
      <nav class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-8">
            <NuxtLink to="/" class="text-2xl font-bold text-primary">
              SocioMile
            </NuxtLink>
            <div class="hidden md:flex items-center gap-4">
              <template v-if="isSuperAdmin()">
                <UButton to="/hub/dashboard" variant="ghost">Dashboard</UButton>
                <UButton to="/hub/organizations" variant="ghost">Organizations</UButton>
              </template>
              <template v-else-if="isOrganizationOwner() || isOrganizationSales()">
                <UButton to="/organization/dashboard" variant="ghost">Dashboard</UButton>
                <UButton to="/organization/conversations" variant="ghost">Conversations</UButton>
                <UButton to="/organization/tickets" variant="ghost">Tickets</UButton>
              </template>
              <template v-else-if="isGuest()">
                <UButton to="/guest/dashboard" variant="ghost">Dashboard</UButton>
                <UButton to="/guest/conversations" variant="ghost">My Conversations</UButton>
              </template>
            </div>
          </div>
          <div class="flex items-center gap-4">
            <UDropdown :items="userMenuItems">
              <UButton variant="ghost" trailing-icon="i-heroicons-chevron-down">
                <div class="flex items-center gap-2">
                  <UAvatar :alt="user?.name" size="sm" />
                  <span class="hidden md:inline">{{ user?.name }}</span>
                </div>
              </UButton>
            </UDropdown>
          </div>
        </div>
      </nav>
    </header>
    <main class="container mx-auto px-4 py-8">
      <slot />
    </main>
  </div>
</template>

<script setup lang="ts">
const { user, isSuperAdmin, isOrganizationOwner, isOrganizationSales, isGuest, clearAuth } = useAuth();

const handleLogout = async () => {
  try {
    await $fetch('/api/auth/logout', { method: 'POST' });
    clearAuth();
    navigateTo('/signin');
  } catch (error) {
    console.error('Logout error:', error);
  }
};

const userMenuItems = computed(() => [
  [
    {
      label: user.value?.email || '',
      disabled: true,
    },
  ],
  [
    {
      label: 'Profile',
      icon: 'i-heroicons-user',
      click: () => navigateTo('/profile'),
    },
    {
      label: 'Settings',
      icon: 'i-heroicons-cog',
      click: () => navigateTo('/settings'),
    },
  ],
  [
    {
      label: 'Logout',
      icon: 'i-heroicons-arrow-right-on-rectangle',
      click: handleLogout,
    },
  ],
]);
</script>
