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
      path: '/admin/berkas',
      name: 'berkas',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/berkas/index.vue'),
    },
    {
      path: '/admin/berkas/create',
      name: 'create_berkas',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/berkas/add.vue'),
    },
    {
      path: '/admin/berkas/edit/:id',
      name: 'edit_berkas',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/berkas/edit.vue'),
    },
    {
      path: '/admin/petugas_ukur',
      name: 'petugas_ukur',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/petugas_ukur/index.vue'),
    },
    {
      path: '/admin/petugas_ukur/detail/:id',
      name: 'detail_petugas_ukur',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/petugas_ukur/detail.vue'),
    },
    {
      path: '/admin/petugas_gambar',
      name: 'petugas_gambar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/petugas_gambar/index.vue'),
    },
    {
      path: '/admin/petugas_gambar/detail/:id',
      name: 'detail_petugas_gambar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/petugas_gambar/detail.vue'),
    },
    {
      path: '/admin/problem',
      name: 'problem',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/problem/index.vue'),
    },
    {
      path: '/admin/problem/detail/:id',
      name: 'detail_problem',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/problem/detail.vue'),
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