import Vue from 'vue'
import Router from 'vue-router'
import utils from './utils'

Vue.use(Router)



let router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [{
      path: '/',
      redirect: '/admin'
    },
    {
      path: '/login',
      name: 'login',
      meta: {
        layout: 'login'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/login.vue')
    },
    {
      path: '*',
      meta: {
        layout: 'notfound'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/notfound.vue')
    },
    {
      path: '/admin',
      name: 'admin',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/admin.vue'),
    },
    {
      path: '/admin/change',
      name: 'change',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/change.vue'),
    },
    {
      path: '/admin/account',
      name: 'account',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/account.vue'),
    },
    {
      path: '/admin/user',
      name: 'user',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/index.vue'),
    },
    {
      path: '/admin/user/create',
      name: 'create_user',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/add.vue'),
    },
    {
      path: '/admin/user/edit/:id',
      name: 'edit_user',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/edit.vue'),
    },
    {
      path: '/admin/settings',
      name: 'wa',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/setting/wa.vue'),
    },
    {
      path: '/admin/kategori',
      name: 'kategori',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/kategori/index.vue'),
    },
    {
      path: '/admin/kategori/create',
      name: 'create_kategori',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/kategori/add.vue'),
    },
    {
      path: '/admin/kategori/edit/:id',
      name: 'edit_kategori',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/kategori/edit.vue'),
    },
    {
      path: '/admin/surat_masuk',
      name: 'surat_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_masuk/index.vue'),
    },
    {
      path: '/admin/surat_masuk/create',
      name: 'create_surat_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_masuk/add.vue'),
    },
    {
      path: '/admin/surat_masuk/edit/:id',
      name: 'edit_surat_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_masuk/edit.vue'),
    },
  ]
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    if (utils.session() === null) {
      next({
        path: '/login',
        params: {
          nextUrl: to.fullPath
        }
      })
    } else {
      next()
    }
  } else {
    next() // make sure to always call next()!
  }
})

export default router;