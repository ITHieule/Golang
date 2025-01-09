import { createRouter, createWebHistory } from 'vue-router';
import LoginPage from '../components/LoginPage.vue';
import TodoApp from '../components/TodoApp.vue';

const routes = [
  { path: '/login', component: LoginPage },
  { path: '/todo', component: TodoApp, meta: { requiresAuth: true } },
  { path: '/:pathMatch(.*)*', redirect: '/login' }, // Wildcard redirect
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guard for authentication
router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('isLoggedIn'); // Check login status
  if (to.meta.requiresAuth && !isLoggedIn) {
    next('/login'); // Redirect to login if not logged in
  } else {
    next(); // Proceed if valid
  }
});

export default router;
