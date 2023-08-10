import {createApp} from 'vue'

import ElementPlus from 'element-plus'  //引入element-plus库
import zhCn from 'element-plus/dist/locale/zh-cn.mjs' //element-plus中文

import VForm3 from 'vform3-builds'  //引入VForm3库
import 'element-plus/dist/index.css'  //引入element-plus样式
import 'vform3-builds/dist/designer.style.css' //引入VForm3样式
import App from "./App.vue";
import router from './router/index.js';
import mitt from 'mitt' // 导入mitt



const app = createApp(App)
app.config.globalProperties.$mitt = new mitt() // mitt在vue3中挂载到全局
app.use(router);
app.use(ElementPlus,{
    locale: zhCn,
})  //全局注册element-plus

app.use(VForm3)  //全局注册VForm3(同时注册了v-form-designe、v-form-render等组件)
app.mount('#app')
