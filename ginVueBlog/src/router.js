import Vue from 'vue'
import VueRouter from 'vue-router'

import Dashboard from '@/views/admin/Dashboard'
import ArticleEdit from '@/views/admin/Article'
import ArticleList from '@/views/admin/ArticleList'
import CommentList from '@/views/admin/CommentList'
import MetaList from '@/views/admin/MetaList'
import MediaList from '@/views/admin/MediaList'
import Setting from '@/views/admin/Setting'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        redirect: 'home',
        component: () => import('@/views/blog/Blog.vue'),
        children: [
            { path: 'home', name: 'Home', component: () => import('@/views/blog/Home.vue') },
            { path: 'about', name: 'About', component: () => import('@/views/blog/About.vue') },
            { path: 'stories', name: 'Stories', component: () => import('@/views/blog/Stories.vue') },
            { path: 'categories', name: 'Categories', component: () => import('@/views/blog/Categories.vue') },
            { path: 'stories/:id', name: 'StoryView', component: () => import('@/views/blog/StoryView.vue') },
        ]
    },
    { path: '/login', name: 'Login', component: () => import('@/views/admin/Login.vue') },
    { path: '/register', name: 'Register', component: () => import('@/views/admin/Register.vue') },
    {
        path: '/admin',
        name: 'Admin',
        component: () => import('@/views/admin/Admin.vue'),
        redirect: 'admin/dashboard',
        children: [
            {
                path: '/admin/dashboard',
                name: 'Dashboard',
                component: Dashboard,
              },
              {
                path: '/admin/article/publish/:id',
                name: 'ArticleEdit',
                component: ArticleEdit,
              },
              {
                path: '/admin/article/publish',
                name: 'ArticleNew',
                component: ArticleEdit,
              },
              {
                path: '/admin/article',
                name: 'ArticleList',
                component: ArticleList,
              },
              {
                path: '/admin/comment',
                name: 'CommentList',
                component: CommentList,
              },
              {
                path: '/admin/meta',
                name: 'MetaList',
                component: MetaList,
              },
              {
                path: '/admin/media',
                name: 'MediaList',
                component: MediaList,
              },
              {
                path: '/admin/setting',
                name: 'Setting',
                component: Setting,
              },
            // { path: 'editor', name: 'Editor', component: () => import('@/views/admin/Editor.vue') },
            // { path: '/comment', name: 'Comment', component: () => import('@/views/admin/Comment.vue') },
        ]
    },
    
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
  })


export default router 