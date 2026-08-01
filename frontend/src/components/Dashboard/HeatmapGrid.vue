<template>
  <div class="heatmap-wrapper">
    <div class="days-labels"><span>Mon</span><span>Wed</span><span>Fri</span><span>Sun</span></div>
    <div class="cubes-scroll-container">
      <div class="cubes-grid">
        <div
          v-for="day in heatmap"
          :key="day.key"
          class="cube"
          :data-date="day.key"
          :data-level="day.level"
          :class="{ today: day.key === todayStr }"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface HeatmapDay {
  key: string
  level: number
}

defineProps<{
  heatmap: HeatmapDay[]
  todayStr: string
}>()
</script>

<style scoped>
.heatmap-wrapper {
  display: flex;
  gap: 16px;
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
</style>
