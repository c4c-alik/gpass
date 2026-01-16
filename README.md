# gpass - Password Manager

一个使用 Wails + Vue3 + SQLite 构建的桌面密码管理器应用。

## 项目结构

- `main.go` - Go 后端代码，包含 Wails 应用逻辑和 SQLite 数据库操作
- `assets.go` - 嵌入前端资源的 Go 文件
- `wails.json` - Wails 项目配置文件
- `frontend/` - Vue3 前端项目
  - `src/` - Vue 组件源码
  - `index.html` - HTML 入口文件
  - `vite.config.ts` - Vite 构建配置
  - `.eslintrc.cjs` - ESLint 配置文件
  - `tsconfig.json` - TypeScript 配置文件
- `assets/` - 构建后的前端资源（由 Vite 生成）

## 依赖

- **Wails** - 用于构建桌面应用框架
- **Vue3** - 前端框架
- **TypeScript** - 类型安全的 JavaScript 超集
- **ESLint** - JavaScript/TypeScript 代码质量检查工具
- **SQLite** - 本地数据库
- **Pinia** - Vue3 状态管理

## 开发

### 后端开发

```bash
# 运行应用
go run main.go
```

### 前端开发

```bash
# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 类型检查
npm run type-check

# 代码检查和修复
npm run lint

# 格式化代码
npm run format
```

### 构建前端资源

```bash
# 在 frontend 目录中构建
cd frontend
npm run build
```

构建后的资源会自动放在项目根目录的 `assets` 文件夹中，Wails 会自动使用这些资源。

## 功能

- 使用 SQLite 存储密码信息
- Go 后端提供数据处理逻辑
- Vue3 + TypeScript 前端提供用户界面
- ESLint 代码质量保证
- 跨平台桌面应用支持