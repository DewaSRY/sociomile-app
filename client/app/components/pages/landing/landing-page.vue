<template>
  <div class="min-h-screen">
    <section class="py-20 px-4">
      <div class="container mx-auto text-center">
        <h1 class="text-5xl font-bold mb-6">Welcome to SocioMile</h1>
        <p class="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
          A multi-tenant customer messaging platform that helps businesses track and manage customer conversations across multiple channels.
        </p>
        <div class="flex items-center justify-center gap-4">
          <UButton to="/signup" size="xl">Get Started</UButton>
          <UButton to="/signin" variant="outline" size="xl">Sign In</UButton>
        </div>
      </div>
    </section>

    <section class="py-16 px-4 bg-white">
      <div class="container mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">Key Features</h2>
        <div class="grid md:grid-cols-3 gap-8">
          <UCard>
            <template #header>
              <div class="text-4xl mb-2">💬</div>
              <h3 class="text-xl font-semibold">Conversation Management</h3>
            </template>
            <p class="text-gray-600">Centralize all customer conversations from multiple channels in one place.</p>
          </UCard>

          <UCard>
            <template #header>
              <div class="text-4xl mb-2">🎫</div>
              <h3 class="text-xl font-semibold">Ticket System</h3>
            </template>
            <p class="text-gray-600">Convert conversations into tickets for better issue tracking and resolution.</p>
          </UCard>

          <UCard>
            <template #header>
              <div class="text-4xl mb-2">👥</div>
              <h3 class="text-xl font-semibold">Team Collaboration</h3>
            </template>
            <p class="text-gray-600">Assign conversations to team members and collaborate efficiently.</p>
          </UCard>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// Landing page
</script>
