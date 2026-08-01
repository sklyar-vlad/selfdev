<template>
  <div class="card-surface" style="padding: 16px 24px; color: var(--error)">
    {{ message }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  message: string
}>()
</script>
