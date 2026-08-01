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
</style>
