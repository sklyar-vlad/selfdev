<template>
  <div v-if="show" class="habit-modal-overlay" @click.self="$emit('close')">
    <form class="habit-create-form card-surface" @submit.prevent="$emit('save')">
      <h3>{{ editingHabitId ? 'Update Habit' : 'Create Habit' }}</h3>
      <input
        v-model="newHabit.name"
        class="habit-input"
        type="text"
        placeholder="Habit name"
        required
      />
      <input
        v-model="newHabit.description"
        class="habit-input"
        type="text"
        placeholder="Description"
      />
      <select v-model="newHabit.category" class="habit-input">
        <option disabled value="">Choose category</option>
        <option
          v-for="category in categories.filter((c) => c !== 'all')"
          :key="category"
          :value="category"
        >
          {{ category }}
        </option>
        <option value="New">+ New category</option>
      </select>

      <!-- ИСПРАВЛЕНО: v-model заменен на :value и @input -->
      <input
        v-if="newHabit.category === 'New'"
        :value="newCategory"
        @input="$emit('update:newCategory', ($event.target as HTMLInputElement).value)"
        class="habit-input"
        placeholder="New category"
      />

      <div class="habit-options">
        <label class="habit-checkbox">
          <input v-model="newHabit.isGood" type="checkbox" />
          <span>Good habit</span>
        </label>
        <label class="color-picker">
          <span>Color</span>
          <input v-model="newHabit.color" type="color" />
        </label>
      </div>
      <div class="habit-create-actions">
        <button class="btn btn-primary btn-sm" type="submit">
          {{ editingHabitId ? 'Update' : 'Save' }}
        </button>
        <button class="btn btn-sm" type="button" @click="$emit('close')">Cancel</button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
interface HabitForm {
  name: string
  description: string
  category: string
  color: string
  isGood: boolean
}

defineProps<{
  show: boolean
  editingHabitId: string | null
  newHabit: HabitForm
  newCategory: string
  categories: string[]
}>()

// ИСПРАВЛЕНО: добавлен эммит для обновления newCategory
defineEmits<{
  close: []
  save: []
  'update:newCategory': [value: string]
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

.habit-create-form {
  width: min(100%, 420px);
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 14px;
  max-height: min(92vh, 780px);
  overflow: auto;
}

.habit-create-form h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
}

.habit-input,
.habit-create-form select {
  width: 100%;
  box-sizing: border-box;
  padding: 12px 14px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-default);
  background: var(--bg-primary);
  color: var(--text-primary);
  font-size: 14px;
  font-family: 'Hind Madurai', sans-serif;
  transition: 0.2s;
}

.habit-create-form select {
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  cursor: pointer;
  background-image:
    linear-gradient(45deg, transparent 50%, var(--text-secondary) 50%),
    linear-gradient(135deg, var(--text-secondary) 50%, transparent 50%);
  background-position:
    calc(100% - 18px) calc(50% - 3px),
    calc(100% - 12px) calc(50% - 3px);
  background-size: 6px 6px;
  background-repeat: no-repeat;
  padding-right: 40px;
}

.habit-create-form select:focus,
.habit-input:focus {
  outline: none;
  border-color: var(--accent-primary);
  box-shadow: 0 0 0 3px color-mix(in srgb, var(--accent-primary) 20%, transparent);
}

.habit-create-form select option {
  background: var(--bg-primary);
  color: var(--text-primary);
}

.habit-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}

.habit-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-secondary);
  font-size: 14px;
  min-height: 44px;
}

.color-picker {
  display: flex;
  align-items: center;
  gap: 10px;
  color: var(--text-secondary);
  font-size: 14px;
  min-height: 44px;
}

.color-picker input[type='color'] {
  width: 44px;
  height: 44px;
  padding: 0;
  border: none;
  border-radius: 50%;
  overflow: hidden;
  cursor: pointer;
  background: transparent;
}

.color-picker input[type='color']::-webkit-color-swatch-wrapper {
  padding: 0;
}
.color-picker input[type='color']::-webkit-color-swatch {
  border: 2px solid var(--border-default);
  border-radius: 50%;
}
.color-picker input[type='color']::-moz-color-swatch {
  border: 2px solid var(--border-default);
  border-radius: 50%;
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

  .habit-create-form {
    width: 100%;
    max-width: 100%;
    max-height: calc(100vh - 12px);
    padding: 18px 16px calc(16px + env(safe-area-inset-bottom));
    border-radius: 24px 24px 0 0;
  }

  .habit-options,
  .habit-create-actions {
    flex-direction: column;
    align-items: stretch;
  }

  .habit-create-actions .btn {
    width: 100%;
  }
}
</style>
