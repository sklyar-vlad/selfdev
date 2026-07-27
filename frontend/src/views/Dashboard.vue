<template>
  <div class="main-layout">
    <DashboardHeader />
  </div>

  <div class="dashboard-container">
    <div class="dashboard-content">
      <!-- ЗАГРУЗКА -->
      <div v-if="loading" class="loading-container card-surface">
        <div class="spinner"></div>
        <span class="loading-text">Loading dashboard...</span>
      </div>

      <!-- ОШИБКА (теперь использует var(--error)) -->
      <div v-else-if="error" class="card-surface" style="padding: 16px 24px; color: var(--error)">
        {{ error }}
      </div>

      <!-- БЛОК ПРОФИЛЯ -->
      <section v-if="!loading && !error" class="user-profile-section card-surface">
        <div class="avatar-block">
          <img class="avatar-image" :src="defaultAvatar" alt="User avatar" />
          <div class="user-name-wrapper">
            <span class="user-label">User</span>
            <h1 class="username">{{ user?.username || 'SelfDev_Hero' }}</h1>
          </div>
        </div>

        <div class="xp-level-block">
          <div class="xp-info">
            <span class="lvl-text">Level {{ level }}</span>
            <span class="xp-count">{{ xpCurrent }} / {{ xpNext }} XP</span>
          </div>
          <div class="xp-bar-container">
            <div class="xp-bar-bg"></div>
            <div class="xp-bar-fill" :style="{ width: `${xpProgress}%` }"></div>
          </div>
        </div>

        <div class="quick-stats">
          <div class="stat-box">
            <span class="stat-n">{{ perfectDays }}</span>
            <span class="stat-t">Perfect Days</span>
          </div>
          <div class="stat-box streak">
            <span class="stat-n">{{ streakDays }}</span>
            <span class="stat-t">Days Streak</span>
          </div>
        </div>
      </section>

      <!-- НИЖНЯЯ ЧАСТЬ: ПРИВЫЧКИ И СТАТИСТИКА -->
      <div v-if="!loading && !error" class="main-grid">
        <!-- ЛЕВАЯ КОЛОНКА: ПРИВЫЧКИ С НАВИГАЦИЕЙ И АНИМАЦИЕЙ -->
        <section class="habits-container">
          <div class="habits-toolbar card-surface">
            <div class="categories-nav">
              <button
                v-for="category in categories"
                :key="category"
                class="nav-tab"
                :class="{ active: currentCategory === category }"
                @click="currentCategory = category"
              >
                {{ category }}
              </button>
            </div>
            <button
              class="btn btn-primary btn-sm"
              type="button"
              @click="showCreateHabitForm = !showCreateHabitForm"
            >
              Create Habit
            </button>
          </div>

          <div
            v-if="showCreateHabitForm"
            class="habit-modal-overlay"
            @click.self="showCreateHabitForm = false"
          >
            <form
              class="habit-create-form card-surface"
              @submit.prevent="editingHabitId ? updateHabit() : createHabit()"
            >
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
              <input
                v-if="newHabit.category === 'New'"
                v-model="newCategory"
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
                <button class="btn btn-sm" type="button" @click="showCreateHabitForm = false">
                  Cancel
                </button>
              </div>
            </form>
          </div>

          <div
            v-if="showDeleteModal"
            class="habit-modal-overlay"
            @click.self="showDeleteModal = false"
          >
            <div class="delete-modal card-surface">
              <h3>Delete "{{ habitToDelete?.name }}"?</h3>
              <p>This action cannot be undone.</p>
              <div class="habit-create-actions">
                <button class="btn btn-danger btn-sm" @click="deleteHabit">Delete</button>
                <button class="btn btn-sm" @click="showDeleteModal = false">Cancel</button>
              </div>
            </div>
          </div>

          <!-- Окно видимости с эффектом плавного затухания -->
          <div class="habits-fade-viewport">
            <div class="habits-scroll-window">
              <TransitionGroup name="habit-fade" tag="div" class="habits-wrapper-layout">
                <article
                  v-for="habit in filteredHabits"
                  :key="habit.id"
                  class="habit-card card-surface"
                  :style="{ '--habit-color': habit.color }"
                >
                  <div class="habit-header">
                    <div class="habit-title-group">
                      <div class="habit-title-row">
                        <h2>{{ habit.name }}</h2>

                        <!-- Кнопка ВВЕРХ -->
                        <button 
                          class="icon-btn reorder" 
                          title="Move Up" 
                          @click="moveHabit(habit, 'up')"
                          :disabled="filteredHabits.findIndex(h => h.id === habit.id) === 0"
                        >
                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="icon-svg">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 15.75l7.5-7.5 7.5 7.5" />
                          </svg>
                        </button>

                        <!-- Кнопка ВНИЗ -->
                        <button 
                          class="icon-btn reorder" 
                          title="Move Down" 
                          @click="moveHabit(habit, 'down')"
                          :disabled="filteredHabits.findIndex(h => h.id === habit.id) === filteredHabits.length - 1"
                        >
                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="icon-svg">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
                          </svg>
                        </button>

                        <!-- Кнопка редактирования -->
                        <button class="icon-btn" title="Edit" @click="openEditHabit(habit)">
                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon-svg">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0115.75 21H5.25A2.25 2.25 0 013 18.75V8.25A2.25 2.25 0 015.25 6H10" />
                          </svg>
                        </button>

                        <!-- Кнопка удаления -->
                        <button class="icon-btn delete" title="Delete" @click="openDeleteModal(habit)">
                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="icon-svg">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
                          </svg>
                        </button>
                      </div>
                      <span class="habit-status"> {{ habit.confirmedCount }}/365 cleared </span>
                    </div>
                    <button class="btn btn-primary btn-sm" @click="toggleHabit(habit)">
                      {{ isHabitDoneToday(habit) ? 'Cancel' : 'Done' }}
                    </button>
                  </div>

                  <!-- СЕТКА КУБИКОВ (HEATMAP НА 365 ДНЕЙ) -->
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
              </TransitionGroup>
            </div>
          </div>
        </section>

        <!-- ПРАВАЯ КОЛОНКА: ФИКСИРОВАННАЯ СТАТИСТИКА -->
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
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import defaultAvatar from '@/assets/default-avatar.jpg'
import { config } from '@/config/env'
import DashboardHeader from '@/components/Header/DashboardHeader.vue'

