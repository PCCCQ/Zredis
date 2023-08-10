import { createRouter, createWebHistory } from 'vue-router';

// 导入你的组件
import SelectDb from '../views/SelectDb.vue';
import index from '../views/index.vue';
import login from '../views/Login.vue';

const routes = [
    { path: '/', component: SelectDb },
    { path: '/index', component: index },
    { path: '/login', component: login }
    // 添加其他路由规则
    // { path: '/:catchAll(.*)', component: SelectDb },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
