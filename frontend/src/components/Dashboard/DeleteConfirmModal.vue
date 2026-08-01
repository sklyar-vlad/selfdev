<template>
  <div v-if="show" class="habit-modal-overlay" @click.self="$emit('cancel')">
    <div class="delete-modal card-surface">
      <h3>Delete "{{ habitName }}"?</h3>
      <p>This action cannot be undone.</p>
      <div class="habit-create-actions">
        <button class="btn btn-danger btn-sm" @click="$emit('confirm')">Delete</button>
        <button class="btn btn-sm" @click="$emit('cancel')">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  show: boolean
  habitName: string
}>()

defineEmits<{
  confirm: []
  cancel: []
}>()
</script>

<style scoped>
.habit-modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  background: rgba(0, 0, 0, 0.6);
}

.delete-modal {
  width: min(100%, 380px);
  padding: 24px;
  max-height: min(92vh, 420px);
  overflow: auto;
}

.delete-modal h3 {
  margin: 0 0 12px;
}
.delete-modal p {
  color: var(--text-secondary);
  margin-bottom: 20px;
}

.habit-create-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

@media (max-width: 768px) {
  .habit-modal-overlay {
    align-items: flex-end;
    padding: 12px 12px 0;
  }

  .delete-modal {
    width: 100%;
    max-width: 100%;
    padding: 18px 16px calc(16px + env(safe-area-inset-bottom));
    border-radius: 24px 24px 0 0;
  }

  .habit-create-actions {
    flex-direction: column;
    align-items: stretch;
  }

  .habit-create-actions .btn {
    width: 100%;
  }
}
</style>
