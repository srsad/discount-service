<script setup lang="ts">
import useSource from '@/composables/useSource'

const { sources, removeSource, openModalForEditSource } = useSource()
</script>

<template>
  <div>
    <h3 class="mt-4 mb-2">Список источников</h3>
  </div>
  <div>
    <ul class="flex flex-col gap-2">
      <li class="flex gap-2" v-for="item of sources" :key="item.id">
        <button class="btn btn-xs rounded-sm" @click="removeSource(item.id)">
          x
        </button>
        <button
          class="btn btn-xs rounded-sm"
          @click="openModalForEditSource(item)"
        >
          E
        </button>
        {{ item.name }}
      </li>
    </ul>
  </div>
</template>
