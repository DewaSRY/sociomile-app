<script setup lang="ts">
import { useRefresh } from "~/composables/auth/useRefresh";

const { fetchRefresh, profile } = useRefresh();

async function handleNavigation() {
  if (!profile.value) return;

  const role = profile.value.roleName;

  if (role === "super_admin") {
    await navigateTo("/hub/dashboard");
  } else if (role === "organization_owner" || role === "organization_sales") {
    await navigateTo("/organization/dashboard");
  } else {
    await navigateTo("/guest/dashboard");
  }
}

await fetchRefresh();
</script>

<template>
  <UButton
    v-if="profile"
    type="submit"
    variant="link"
    size="xs"
    @click="handleNavigation"
  >
    <slot>
      <div>
        <span> sing in as </span>
        <span>
          {{ profile.name }}
        </span>
      </div>
    </slot>
  </UButton>
</template>
