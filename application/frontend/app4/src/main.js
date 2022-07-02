/*
 * @Author: Wen Jiajun
 * @Date: 2022-07-02 18:41:27
 * @LastEditors: Wen Jiajun
 * @LastEditTime: 2022-07-02 19:30:34
 * @FilePath: \application\frontend\app4\src\main.js
 * @Description: 
 */
import { createApp } from 'vue'
import App from './App.vue'
import './index.css'

import ElementPlus from 'element-plus'

import axios from 'axios'



var app = createApp(App)
app.use(ElementPlus)
// var app = Vue.createApp({}) 

axios.defaults.baseURL = "http://localhost:4000/v1"
app.config.globalProperties.$axios=axios


app.mount('#app')
