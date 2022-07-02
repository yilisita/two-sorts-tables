/*
 * @Author: Wen Jiajun
 * @Date: 2022-04-22 21:01:12
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 17:00:36
 * @FilePath: \application\frontend\app1\src\router\router.js
 * @Description: 
 */
import { createApp } from 'vue' // 到底应该怎么引入？？？？？？
import { createRouter, createWebHashHistory } from "vue-router"
import ElementPlus from 'element-plus'
import request from './app.vue'
// import app from './navgator.vue'
import axios from 'axios'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import home from "./Home.vue"
import table from "./table.vue"


var app = createApp(home)
// var app = Vue.createApp({}) 

app.config.globalProperties.$axios=axios

axios.defaults.baseURL = "http://localhost:4000/v1"
app.use(ElementPlus)

for (let key of Object.keys(ElementPlusIconsVue)) {
    app.component(key, ElementPlusIconsVue[key]);
  }






const Home = { template: '<div>首页</div>' }
// const About = { template: '<div>About</div>' }

// // 2. Define some routes
// // Each route should map to a component.
// // We'll talk about nested routes later.
const routes = [
  { path: '/', component: Home },
  { path: '/v1/table', component: table },
  { path: '/v1/request', component: request},
  { path: '/v1/login', component: login},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes, 
})

app.use(router)

app.mount('#app')