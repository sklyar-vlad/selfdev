<template>
  <div class="heatmap-legend">
    <div class="l-cubes">
      <div class="cube" data-level="0"></div>
      <div class="cube" data-level="1"></div>
      <div class="cube" data-level="2"></div>
      <div class="cube" data-level="3"></div>
      <div class="cube" data-level="4"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
// No props or emits needed
</script>

<style scoped>
.heatmap-legend {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
  font-size: 11px;
  color: var(--text-secondary);
}

.l-cubes {
  display: flex;
  gap: 4px;
}

@media (max-width: 768px) {
  .heatmap-legend {
    justify-content: flex-start;
    margin-top: 10px;
  }
}
</style>
