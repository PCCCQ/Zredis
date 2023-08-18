<template>
  <el-drawer
      v-model="drawer"
      title="redis命令执行"
      direction="btt"
      size="60%"
      :with-header="false"
  >
<div class="showCase">
  <el-scrollbar>
    <p v-for="item in execData" :key="item">
      <template v-if="item.code===200">
      {{ item.command }} >>>   <el-tag type="success" size="small">success</el-tag> {{ item.data }}
      </template>
      <template v-else>
        {{ item.command }} >>> <el-tag type="danger" size="small">err</el-tag> {{ item.data }}
      </template>
    </p>
  </el-scrollbar>
</div>

    <el-input v-model="command" @keyup.enter="exec" placeholder="请输入命令,按回车键执行"></el-input>
    <el-link href="https://redis.com.cn/commands.html" target="_blank" :underline="false">redis命令手册</el-link>
  </el-drawer>

</template>

<script setup>
import {ref} from "vue";
import {AnyCmd} from "../../wailsjs/go/main/App.js";
const command=ref('')
const execData=ref([])
function  exec(){
  AnyCmd(command.value).then(res=>{
    console.log(res)
    execData.value.push({code:res.code,data:res.data,command:command.value})
    if(res.code===200){
      command.value=''
    }
  })
}
const drawer=ref(false)
defineExpose({drawer})

</script>


<style scoped>
.showCase{
  height: 85%;
}

</style>