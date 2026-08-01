<template>
  <div class="dashboard-error card-surface">
    {{ message }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  message: string
}>()
</script>

<style scoped>
.dashboard-error {
  width: 100%;
  padding: 16px 20px;
  color: var(--error);
  word-break: break-word;
}

@media (max-width: 768px) {
  .dashboard-error {
    padding: 14px 16px;
  }
}
</style>
