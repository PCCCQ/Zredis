<script setup>
import {Delete, Edit, Plus} from "@element-plus/icons-vue";
import {DelSetOne, GetSet, SetSet} from "../../../wailsjs/go/main/App.js";
import {defineProps, reactive, ref, toRefs} from "vue";
import {ElMessage} from "element-plus";
const props = defineProps({
//子组件接收父组件传递过来的值
  operateKey:""
})
//使用父组件传递过来的值
const {operateKey} =toRefs(props)
let data = reactive({
  key:operateKey.value,
  value:[],
  length:0,
  editValue:"",
  isEditValue:false
})
function getSet(){GetSet({key:data.key}).then(res=>{
  console.log("执行")
  if(res.code===200){
    data.value=res.data['value']
    data.length=data.value.length
  }else {
    ElMessage.error(res.data)
  }
})}
getSet()
const textarea = ref('')
function  addSetOne(){
  if(textarea.value===""){
    ElMessage.error("请输入内容")
    return
  }
  SetSet({key:data.key,value:textarea.value}).then(res=>{
    if(res.code===200){
      getSet()
    }else {
      ElMessage.error(res.data)
    }
  })
}




function delSetOne(value){
  DelSetOne({key:data.key,value:value}).then(res=>{
    if(res.code===200){
      getSet()
      // ElMessage.success(res.data)
    }else {
      ElMessage.error(res.data)
    }
  })
}




</script>

<template>
  <el-text> length: {{data.length}}</el-text>
  <el-table :data="data.value"  stripe style="width: 100%">
    <el-table-column type="index"  width="50" >
<!--      <template #header>-->
<!--        <el-text>-->
<!--          {{data.length}}-->

<!--        </el-text>-->
<!--      </template>-->
    </el-table-column>
    <el-table-column  label="Date" >
      <template #header>
        <el-input
            v-model="textarea"
            autosize
            type="textarea"
            placeholder="新增数据"
        />
      </template>
    <template #default="scope">
      <template v-for="item in scope.row">{{item}}</template>
    </template>
    </el-table-column>
    <el-table-column label="操作" width="70">

      <template #header>
        <el-button type="text" @click="addSetOne">确定</el-button>
      </template>
      <template #default="scope">
<!--        <el-button type="text"  :icon="Edit"></el-button>-->
        <el-button type="text"  :icon="Delete" @click="delSetOne(scope.row)"></el-button>
      </template>

    </el-table-column>
  </el-table>

</template>

<style scoped>

</style>