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
  gap: clamp(14px, 2vw, 24px);
}

.stat-card {
  padding: clamp(16px, 3vw, 24px);
  flex: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
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

@media (max-width: 768px) {
  .right-stats {
    gap: 12px;
  }
}
</style>
