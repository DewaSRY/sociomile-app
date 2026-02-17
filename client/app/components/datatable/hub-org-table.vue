<script setup lang="ts">
import { useOrganizations } from "~/composables/hub/useGetOrganization";
import type { TableColumn } from "@nuxt/ui";
import { getPaginationRowModel } from "@tanstack/table-core";

const toast = useToast();
const table = useTemplateRef("table");

const { fetchOrganizations, organizations } = useOrganizations();

const tableData = computed(() => organizations.value?.data ?? []);

const pagination = ref({
  pageIndex: 0,
  pageSize: 10,
});

const columns: TableColumn<HubOrganizationResponse>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "ownerName",
    header: "Owner Name",
  },
  {
    accessorKey: "createdAt",
    header: "Created At",
  },
  {
    id: "actions",
  },
];


async function refresh() {
  await fetchOrganizations()
}


await fetchOrganizations();

defineExpose({
  refresh
})

</script>

<template>
  <UTable
    v-model:pagination="pagination"
    :pagination-options="{
      getPaginationRowModel: getPaginationRowModel(),
    }"
    :data="tableData"
    :columns="columns"
    :ui="{
      thead: '',
      th: 'px-6 py-4 text-xs font-semibold uppercase tracking-wider text-gray-500 dark:text-gray-400',
      td: 'px-6 py-4',
      tr: 'hover:bg-gray-50/60 dark:hover:bg-gray-600/60 transition-colors',
    }"
  >
    <template #name-data="{ row }">
      <div class="flex flex-col">
        <span class="font-semibold text-gray-900 dark:text-white">
          {{ row.original.name }}
        </span>
        <span class="text-sm text-gray-500">
          @{{ row.original.name.toLowerCase().replace(/\s+/g, "") }}
        </span>
      </div>
    </template>

    <template #ownerName-data="{ row }">
      <div class="flex items-center gap-3">
        <div>
          <p class="font-medium text-highlighted">
            {{ row.original.ownerName }}
          </p>
          <p>@{{ row.original.ownerName }}</p>
        </div>
      </div>
    </template>

    <template #createdAt-data="{ row }">
      <span class="text-sm text-gray-600 dark:text-gray-400">
        {{
          new Date(row.original.createdAt).toLocaleDateString("id-ID", {
            day: "2-digit",
            month: "short",
            year: "numeric",
          })
        }}
      </span>
    </template>

    <template #actions-data="{ row }">
      <!-- <UButton size="xs" @click="edit(row.original)"> Edit </UButton> -->
    </template>
  </UTable>
</template>
