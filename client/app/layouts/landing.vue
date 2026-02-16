<template>
  <UHeader :toggle="false">
    <template #left>
      <NuxtLink to="/">
        <div class="grid grid-cols-4 gap-2">
          <NuxtImg
            src="https://avatars.githubusercontent.com/u/114393211?v=4"
            alt="Profile Avatar"
            width="50"
            height="50"
            class="rounded-full"
          />
          <div class="col-span-3 items-center">
            <h2 class="text-lg text-primary-500">Dewa Surya Ariesta</h2>
            <p>Full Stack Developer</p>
          </div>
        </div>
      </NuxtLink>
    </template>

    <template #right>
      <UColorModeButton />

      <UButton
        to="https://github.com/DewaSRY/sociomile-app"
        target="_blank"
        icon="i-simple-icons-github"
        aria-label="GitHub"
        color="neutral"
        variant="ghost"
      />
    </template>
  </UHeader>

  <UMain>
    <slot />
  </UMain>

  <USeparator icon="i-simple-icons-nuxtdotjs" />

  <UFooter>
    <template #left>
      <p class="text-sm text-muted">
        Built with Nuxt UI • © {{ new Date().getFullYear() }}
      </p>
    </template>

    <template #right>
      <UButton
        to="https://github.com/DewaSRY/sociomile-app"
        target="_blank"
        icon="i-simple-icons-github"
        aria-label="GitHub"
        color="neutral"
        variant="ghost"
      />
    </template>
  </UFooter>
</template>
