<template>
  <article class="habit-card card-surface" :style="{ '--habit-color': habit.color }">
    <div class="habit-header">
      <div class="habit-title-group">
        <div class="habit-title-row">
          <h2>{{ habit.name }}</h2>

          <button
            class="icon-btn reorder"
            title="Move Up"
            @click="$emit('move-up')"
            :disabled="!canMoveUp"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="icon-svg"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.75l7.5-7.5 7.5 7.5" />
            </svg>
          </button>

          <button
            class="icon-btn reorder"
            title="Move Down"
            @click="$emit('move-down')"
            :disabled="!canMoveDown"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="icon-svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M19.5 8.25l-7.5 7.5-7.5-7.5"
              />
            </svg>
          </button>

          <button class="icon-btn" title="Edit" @click="$emit('edit')">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="icon-svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10"
              />
            </svg>
          </button>

          <button class="icon-btn delete" title="Delete" @click="$emit('delete')">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="icon-svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
              />
            </svg>
          </button>
        </div>
        <span class="habit-status"> {{ habit.confirmedCount }}/365 cleared </span>
      </div>
      <button class="btn btn-primary btn-sm" @click="$emit('toggle')" :disabled="isUpdating">
        {{ isDoneToday ? 'Cancel' : 'Done' }}
      </button>
    </div>

    <div class="heatmap-wrapper">
      <div class="days-labels">
        <span>Mon</span><span>Wed</span><span>Fri</span><span>Sun</span>
      </div>
      <div class="cubes-scroll-container">
        <div class="cubes-grid">
          <div
            v-for="day in habit.heatmap"
            :key="day.key"
            class="cube"
            :data-date="day.key"
            :data-level="day.level"
            :class="{ today: day.key === todayStr }"
          ></div>
        </div>
      </div>
    </div>

    <div class="heatmap-legend">
      <div class="l-cubes">
        <div class="cube" data-level="0"></div>
        <div class="cube" data-level="1"></div>
        <div class="cube" data-level="2"></div>
        <div class="cube" data-level="3"></div>
        <div class="cube" data-level="4"></div>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
interface HeatmapDay {
  key: string
  level: number
}

interface Habit {
  id: string
  name: string
  description: string
  isGood: boolean
  color: string
  category: string
  confirmedDates: string[]
  confirmedCount: number
  heatmap: HeatmapDay[]
  order: number
}

defineProps<{
  habit: Habit
  todayStr: string
  canMoveUp: boolean
  canMoveDown: boolean
  isDoneToday: boolean
  isUpdating: boolean
}>()

defineEmits<{
  toggle: []
  edit: []
  delete: []
  'move-up': []
  'move-down': []
}>()
</script>

<style scoped>
.habit-card {
  padding: clamp(16px, 3vw, 20px);
  position: relative;
  transition: all 0.25s ease;
  width: 100%;
  box-sizing: border-box;
}

.habit-card:hover {
  border-color: var(--accent-primary);
  box-shadow: 0 8px 30px color-mix(in srgb, var(--text-primary) 8%, transparent);
}

.habit-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  gap: 16px;
}

.habit-header h2 {
  font-size: clamp(20px, 3vw, 24px);
  font-weight: 700;
  margin: 0;
  min-width: 0;
}

.habit-title-row {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  padding: 0;
  border: 1px solid var(--border-default);
  background: transparent;
  border-radius: var(--radius-sm);
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s ease;
  flex-shrink: 0;
}

.icon-svg {
  width: 16px;
  height: 16px;
  stroke: currentColor;
  display: block;
}

.icon-btn:hover:not(:disabled) {
  color: var(--accent-primary);
  border-color: var(--accent-primary);
  background: color-mix(in srgb, var(--accent-primary) 10%, transparent);
}

.icon-btn.delete:hover:not(:disabled) {
  color: var(--error);
  border-color: var(--error);
  background: color-mix(in srgb, var(--error) 10%, transparent);
}

.icon-btn.reorder {
  opacity: 0.5;
}

.icon-btn.reorder:hover:not(:disabled) {
  opacity: 1;
}

.icon-btn.reorder:disabled {
  cursor: not-allowed;
  opacity: 0.15;
  border-color: transparent;
  background: transparent;
}

/* ВРЕДНЫЙ СЕЛЕКТОР УДАЛЕН ОТСЮДА */

.habit-status {
  font-size: 13px;
  color: var(--text-secondary);
  display: block;
  margin-top: 4px;
}

.heatmap-wrapper {
  display: flex;
  gap: 12px;
  background: var(--bg-primary);
  padding: 12px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-default);
  width: 100%;
  box-sizing: border-box;
}

.days-labels {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  font-size: 10px;
  color: var(--text-secondary);
  padding: 2px 0;
  flex-shrink: 0;
}

.cubes-scroll-container {
  flex: 1;
  overflow: hidden;
  width: 100%;
}

.cubes-scroll-container::-webkit-scrollbar {
  height: 4px;
}
.cubes-scroll-container::-webkit-scrollbar-thumb {
  background: var(--border-default);
  border-radius: 4px;
}

.cubes-grid {
  display: grid;
  grid-template-rows: repeat(7, minmax(0, 1fr));
  grid-auto-flow: column;
  grid-auto-columns: minmax(0, 1fr);
  gap: 2px;
  width: 100%;
}

.cube {
  width: 100%;
  aspect-ratio: 1;
  background: var(--border-default);
  border-radius: 3px;
}

.cube.today {
  box-shadow: inset 0 0 0 1.5px var(--accent-primary);
  position: relative;
  z-index: 1;
}

.cube[data-level='0'] {
  background: var(--border-default);
}
.cube[data-level='1'] {
  background: color-mix(in srgb, var(--habit-color) 25%, transparent);
}
.cube[data-level='2'] {
  background: color-mix(in srgb, var(--habit-color) 50%, transparent);
}
.cube[data-level='3'] {
  background: color-mix(in srgb, var(--habit-color) 75%, transparent);
}
.cube[data-level='4'] {
  background: var(--habit-color);
}

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
  .habit-header {
    flex-direction: column;
    align-items: stretch;
    margin-bottom: 16px;
  }

  .habit-title-row {
    align-items: flex-start;
  }

  .habit-header .btn {
    width: 100%;
  }

  .icon-btn {
    width: 44px;
    height: 44px;
  }

  .heatmap-wrapper {
    gap: 8px;
    padding: 10px;
  }

  .days-labels {
    font-size: 9px;
  }

  .cubes-scroll-container {
    overflow: hidden;
  }

  .cubes-grid {
    gap: 1px;
  }

  .cube {
    width: 100%;
    min-width: 0;
    border-radius: 2px;
  }

  .heatmap-legend {
    justify-content: flex-start;
  }
}
</style>