<script setup lang="ts">
import TheModal from '@/components/TheModal.vue'
import useSource from '@/composables/useSource'
const {
  showEditMobile,
  updateSourceItem,
  updateSource,
  clearUpdateSourceItem,
} = useSource()
</script>

<template>
  <TheModal
    v-model="showEditMobile"
    title="Редактировать источник"
    @closed="clearUpdateSourceItem"
  >
    <template #default="{ close }">
      <form @submit.prevent="updateSource(updateSourceItem)" class="grid gap-2">
        <label class="mb-4">
          <p class="text-sm font-semibold mb-2">Имя источника</p>
          <input
            v-model="updateSourceItem.name"
            type="text"
            class="form-control"
            required
          />
        </label>
        <div class="flex justify-center gap-2">
          <button class="btn btn-success">Обновить</button>
          <button @click="close" class="btn btn-danger">Отменить</button>
        </div>
      </form>
    </template>
  </TheModal>
</template>
