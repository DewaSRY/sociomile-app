<script setup lang="ts">
import HubOrgTable from "~/components/datatable/hub-org-table.vue";
import HubOrgForm from "~/components/forms/hub-org-forms.vue";
const tableRef = ref<InstanceType<typeof HubOrgTable> | null>(null);

const open = ref(false);


async function handleSubmit() {
  open.value = false;
  await tableRef.value?.refresh();
}

</script>

<template>
  <UContainer>
    <div class="mb-4 flex justify-between">
      <div>
        <h2 class="text-primary">Organization List</h2>
      </div>
      <div>
        <UModal v-model:open="open" title="Modal with title">
          <UButton label="New Organization" icon="i-lucide-plus" />
          <template #body>
            <HubOrgForm @on-submit="handleSubmit" />
          </template>
        </UModal>
      </div>
    </div>
    <HubOrgTable ref="tableRef" />
  </UContainer>
</template>
