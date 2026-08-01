<template>
  <div class="habits-toolbar card-surface">
    <div class="categories-nav">
      <button
        v-for="category in categories"
        :key="category"
        class="nav-tab"
        :class="{ active: currentCategory === category }"
        @click="$emit('update:currentCategory', category)"
      >
        {{ category }}
      </button>
    </div>
    <button class="btn btn-primary btn-sm" type="button" @click="$emit('toggle-create')">
      Create Habit
    </button>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  categories: string[]
  currentCategory: string
}>()

defineEmits<{
  'update:currentCategory': [value: string]
  'toggle-create': []
}>()
</script>

<style scoped>
.habits-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  width: 100%;
  box-sizing: border-box;
}

.categories-nav {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  flex: 1;
}

.nav-tab {
  background: transparent;
  border: none;
  color: var(--text-secondary);
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Hind Madurai', sans-serif;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s ease;
}

.nav-tab:hover {
  color: var(--text-primary);
}

.nav-tab.active {
  background: var(--accent-primary);
  color: #ffffff;
}

:global(html[data-theme='dark']) .nav-tab.active {
  color: #0f1115;
}
</style>
