<template>
  <template v-if="formInline.key">
  <el-form  :inline="true" :model="formInline" class="demo-form-inline">
    <el-form-item label="Key:">
      <el-input v-if="formInline.isEditKey" v-model="formInline.editKey" placeholder="请输入Key" >
        <template #append> <el-button :icon="Select" type="primary" size="default" @click="editKey"></el-button></template>
      </el-input>
      <el-button v-else :icon="Edit" type="primary" :underline="false" text size="default" @click="formInline.isEditKey=true;"> {{formInline.key}}</el-button>
    </el-form-item>
    <el-form-item label="Type:">
<!--      <el-select-->
<!--          v-model="formInline.type"-->
<!--          placeholder="请选择数据类型"-->
<!--          :disabled="formInline.operateTypeIsRead"-->
<!--      >-->
<!--        -->
<!--        <el-option label="string（字符）" value="string" />-->
<!--        <el-option label="list（列表）" value="list" />-->
<!--        <el-option label="set（集合）" value="set" />-->
<!--        <el-option label="zset（有序集）" value="zset" />-->
<!--        <el-option label="hash（哈希表）" value="hash" />-->
<!--      </el-select>-->
     <el-text type="danger">{{formInline.type}}</el-text>
    </el-form-item>
    <el-form-item label="TTL:">
<!--      <el-input v-model="formInline.ttl" placeholder="超时时间（单位秒）" clearable :disabled="formInline.operateTypeIsRead"/>-->
      <el-input v-if="formInline.isEditTTL" v-model="formInline.editTTL" >
        <template #append> <el-button :icon="Select" type="primary" size="default" @click="editTTL"></el-button></template>

      </el-input>
      <el-button v-else :icon="Edit" type="primary" :underline="false" text size="default" @click="formInline.isEditTTL=true"> {{formInline.ttl}}</el-button>
    </el-form-item>
    <el-form-item>
<!--      <el-button v-if="formInline.operateTypeIsRead" type="primary" @click="onRevise" :icon="EditPen" size="small">修改</el-button>-->
<!--      <el-button v-else type="primary" @click="onSave" :icon="EditPen" size="small">保存</el-button>-->
      <el-button type="danger" @click="onDelete" :icon="Delete" size="small">删除</el-button>
    </el-form-item>

  </el-form>
<template v-if="formInline.type=='string'">
      <string-components  :operateKey="formInline.key" :key="formInline.key"/>
</template>
  <template v-if="formInline.type=='hash'">
    <hash-components  :operateKey="formInline.key" :key="formInline.key"/>
  </template>
  <template v-if="formInline.type=='list'">
    <list-components  :operateKey="formInline.key" :key="formInline.key"/>
  </template>
  <template v-if="formInline.type=='set'">
    <set-components  :operateKey="formInline.key" :key="formInline.key"/>
  </template>
  <template v-if="formInline.type=='zset'">
    <zset-components  :operateKey="formInline.key" :key="formInline.key"/>
  </template>
</template>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { defineComponent,ref,getCurrentInstance } from 'vue'
import {Del, GetTTL, SetKey, SetTTL} from "../../wailsjs/go/main/App";
import {Delete, EditPen, Edit, Select} from "@element-plus/icons-vue";
import StringComponents from "./operateComponents/String.vue";
import hashComponents from "./operateComponents/Hash.vue";
import ListComponents from "./operateComponents/List.vue";
import SetComponents from "./operateComponents/Set.vue";
import ZsetComponents from "./operateComponents/Zset.vue";
import {ElMessage} from "element-plus";


let formInline = reactive({
  key: '',
  type: "string",
  ttl: "-1",
  operateTypeIsRead:true,
  editKey:"",
  isEditKey:false,
  editTTL:"-1",
  isEditTTL:false
})

function editKey(){
SetKey(formInline.key,formInline.editKey).then(res=>{
  if (res.code===200){
    ElMessage.success(res.data)
    formInline.key=formInline.editKey
    formInline.isEditKey=false
  }else {
    ElMessage.error(res.data)
  }
})
}

function editTTL(){
  SetTTL(formInline.key,formInline.editTTL).then(res=>{
    if (res.code===200){
      ElMessage.success(res.data)
      formInline.ttl=formInline.editTTL
      formInline.isEditTTL=false
    }else {
      ElMessage.error(res.data)
    }
  })
}

// 删除
const onDelete = () => {
  Del(formInline.key).then(res=>{
    if (res.code===200){
      formInline.key=""
      ElMessage.success(res.data)

    }else {
      ElMessage.error(res.data)
    }
  })
  console.log('onDelete!')
}




let { proxy } = getCurrentInstance()

// 拿到info，获取info内部的值
proxy.$mitt.on('key', (res) => {
  formInline.key=res.key
  formInline.editKey=res.key
  formInline.type=res.type
  GetTTL(res.key).then(result => {
    console.log(result)
    formInline.ttl = result.data
    formInline.editTTL = result.data
  })
  formInline.operateTypeIsRead=true
  console.log(res)
  console.log(formInline)
})



</script>

<style>
.demo-form-inline .el-input {
  --el-input-width: 220px;
}
</style>