const todayStr = computed(() => toIsoDate(new Date()))
const updatingHabits = ref(new Set<string>())

interface User {
  user_id: string
  role: string
  username: string
  email: string
}

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

const currentCategory = ref('all')
const newCategory = ref('')
const user = ref<User | null>(null)
const habits = ref<Habit[]>([])
const loading = ref(true)
const error = ref('')
const showCreateHabitForm = ref(false)
const newHabit = ref({
  name: '',
  description: '',
  category: '',
  color: '#39d353',
  isGood: true,
})

const editingHabitId = ref<string | null>(null)
const showDeleteModal = ref(false)
const habitToDelete = ref<Habit | null>(null)

function openDeleteModal(habit: Habit) {
  habitToDelete.value = habit
  showDeleteModal.value = true
}

function openEditHabit(habit: Habit) {
  editingHabitId.value = habit.id
  newHabit.value = {
    name: habit.name,
    description: habit.description,
    category: habit.category,
    color: habit.color,
    isGood: habit.isGood,
  }
  newCategory.value = ''
  showCreateHabitForm.value = true
}

const categories = computed(() => [
  'all',
  ...new Set(habits.value.map((h) => h.category).filter(Boolean)),
])

const filteredHabits = computed(() => {
  let filtered =
    currentCategory.value === 'all'
      ? habits.value
      : habits.value.filter((h) => h.category === currentCategory.value)
  return filtered.sort((a, b) => a.order - b.order)
})

function toIsoDate(value: string | Date) {
  const date = typeof value === 'string' ? new Date(value) : value
  return new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()))
    .toISOString()
    .slice(0, 10)
}

function buildHeatmap(completedDates: string[]) {
  const completed = new Set(completedDates.map(toIsoDate))
  const days: HeatmapDay[] = []
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  const end = new Date(today)
  const day = end.getDay()
  const mondayOffset = day === 0 ? -6 : 1 - day
  end.setDate(end.getDate() + mondayOffset + 6)

  const start = new Date(end)
  start.setDate(start.getDate() - 52 * 7 + 1)
  const cursor = new Date(start)

  while (cursor <= end) {
    const key = toIsoDate(cursor)
    if (cursor <= today) {
      days.push({
        key,
        level: completed.has(key) ? 4 : 0,
      })
    }
    cursor.setDate(cursor.getDate() + 1)
  }
  return days
}

function isHabitDoneToday(habit: Habit) {
  const today = toIsoDate(new Date())
  return habit.confirmedDates.map(toIsoDate).includes(today)
}

