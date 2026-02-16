<template>
  <div>
    <v-breadcrumbs :items="navList">
      <template v-slot:divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </template>
    </v-breadcrumbs>
  </div>

  <slot />
</template>

<script setup lang="ts">
import type { BreadcrumbItem } from "$types/common";

interface Props {
  navList?: BreadcrumbItem[];
}

const props = withDefaults(defineProps<Props>(), {
  navList: () => [],
});
</script>
