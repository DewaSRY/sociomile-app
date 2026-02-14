<template>
  <div class="min-h-screen bg-gray-50">
    <header class="bg-white shadow-sm">
      <nav class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <NuxtLink to="/" class="text-2xl font-bold text-primary">
            SocioMile
          </NuxtLink>
          <div class="flex items-center gap-4">
            <UButton to="/signin" variant="ghost">Sign In</UButton>
            <UButton to="/signup">Sign Up</UButton>
          </div>
        </div>
      </nav>
    </header>
    <main>
      <slot />
    </main>
    <footer class="bg-white border-t mt-auto">
      <div class="container mx-auto px-4 py-6 text-center text-gray-600">
        <p>&copy; 2026 SocioMile. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>
