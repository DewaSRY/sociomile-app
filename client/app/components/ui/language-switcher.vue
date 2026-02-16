<template>
  <v-menu>
    <template v-slot:activator="{ props }">
      <v-btn icon v-bind="props" class="ml-2">
        <v-icon>mdi-translate</v-icon>
      </v-btn>
    </template>
    <v-list>
      <v-list-item @click="switchLocale('en')">
        <v-list-item-title>English</v-list-item-title>
      </v-list-item>
      <v-list-item @click="switchLocale('id')">
        <v-list-item-title>Indonesia</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script setup lang="ts">
import type { LocaleType } from "$types/common";
const { locale } = useI18n();
const switchLocale = (newLocale: LocaleType) => {
  locale.value = newLocale;
};
</script>
