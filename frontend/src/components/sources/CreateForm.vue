<script setup lang="ts">
import type { Ref } from 'vue'
import { ref } from 'vue'
import { useToast } from 'vue-toastification'

import api from '@/api/index'
import useSource from '@/composables/useSource'
import type { SourceRequest } from '@/types/sourse'
import { SourceTypesEnum } from '@/types/sourse'

const { getSourcesList } = useSource()

const toast = useToast()

const form: Ref<SourceRequest> = ref({
  name: '',
  type: SourceTypesEnum.STORE,
})

async function onCreate() {
  try {
    const result = await api.source.create(form.value)
    form.value.name = ''
    toast.success(result.data.message)
    getSourcesList()
  } catch (error: any) {
    console.error('Не удалось создать источник', error)
    toast.error(error.data.message)
  }
}
</script>

<template>
  <div>
    <h3 class="mb-2">Добавить источник</h3>
    <form @submit.prevent="onCreate" class="flex gap-2">
      <input
        v-model.trim="form.name"
        type="text"
        class="form-control"
        placeholder="Источник"
        required
      />
      <button type="submit" class="btn btn-primary">Создать</button>
    </form>
  </div>
</template>
