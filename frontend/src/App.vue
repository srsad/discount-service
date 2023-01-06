<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <nav class="grid grid-cols-12 gap-4">
      <div class="col-span-9">logo</div>
      <div>
        <RouterLink class="navbar-link" to="/"> Home </RouterLink>
      </div>
      <div>
        <RouterLink class="navbar-link" to="/about"> About </RouterLink>
      </div>
      <div>
        <RouterLink class="navbar-link" to="/source"> Source </RouterLink>
      </div>
    </nav>
  </header>

  <RouterView />
</template>
