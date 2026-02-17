<script setup lang="ts">
import type { NavigationMenuItem } from "@nuxt/ui";
import AuthorProfile from "~/components/ui/author-profile.vue";
import UserProfile from "~/components/ui/user-profile.vue";
import { useProfile } from "~/composables/auth/useProfile";

const { profile, fetchProfile, reset } = useProfile();
const open = ref(false);
const links = [
  [
    {
      label: "Home",
      icon: "i-lucide-house",
      to: "/guest/dashboard",
      onSelect: () => {
        open.value = false;
      },
    },
  ],
] satisfies NavigationMenuItem[][];

const groups = computed(() => [
  {
    id: "links",
    label: "Go to",
    items: links.flat(),
  },
]);

onUnmounted(() => {
  reset();
});

await fetchProfile();
</script>

<template>
  <UDashboardGroup unit="rem" class="h-screen" orientation="horizontal">
    <UDashboardSidebar
      id="default"
      v-model:open="open"
      collapsible
      resizable
      class="bg-elevated/25"
      :ui="{ footer: 'lg:border-t lg:border-default' }"
    >
      <template #header="{ collapsed }">
        <UserProfile
          v-if="profile"
          :name="profile.name"
          :user-role="profile.roleName"
          :organization-name="profile.organization?.name"
          :collapsed="collapsed"
        />
      </template>

      <template #default="{ collapsed }">
        <UDashboardSearchButton
          :collapsed="collapsed"
          class="bg-transparent ring-default"
        />

        <UNavigationMenu
          :collapsed="collapsed"
          :items="links[0]"
          orientation="vertical"
          tooltip
          popover
        />

        <UNavigationMenu
          :collapsed="collapsed"
          :items="links[1]"
          orientation="vertical"
          tooltip
          class="mt-auto"
        />
      </template>

      <template #footer>
        <AuthorProfile />
      </template>
    </UDashboardSidebar>
    <UDashboardSearch :groups="groups" />
    <slot />
  </UDashboardGroup>
</template>
