<script setup lang="ts">
import AppCreateForm from '@/components/sources/CreateForm.vue'
import AppSourceUpdateModal from '@/components/sources/modal/SourceUpdateModal.vue'
import AppSourceItems from '@/components/sources/SourceItemsList.vue'
</script>

<template>
  <div class="grid grid-cols-12">
    <div class="col-span-2">
      <AppCreateForm />
    </div>

    <div class="col-span-12 mt-8">
      <hr />
    </div>

    <div class="col-span-12">
      <AppSourceItems />
    </div>
  </div>

  <AppSourceUpdateModal />
</template>
