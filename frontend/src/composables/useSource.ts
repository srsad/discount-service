import type { Ref } from 'vue'
import { onMounted, readonly, ref } from 'vue'
// import { $vfm } from 'vue-final-modal'
import { useToast } from 'vue-toastification'

import api from '@/api/index'
import type { Source } from '@/types/sourse'

const sources: Ref<Source[]> = ref([])

const showEditMobile: Ref<boolean> = ref(false)
const updateSourceItem: Ref<Source> = ref({
  id: null,
  name: '',
})

export default function useSource() {
  const toast = useToast()

  onMounted(() => {
    getSourcesList()
  })

  /**
   * Список источников
   */
  async function getSourcesList() {
    try {
      const result = await api.source.getAll()
      sources.value = result.data.data
    } catch (error: any) {
      console.error('Не получить список источников', error)
      toast.error(error.data.message)
    }
  }

  /**
   * Удалить источник
   */
  async function removeSource(sourceId: number) {
    try {
      const result = await api.source.remove(sourceId)

      const itemIndex = sources.value.findIndex(
        (el: Source) => el.id === sourceId
      )
      sources.value.splice(itemIndex, 1)

      toast.success(result.data.message)
    } catch (error: any) {
      console.error('Не удалось удалить источник', error)
      toast.error(error.data.message)
    }
  }

  /**
   * Обновить источник
   */
  async function updateSource(source: Source) {
    if (!source) {
      return
    }

    try {
      const result = await api.source.update(source)

      const itemIndex = sources.value.findIndex(
        (el: Source) => el.id === source.id
      )
      sources.value.splice(itemIndex, 1, result.data.data)

      toast.success(result.data.message)
      clearUpdateSourceItem()
    } catch (error: any) {
      console.error('Не удалось обновить источник', error)
      toast.error(error.data.message)
    }
  }

  /**
   * Открыть модальное окно для редактирования источника
   */
  function openModalForEditSource(source: Source) {
    updateSourceItem.value = JSON.parse(JSON.stringify(source))
    showEditMobile.value = true
  }

  function clearUpdateSourceItem() {
    updateSourceItem.value = {
      id: null,
      name: '',
    }
    showEditMobile.value = false
  }

  return {
    sources: readonly(sources),
    showEditMobile: showEditMobile,
    updateSourceItem,

    openModalForEditSource,
    updateSource,
    clearUpdateSourceItem,

    getSourcesList,
    removeSource,
  }
}
