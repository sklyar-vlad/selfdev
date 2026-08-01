<template>
  <div class="main-layout">
    <DashboardHeader />
  </div>

  <div class="dashboard-container">
    <div class="dashboard-content">
      <!-- ЗАГРУЗКА -->
      <LoadingState v-if="loading" />

      <!-- ОШИБКА (теперь использует var(--error)) -->
      <DashboardError v-else-if="error" :message="error" />

      <!-- БЛОК ПРОФИЛЯ -->
      <UserProfile
        v-if="!loading && !error"
        :user="user"
        :level="level"
        :xpCurrent="xpCurrent"
        :xpNext="xpNext"
        :xpProgress="xpProgress"
        :perfectDays="perfectDays"
        :streakDays="streakDays"
        :defaultAvatar="defaultAvatar"
      />

      <!-- НИЖНЯЯ ЧАСТЬ: ПРИВЫЧКИ И СТАТИСТИКА -->
      <div v-if="!loading && !error" class="main-grid">
        <!-- ЛЕВАЯ КОЛОНКА: ПРИВЫЧКИ С НАВИГАЦИЕЙ И АНИМАЦИЕЙ -->
        <section class="habits-container">
          <HabitsToolbar
            :categories="categories"
            :currentCategory="currentCategory"
            @update:currentCategory="currentCategory = $event"
            @toggle-create="showCreateHabitForm = !showCreateHabitForm"
          />

          <HabitCreateModal
            :show="showCreateHabitForm"
            :editingHabitId="editingHabitId"
            :newHabit="newHabit"
            v-model:newCategory="newCategory"
            :categories="categories"
            @close="closeHabitForm"
            @save="editingHabitId ? updateHabit() : createHabit()"
          />

          <DeleteConfirmModal
            :show="showDeleteModal"
            :habitName="habitToDelete?.name || ''"
            @confirm="deleteHabit"
            @cancel="cancelDelete"
          />

          <!-- Окно видимости с эффектом плавного затухания -->
          <div class="habits-fade-viewport">
            <div class="habits-scroll-window">
              <TransitionGroup name="habit-fade" tag="div" class="habits-wrapper-layout">
                <HabitCard
                  v-for="habit in filteredHabits"
                  :key="habit.id"
                  :habit="habit"
                  :todayStr="todayStr"
                  :canMoveUp="filteredHabits.findIndex(h => h.id === habit.id) > 0"
                  :canMoveDown="filteredHabits.findIndex(h => h.id === habit.id) < filteredHabits.length - 1"
                  :isDoneToday="isHabitDoneToday(habit)"
                  :isUpdating="updatingHabits.has(habit.id)"
                  @toggle="toggleHabit(habit)"
                  @edit="openEditHabit(habit)"
                  @delete="openDeleteModal(habit)"
                  @move-up="moveHabit(habit, 'up')"
                  @move-down="moveHabit(habit, 'down')"
                />
              </TransitionGroup>
            </div>
          </div>
        </section>

        <!-- ПРАВАЯ КОЛОНКА: ФИКСИРОВАННАЯ СТАТИСТИКА -->
        <RightStats />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import defaultAvatar from '@/assets/default-avatar.jpg'
import { config } from '@/config/env'
import DashboardHeader from '@/components/Header/DashboardHeader.vue'
import LoadingState from '@/components/Dashboard/LoadingState.vue'
import DashboardError from '@/components/Dashboard/DashboardError.vue'
import HabitsToolbar from '@/components/Dashboard/HabitsToolbar.vue'
import HabitCreateModal from '@/components/Dashboard/HabitCreateModal.vue'
import UserProfile from '@/components/Dashboard/UserProfile.vue'
import DeleteConfirmModal from '@/components/Dashboard/DeleteConfirmModal.vue'
import RightStats from '@/components/Dashboard/RightStats.vue'
import HabitCard from '@/components/Dashboard/HabitCard.vue'

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

function cancelDelete() {
  showDeleteModal.value = false
  habitToDelete.value = null
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
  LAYOUT & CONTAINERS
========================================= */
.dashboard-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow-x: clip;
  overflow-y: visible;
  padding: clamp(84px, 12vw, 100px) clamp(12px, 2vw, 24px) clamp(24px, 4vw, 40px);
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
  gap: clamp(16px, 2vw, 24px);
  min-height: calc(100vh - 150px);
}

/* =========================================
  MAIN GRID & HABITS NAVIGATION
========================================= */
.main-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(280px, 380px);
  gap: clamp(16px, 2vw, 24px);
  flex: 1;
  overflow: visible;
  min-height: 0;
  padding-bottom: 24px;
}

.habits-container {
  overflow: visible;
  display: flex;
  flex-direction: column;
  gap: clamp(12px, 2vw, 16px);
  min-height: 0;
  min-width: 0;
}
/* =========================================
  ICON BUTTONS (Unified & Clean)
========================================= */
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
  padding: clamp(20px, 3vw, 40px) clamp(16px, 3vw, 36px);
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
   BUTTONS
========================================= */

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-dark));
  color: #ffffff;
  box-shadow: 0 4px 15px color-mix(in srgb, var(--accent-primary) 30%, transparent);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px color-mix(in srgb, var(--accent-primary) 40%, transparent);
}

/* =========================================
   RESPONSIVE
========================================= */
@media (max-width: 1024px) {
  .dashboard-content {
    min-height: auto;
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

@media (max-width: 768px) {
  .dashboard-container {
    padding-top: 84px;
  }

  .dashboard-content {
    gap: 16px;
  }

  .main-grid {
    grid-template-columns: 1fr;
    padding-bottom: 12px;
  }

  .right-stats {
    order: -1;
  }

  .user-profile-section {
    flex-direction: column;
    align-items: flex-start;
  }

  .habits-scroll-window {
    padding: 16px 12px 20px;
  }
}
</style>
