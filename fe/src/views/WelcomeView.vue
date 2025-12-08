<template>
  <div class="welcome-container">
    <div class="content">
      <img src="@/assets/logo.svg" alt="PMail Logo" class="logo" />
      <h1>{{ lang.tks_pmail }}</h1>
      <p class="description">PMail Pro</p>
      <div class="actions">
        <el-button type="primary" size="large" class="enter-btn" @click="start">
          {{ lang.login }} / {{ lang.inbox }}
          <el-icon class="el-icon--right"><ArrowRight /></el-icon>
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { ArrowRight } from '@element-plus/icons-vue';
import lang from '../i18n/i18n';

const router = useRouter();

const start = () => {
  router.push('/list');
};
</script>

<style scoped>
.welcome-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-color: #f5f7fa;
  /* Use a subtle gradient or pattern if desired */
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.content {
  text-align: center;
  background: rgba(255, 255, 255, 0.9);
  padding: 60px 80px;
  border-radius: 16px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1);
  animation: fadeIn 0.8s ease-out;
}

.logo {
  width: 120px;
  height: 120px;
  margin-bottom: 24px;
  animation: float 6s ease-in-out infinite;
}

h1 {
  font-size: 28px;
  color: #303133;
  margin-bottom: 12px;
  font-weight: 600;
}

.description {
  font-size: 18px;
  color: #606266;
  margin-bottom: 40px;
  letter-spacing: 1px;
}

.enter-btn {
  padding: 12px 30px;
  font-size: 16px;
  border-radius: 8px;
  transition: transform 0.2s;
}

.enter-btn:hover {
  transform: translateY(-2px);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
  100% {
    transform: translateY(0px);
  }
}
</style>
