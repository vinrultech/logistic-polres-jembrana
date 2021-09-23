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
      path: '/admin/user/reset/:id',
      name: 'reset_user',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/user/reset_password.vue'),
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
      path: '/admin/surat_masuk/edit/:row_id',
      name: 'edit_surat_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_masuk/edit.vue'),
    },
    {
      path: '/admin/surat_masuk/info/:row_id',
      name: 'info_surat_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_masuk/info.vue'),
    },
    {
      path: '/admin/surat_keluar',
      name: 'surat_keluar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_keluar/index.vue'),
    },
    {
      path: '/admin/surat_keluar/create',
      name: 'create_surat_keluar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_keluar/add.vue'),
    },
    {
      path: '/admin/surat_keluar/edit/:id',
      name: 'edit_surat_keluar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/surat_keluar/edit.vue'),
    },
    {
      path: '/admin/juknis',
      name: 'juknis',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/juknis/index.vue'),
    },
    {
      path: '/admin/juknis/create',
      name: 'create_juknis',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/juknis/add.vue'),
    },
    {
      path: '/admin/juknis/edit/:id',
      name: 'edit_juknis',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/juknis/edit.vue'),
    },
    {
      path: '/admin/inventory',
      name: 'inventory',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory/index.vue'),
    },
    {
      path: '/admin/inventory/create',
      name: 'create_inventory',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory/add.vue'),
    },
    {
      path: '/admin/inventory/edit/:id',
      name: 'edit_inventory',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory/edit.vue'),
    },
    {
      path: '/admin/asset_tetap',
      name: 'asset_tetap',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/asset_tetap/index.vue'),
    },
    {
      path: '/admin/asset_tetap/create',
      name: 'create_asset_tetap',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/asset_tetap/add.vue'),
    },
    {
      path: '/admin/asset_tetap/edit/:id',
      name: 'edit_asset_tetap',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/asset_tetap/edit.vue'),
    },
    {
      path: '/admin/inventory_masuk',
      name: 'inventory_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory_masuk/index.vue'),
    },
    {
      path: '/admin/inventory_masuk/create',
      name: 'create_inventory_masuk',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory_masuk/add.vue'),
    },
    {
      path: '/admin/inventory_keluar',
      name: 'inventory_keluar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory_keluar/index.vue'),
    },
    {
      path: '/admin/inventory_keluar/create',
      name: 'create_inventory_keluar',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/inventory_keluar/add.vue'),
    },
    {
      path: '/admin/unit_kerja',
      name: 'unit_kerja',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/unit_kerja/index.vue'),
    },
    {
      path: '/admin/unit_kerja/create',
      name: 'create_unit_kerja',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/unit_kerja/add.vue'),
    },
    {
      path: '/admin/unit_kerja/edit/:id',
      name: 'edit_unit_kerja',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/unit_kerja/edit.vue'),
    },
    {
      path: '/admin/metric',
      name: 'metric',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/metric/index.vue'),
    },
    {
      path: '/admin/metric/create',
      name: 'create_metric',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/metric/add.vue'),
    },
    {
      path: '/admin/metric/edit/:id',
      name: 'edit_metric',
      meta: {
        requiresAuth: true,
        layout: 'admin'
      },
      component: () => import( /* webpackChunkName: "about" */ './views/metric/edit.vue'),
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