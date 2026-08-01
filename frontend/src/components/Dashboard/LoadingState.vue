<template>
  <div class="loading-container card-surface">
    <div class="spinner"></div>
    <span class="loading-text">Loading dashboard...</span>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: clamp(40px, 8vw, 64px) 16px;
  gap: 20px;
  min-height: 200px;
}

.spinner {
  width: 48px;
  height: 48px;
  border: 3px solid var(--border-default);
  border-top-color: var(--accent-primary);
  border-radius: 50%;
  animation: spin 0.8s cubic-bezier(0.5, 0, 0.5, 1) infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.loading-text {
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.05em;
  text-transform: uppercase;
  text-align: center;
}
</style>
