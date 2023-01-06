<script setup lang="ts">
import type { PropType, Ref } from 'vue'
import { ref } from 'vue'

const props = defineProps({
  value: {
    type: Boolean as PropType<boolean>,
    default: false,
  },
  title: {
    type: String as PropType<string>,
    default: '',
  },
})

const showModal: Ref<boolean> = ref(props.value)
</script>

<template>
  <vue-final-modal
    v-model="showModal"
    v-bind="$attrs"
    :max-width="200"
    classes="flex justify-center items-center"
    content-class="relative flex flex-col max-h-full mx-4 p-4 border dark:border-gray-800 rounded bg-white dark:bg-gray-900"
    @closed="$emit('closed')"
  >
    <template #default="{ close }">
      <div>
        <h3 v-if="title" class="text-2xl mb-4">
          {{ title }}
        </h3>
        <button class="absolute top-0 right-2" @click="close">x</button>
      </div>
      <slot v-bind="{ close }" />
    </template>
  </vue-final-modal>
</template>
