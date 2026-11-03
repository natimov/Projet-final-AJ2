import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import AdminUsersView from '../views/AdminUsersView.vue'
import ComingSoonView from '../views/ComingSoonView.vue'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    {
      path:'/login',
      name:'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/espace-particulier',
      name: 'particular-space',
      component: ComingSoonView,
      meta: { titleKey: 'spaces.particular' },
    },
    {
      path: '/espace-professionnel',
      name: 'professional-space',
      component: ComingSoonView,
      meta: { titleKey: 'spaces.professional' },
    },
    {
      path: '/espace-salarie',
      name: 'employee-space',
      component: ComingSoonView,
      meta: { titleKey: 'spaces.employee' },
    },
    {
      path: '/admin',
      name: 'admin-home',
      component: AdminUsersView,
      meta: { requiresAdmin: true },
    },
    {
      path: '/admin/utilisateurs',
      name: 'admin-users',
      redirect: '/admin',
    },
  ],
})
router.beforeEach((to) => {
  if (to.meta.requiresAdmin) {
    const authStore = useAuthStore()

    if (!authStore.user || authStore.user.role !== 'admin') {
      return { path: '/login', query: { redirect: to.fullPath } }
    }
  }
})
export default router
