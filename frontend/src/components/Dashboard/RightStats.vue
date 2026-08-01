<template>
  <aside class="right-stats">
    <div class="stat-card card-surface">
      <h3>Activity Balance</h3>
      <div class="placeholder-chart"></div>
    </div>
    <div class="stat-card card-surface">
      <h3>Weekly Progress</h3>
      <div class="placeholder-chart bar"></div>
    </div>
  </aside>
</template>

<script setup lang="ts"></script>

<style scoped>
.right-stats {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.stat-card {
  padding: 24px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.stat-card h3 {
  font-size: 18px;
  font-weight: 700;
  margin: 0 0 16px 0;
}

.placeholder-chart {
  flex: 1;
  min-height: 120px;
  background: var(--bg-primary);
  border-radius: var(--radius-lg);
  border: 1px dashed var(--border-default);
}
</style>