async function fetchJson<T>(path: string, options: RequestInit = {}) {
  const response = await fetch(`${config.apiUrl}${path}`, {
    credentials: 'include',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  const text = await response.text()
  return text ? (JSON.parse(text) as T) : ({} as T)
}

async function fetchHabitDates(habitId: string) {
  const data = await fetchJson<{
    Dates?: Array<{
      HabitId: string
      Date: string
    }>
  }>(`/api/habit/${encodeURIComponent(habitId)}/confirm`)

  return (data.Dates || []).map((item: { HabitId: string; Date: string }) => item.Date)
}

async function refreshHabit(habitId: string) {
  const dates = await fetchHabitDates(habitId)
  applyHabitDates(habitId, dates)
}

function applyHabitDates(habitId: string, dates: string[]) {
  habits.value = habits.value.map((habit) =>
    habit.id === habitId
      ? {
          ...habit,
          confirmedDates: dates,
          confirmedCount: dates.length,
          heatmap: buildHeatmap(dates),
        }
      : habit,
  )
}

async function fetchHabits() {
  const data = await fetchJson<{
    habits?: Array<{
      habit_id?: string
      HabitId?: string
      name: string
      Name?: string
      description?: string
      Description?: string
      category?: string
      Category?: string
      color?: string
      Color?: string
      is_good?: boolean
      IsGood?: boolean
      order?: number
    }>
    Habits?: Array<{
      habit_id?: string
      HabitId?: string
      name: string
      Name?: string
      description?: string
      Description?: string
      category?: string
      Category?: string
      color?: string
      Color?: string
      is_good?: boolean
      IsGood?: boolean
      order?: number
    }>
  }>('/api/habits')

  const nextHabits = await Promise.all(
    (data.habits || data.Habits || []).map(async (habit, index) => {
      const id = habit.habit_id || habit.HabitId || ''
      const name = habit.name || habit.Name || ''
      const description = habit.description || habit.Description || ''
      const isGood = habit.is_good ?? habit.IsGood ?? false
      const category = habit.category || habit.Category || ''
      const color = habit.color || habit.Color || '#39d353'
      const confirmedDates = id ? await fetchHabitDates(id) : []

      return {
        id,
        name,
        description,
        isGood,
        color,
        category,
        confirmedDates,
        confirmedCount: confirmedDates.length,
        heatmap: buildHeatmap(confirmedDates),
        order: habit.order ?? index,
      }
    }),
  )

  habits.value = nextHabits
}

async function confirmHabit(habitId: string) {
  if (updatingHabits.value.has(habitId)) return
  updatingHabits.value.add(habitId)

  try {
    await fetchJson(`/api/habit/${encodeURIComponent(habitId)}/confirm`, {
      method: 'POST',
    })
    await refreshHabit(habitId)
  } finally {
    updatingHabits.value.delete(habitId)
  }
}

function closeHabitForm() {
  showCreateHabitForm.value = false
  editingHabitId.value = null
  newHabit.value = {
    name: '',
    description: '',
    category: '',
    color: '#39d353',
    isGood: true,
  }
  newCategory.value = ''
}

async function deleteHabit() {
  if (!habitToDelete.value) return
  await fetchJson(`/api/habit/${encodeURIComponent(habitToDelete.value.id)}`, {
    method: 'DELETE',
  })
  habits.value = habits.value.filter((h) => h.id !== habitToDelete.value?.id)
  showDeleteModal.value = false
  habitToDelete.value = null
}

async function cancelHabit(habitId: string) {
  await fetchJson(`/api/habit/${encodeURIComponent(habitId)}/confirm`, {
    method: 'DELETE',
  })
  await refreshHabit(habitId)
}

async function toggleHabit(habit: Habit) {
  if (isHabitDoneToday(habit)) {
    await cancelHabit(habit.id)
    return
  }
  await confirmHabit(habit.id)
}

async function createHabit() {
  await fetchJson('/api/habit', {
    method: 'POST',
    body: JSON.stringify({
      name: newHabit.value.name,
      description: newHabit.value.description,
      category:
        newHabit.value.category === 'New' ? newCategory.value.trim() : newHabit.value.category,
      color: newHabit.value.color,
      is_good: newHabit.value.isGood,
    }),
  })
  closeHabitForm()
  await fetchHabits()
}

async function updateHabit() {
  if (!editingHabitId.value) return
  await fetchJson(`/api/habit/${editingHabitId.value}`, {
    method: 'PUT',
    body: JSON.stringify({
      name: newHabit.value.name,
      description: newHabit.value.description,
      category:
        newHabit.value.category === 'New' ? newCategory.value.trim() : newHabit.value.category,
      color: newHabit.value.color,
      is_good: newHabit.value.isGood,
    }),
  })
  closeHabitForm()
  await fetchHabits()
}

const perfectDays = computed(() =>
  habits.value.reduce((sum, habit) => sum + habit.confirmedCount, 0),
)
const streakDays = computed(() => Math.max(...habits.value.map((habit) => habit.confirmedCount), 0))
const level = computed(() => Math.max(1, Math.floor(perfectDays.value / 10) + 1))
const xpCurrent = computed(() => perfectDays.value * 100)
const xpNext = computed(() => level.value * 1000)
const xpProgress = computed(() => Math.min(100, Math.round((xpCurrent.value / xpNext.value) * 100)))

function moveHabit(habit: Habit, direction: 'up' | 'down') {
  const visibleHabits = filteredHabits.value
  const visualIdx = visibleHabits.findIndex((h) => h.id === habit.id)
  if (visualIdx === -1) return

  const targetVisualIdx = direction === 'up' ? visualIdx - 1 : visualIdx + 1

  if (targetVisualIdx < 0 || targetVisualIdx >= visibleHabits.length) return

  const currentHabit = visibleHabits[visualIdx]
  const targetHabit = visibleHabits[targetVisualIdx]

  if (!currentHabit || !targetHabit) return

  const tempOrder = currentHabit.order
  currentHabit.order = targetHabit.order
  targetHabit.order = tempOrder

  habits.value = [...habits.value]
}

onMounted(async () => {
  loading.value = true
  error.value = ''
  try {
    await fetchHabits()
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to load dashboard'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
/* =========================================
   LOADING STATE
========================================= */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 64px 24px;
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
  to { transform: rotate(360deg); }
}

.loading-text {
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.05em;
  text-transform: uppercase;
}

/* =========================================
   LAYOUT & CONTAINERS
========================================= */
.dashboard-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  padding: 100px 24px 40px;
  box-sizing: border-box;
}

.dashboard-content {
  position: relative;
  z-index: 2;
  max-width: 1400px;
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: calc(100vh - 150px);
}

/* ЧИСТЫЕ КАРТОЧКИ, ПОЛНОСТЬЮ НА ПЕРЕМЕННЫХ */
.card-surface {
  background: var(--surface);
  border: 1px solid var(--surface-border);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
  color: var(--text-primary);
  transition: all 0.25s ease;
}

/* =========================================
   USER PROFILE SECTION
========================================= */
.user-profile-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px 32px;
  flex-shrink: 0;
}

.avatar-block {
  display: flex;
  align-items: center;
  gap: 20px;
}

.avatar-image {
  width: 80px;
  height: 80px;
  background: var(--bg-primary);
  border-radius: var(--radius-lg);
  border: 1px solid var(--surface-border);
}

.user-label {
  font-size: 13px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.username {
  font-size: 28px;
  font-weight: 800;
  margin: 4px 0 0 0;
  color: var(--accent-primary);
}

.xp-level-block {
  width: 360px;
}

.xp-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 14px;
  font-weight: 700;
}

.lvl-text { color: var(--text-primary); }
.xp-count { color: var(--text-secondary); }

.xp-bar-container {
  position: relative;
  height: 12px;
}

.xp-bar-bg {
  position: absolute;
  width: 100%;
  height: 100%;
  background: var(--border-default);
  border-radius: 6px;
}

.xp-bar-fill {
  position: absolute;
  height: 100%;
  background: linear-gradient(90deg, var(--accent-primary), var(--accent-dark));
  border-radius: 6px;
  box-shadow: 0 0 12px var(--accent-light);
}

.quick-stats {
  display: flex;
  gap: 32px;
}

.stat-box {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-n {
  font-size: 32px;
  font-weight: 800;
  color: var(--text-primary);
}

.stat-t {
  font-size: 11px;
  color: var(--text-secondary);
  text-transform: uppercase;
  margin-top: 2px;
}

.streak .stat-n {
  color: var(--accent-primary);
}

/* =========================================
   MAIN GRID & HABITS NAVIGATION
========================================= */
.main-grid {
  display: grid;
  grid-template-columns: 1fr 380px;
  gap: 24px;
  flex: 1;
  overflow: visible;
  min-height: 0;
  padding-bottom: 24px;
}

.habits-container {
  overflow: visible;
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-height: 0;
}

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

/* =========================================
   MODALS & FORMS
========================================= */
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
}

.habit-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-secondary);
  font-size: 14px;
}

