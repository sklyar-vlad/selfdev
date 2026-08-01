<template>
  <section class="user-profile-section card-surface">
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
</template>

<script setup lang="ts">
defineProps<{
  user: { user_id: string; role: string; username: string; email: string } | null
  level: number
  xpCurrent: number
  xpNext: number
  xpProgress: number
  perfectDays: number
  streakDays: number
  defaultAvatar: string
}>()
</script>

<style scoped>
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

.lvl-text {
  color: var(--text-primary);
}
.xp-count {
  color: var(--text-secondary);
}

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
</style>
