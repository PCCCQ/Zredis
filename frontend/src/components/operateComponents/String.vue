<script setup>
import {toRefs, defineProps, ref, onBeforeMount, onMounted, onBeforeUpdate, onUpdated, watch, reactive} from 'vue'
import {GetString, SetString} from "../../../wailsjs/go/main/App.js";
import {Edit,Select} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";
const props = defineProps({
//子组件接收父组件传递过来的值
  operateKey:""
})
//使用父组件传递过来的值
const {operateKey} =toRefs(props)
let data = reactive({
  key:operateKey.value,
  value:"",
  editValue:"",
  ttl:-1,
  isEditValue:false
})
const tmp = {key:operateKey.value}
GetString({key:data.key}).then(res => {
  if (res.code===200){
    data.value=res.data.value
    console.log(res.data)
  }else {
    ElMessage.error(res.data)
  }
})
// onUpdated( () => {
//   console.log('updated');
//   const tmp = {key:operateKey.value}
//   GetString(tmp).then(res => {
//     if (res.code===200){
//       data.value=res.data.value
//       console.log(res.data)
//     }
//   })
//
// });


// 编辑
function editValue(){
  data.isEditValue=true
  data.editValue=data.value
}

// 保存
function saveValue(){
  data.value= data.editValue
  if(data.value===""){
    ElMessage.error("请输入内容")
    return
  }
  SetString({key:data.key,value:data.value}).then(res=>{
    if(res.code===200){
      ElMessage.success(res.data)

      data.isEditValue=false
    }else {
      ElMessage.error(res.data)
    }
  })
}
// watch(operateKey,(newValue,oldValue)=>{
//   // const tmp = {key:newValue}
//   // GetString(tmp).then(res => {
//   //   if (res.code===200){
//   //     data.value=res.data.value
//   //     console.log(res.data)
//   //   }
//   // })
//   console.log(newValue,oldValue)
// })

</script>

<template>
  <el-text>Value:
    <el-button v-if="data.isEditValue" type="primary" size="small" :icon="Select" text @click="saveValue" >保存</el-button>
    <el-button  v-else type="primary" size="small" :icon="Edit" text @click="editValue">编辑</el-button>
  </el-text>
  <div v-if="data.isEditValue">
    <el-input
        v-model="data.editValue"
        :rows="4"
        type="textarea"
        placeholder="Please input"
    />
  </div>
<div v-else>
{{data.value}}
</div>

</template>

<style scoped>



</style>