.color-picker {
  display: flex;
  align-items: center;
  gap: 10px;
  color: var(--text-secondary);
  font-size: 14px;
}

.color-picker input[type='color'] {
  width: 20px;
  height: 20px;
  padding: 0;
  border: none;
  border-radius: 50%;
  overflow: hidden;
  cursor: pointer;
  background: transparent;
}

.color-picker input[type='color']::-webkit-color-swatch-wrapper { padding: 0; }
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

.delete-modal {
  width: min(100%, 380px);
  padding: 24px;
}

.delete-modal h3 { margin: 0 0 12px; }
.delete-modal p {
  color: var(--text-secondary);
  margin-bottom: 20px;
}

/* =========================================
   ICON BUTTONS (Unified & Clean)
========================================= */
.habit-title-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
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

.icon-btn.reorder + .icon-btn {
  margin-left: 6px;
  border-left: 1px solid var(--border-default);
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

/* =========================================
   HABIT CARD & HEATMAP
========================================= */
.habit-card {
  padding: 20px;
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
}

.habit-header h2 {
  font-size: 24px;
  font-weight: 700;
  margin: 0;
}

.habit-status {
  font-size: 13px;
  color: var(--text-secondary);
  display: block;
  margin-top: 4px;
}

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

.cube[data-level='0'] { background: var(--border-default); }
.cube[data-level='1'] { background: color-mix(in srgb, var(--habit-color) 25%, transparent); }
.cube[data-level='2'] { background: color-mix(in srgb, var(--habit-color) 50%, transparent); }
.cube[data-level='3'] { background: color-mix(in srgb, var(--habit-color) 75%, transparent); }
.cube[data-level='4'] { background: var(--habit-color); }

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

/* =========================================
   FADE VIEWPORT & ANIMATIONS
========================================= */
.habits-fade-viewport {
  flex: 1;
  min-height: 0;
  position: relative;
  overflow: hidden;
  -webkit-mask-image: linear-gradient(
    to bottom,
    transparent 0%,
    black 32px,
    black calc(100% - 32px),
    transparent 100%
  );
  mask-image: linear-gradient(
    to bottom,
    transparent 0%,
    black 32px,
    black calc(100% - 32px),
    transparent 100%
  );
}

.habits-scroll-window {
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 40px 36px 40px 36px;
  box-sizing: border-box;
}

.habits-scroll-window::-webkit-scrollbar {
  width: 6px;
}
.habits-scroll-window::-webkit-scrollbar-thumb {
  background: var(--border-default);
  border-radius: 10px;
}
.habits-scroll-window::-webkit-scrollbar-thumb:hover {
  background: var(--text-tertiary);
}

.habits-wrapper-layout {
  display: flex;
  flex-direction: column;
  gap: 20px;
  position: relative;
}

.habit-fade-enter-active,
.habit-fade-leave-active {
  transition: all 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

.habit-fade-enter-from {
  opacity: 0;
  transform: translateY(24px) scale(0.98);
}

.habit-fade-leave-to {
  opacity: 0;
  transform: translateY(-24px) scale(0.96);
}

.habit-fade-leave-active {
  position: absolute;
  left: 0;
  right: 0;
  z-index: 0;
  pointer-events: none;
}

.habit-fade-move {
  transition: transform 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

/* =========================================
   RIGHT SIDEBAR & CHARTS
========================================= */
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

/* =========================================
   BUTTONS
========================================= */
.btn {
  padding: 10px 20px;
  border-radius: var(--radius-md);
  font-weight: 700;
  font-family: 'Hind Madurai', sans-serif;
  cursor: pointer;
  transition: 0.25s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
}

.btn-sm {
  padding: 8px 18px;
  font-size: 14px;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-dark));
  color: #ffffff;
  box-shadow: 0 4px 15px color-mix(in srgb, var(--accent-primary) 30%, transparent);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px color-mix(in srgb, var(--accent-primary) 40%, transparent);
}

.btn-danger {
  background: var(--error);
  color: #ffffff;
}

.btn-danger:hover {
  background: color-mix(in srgb, var(--error) 85%, black);
  transform: translateY(-2px);
}

/* =========================================
   RESPONSIVE
========================================= */
@media (max-width: 1024px) {
  .dashboard-content {
    height: auto;
    overflow: visible;
  }
  .main-grid {
    grid-template-columns: 1fr;
  }
  .habits-scroll-window {
    overflow-y: visible;
    height: auto;
  }
  .user-profile-section {
    flex-direction: column;
    gap: 20px;
    align-items: flex-start;
  }
  .xp-level-block {
    width: 100%;
  }
}
</style>