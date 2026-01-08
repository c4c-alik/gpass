<script setup lang="ts">
import {onMounted, ref, Ref} from 'vue'

interface Password {
  id: number
  account: string
  username: string
  password: string
  created_at: string
}

const greeting = ref<string>('Hello Wails + Vue3 + SQLite!')
const passwords: Ref<Password[]> = ref<Password[]>([])
const newAccount = ref<string>('')
const newUsername = ref<string>('')
const newPassword = ref<string>('')

// 示例：调用 Go 后端方法
const callBackend = async (): Promise<void> => {
  try {
    // 这里是调用 Go 后端方法的示例
    // @ts-ignore
    const result: string = await window.go.main.App.Greet('Wails')
    greeting.value = result
  } catch (error) {
    console.error('Error calling backend:', error)
  }
}

// 获取所有密码
const getAllPasswords = async (): Promise<void> => {
  try {
    // @ts-ignore
    const result: Password[] = await window.go.main.App.GetAllPasswords()
    passwords.value = result
  } catch (error) {
    console.error('Error getting passwords:', error)
  }
}

// 添加新密码
const addPassword = async (): Promise<void> => {
  try {
    // @ts-ignore
    await window.go.main.App.CreatePassword(newAccount.value, newUsername.value, newPassword.value)
    // 重置输入框
    newAccount.value = ''
    newUsername.value = ''
    newPassword.value = ''
    // 刷新列表
    await getAllPasswords()
  } catch (error) {
    console.error('Error adding password:', error)
  }
}

onMounted(async () => {
  await getAllPasswords()
})
</script>

<template>
  <div id="app">
    <header>
      <h1>gpass - Password Manager</h1>
    </header>
    <main>
      <div class="card">
        <h2>{{ greeting }}</h2>
        <button @click="callBackend">Call Backend</button>
        <p>Manage your passwords securely with Wails + Vue3 + SQLite</p>
      </div>

      <!-- 添加密码表单 -->
      <div class="card">
        <h3>Add New Password</h3>
        <form @submit.prevent="addPassword">
          <div class="input-group">
            <label for="account">Account:</label>
            <input
                id="account"
                v-model="newAccount"
                type="text"
                placeholder="e.g., Google Account"
                required
            />
          </div>
          <div class="input-group">
            <label for="username">Username:</label>
            <input
                id="username"
                v-model="newUsername"
                type="text"
                placeholder="your-email@example.com"
                required
            />
          </div>
          <div class="input-group">
            <label for="password">Password:</label>
            <input
                id="password"
                v-model="newPassword"
                type="password"
                placeholder="••••••••"
                required
            />
          </div>
          <button type="submit">Add Password</button>
        </form>
      </div>

      <!-- 密码列表 -->
      <div class="card">
        <h3>Saved Passwords</h3>
        <div v-if="passwords.length === 0" class="empty-state">
          <p>No passwords saved yet.</p>
        </div>
        <ul v-else class="password-list">
          <li v-for="pwd in passwords" :key="pwd.id" class="password-item">
            <div class="password-info">
              <strong>{{ pwd.account }}</strong
              ><br/>
              <span>Username: {{ pwd.username }}</span
              ><br/>
              <small>Created: {{ new Date(pwd.created_at).toLocaleString() }}</small>
            </div>
          </li>
        </ul>
      </div>
    </main>
  </div>
</template>

<style>
#app {
  max-width: 1280px;
  margin: 0 auto;
  padding: 2rem;
  font-family: Inter, Avenir, Helvetica, Arial, sans-serif;
}

header {
  text-align: center;
  margin-bottom: 2rem;
}

h1 {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

h2,
h3 {
  font-weight: 600;
  margin-bottom: 1rem;
}

.card {
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
  background: white;
}

.input-group {
  margin-bottom: 1rem;
}

.input-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
}

.input-group input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

button {
  padding: 0.5rem 1rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin: 1rem 0;
}

button:hover {
  background-color: #369870;
}

.empty-state {
  text-align: center;
  color: #888;
  font-style: italic;
}

.password-list {
  list-style: none;
  padding: 0;
}

.password-item {
  padding: 1rem;
  border: 1px solid #eee;
  border-radius: 4px;
  margin-bottom: 0.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.password-info {
  flex-grow: 1;
}

.password-actions button {
  padding: 0.25rem 0.5rem;
  font-size: 0.8rem;
  margin-left: 0.5rem;
}
</style>
