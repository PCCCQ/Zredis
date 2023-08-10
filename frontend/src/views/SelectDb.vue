<template>
  <div class="box-father">
    <el-card class="box-card" shadow="hover">
      <div style="text-align: center">
<!--        <el-image src="../assets/images/add.png" fit="cover" />-->
        <img src="../assets/images/add.png" style="width: 80%; height: 80%"  alt="添加">
        <el-link type="primary" @click="toAdd">添加新连接</el-link>
      </div></el-card>




  <el-card class="box-card" shadow="hover" v-for="(item,index) in redisList">
    <template #header>
      <div class="card-header">
        <el-link type="primary" :icon="Pointer" @click="toLink(index)">{{item.name}}</el-link>

        <el-button type="danger" plain size="small" @click="delRedis(index)">删除</el-button>
      </div>
    </template>
    <div>{{ item.IP}}:{{item.port}}</div>
    <div>{{item.userName}}</div>
  </el-card>



  </div>
</template>

<style scoped>
.box-father{
  margin: 20px;
  display: flex;
  flex-wrap: wrap;
  //background: #67c23a;
  //height: 100vh;
}
.card-header {
  display: flex;
  justify-content: space-between;
  //align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  margin: 10px;
  height: 200px;
  width: 200px;

}
</style>
<script setup>
import { Pointer} from '@element-plus/icons-vue'
import {RedisInit} from "../../wailsjs/go/main/App.js";
import router from "../router/index.js";
import {ElMessage} from "element-plus";
import {ref} from "vue";

// export default {
//   methods: {
//     // 使用方法一：使用名称跳转
//     navigateToAboutPage() {
//       this.$router.push({ name: 'About' });
//     },
//
//     // 使用方法二：使用路径跳转
//     navigateToHome() {
//       this.$router.push('/');
//     }
//   }
// }\
function toAdd(){
  router.push('/login');
}

// window.localStorage.setItem("redisList",'[{name:"test1",IP:"127.0.0.1",port:"6379",userName:"",password:""}]')
let localStorageRedisList=eval('(' + window.localStorage.getItem("redisList") + ')');

let redisList=ref(localStorageRedisList)
function delRedis(index){
  console.log("删除")
  redisList.value.splice(index,1)
  window.localStorage.setItem("redisList",JSON.stringify(redisList.value))
}





function toLink(index) {
  let redisTmp=redisList.value
  RedisInit(redisTmp[index]).then(result => {
    console.log(result)
    if(result.code===200){
      console.log(111)
      router.push('/index');
    }else {
      ElMessage.error('失败')
    }
  })
}

</script>