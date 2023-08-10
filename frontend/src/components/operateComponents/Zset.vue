<template>
  <el-table :data="tableData.valueList" stripe style="width: 100%">
    <el-table-column type="index">
      <template #header>
        <el-text>{{tableData.valueList.length}}</el-text>
      </template>
    </el-table-column>
    <el-table-column prop="Member" label="Member"  >
      <template #header>
        <el-input v-model="tableData.addValue.Member" placeholder="Member"></el-input>
      </template>
    </el-table-column>
    <el-table-column prop="Score" label="Score" width="100" >
      <template #header>
        <el-input v-model.number="tableData.addValue.Score" type="number" placeholder="Score"></el-input>
      </template>
      <template #default="scope">
        <el-input v-model.number="tableData.editValue.Score" v-if="scope.$index===tableData.editIndex" type="number"></el-input>
        <span v-else>{{scope.row.Score}}</span>

      </template>
    </el-table-column>
    <el-table-column label="操作" width="70" >
      <template #header>
        <el-button type="text" @click="addZset">新增</el-button>
      </template>
      <template #default="scope">
        <el-button type="text"  :icon="Select" v-if="scope.$index===tableData.editIndex" @click="editZset(scope.$index)"></el-button>
        <el-button type="text"  :icon="Edit" v-else @click="editButton(scope)"></el-button>
        <el-button type="text"  :icon="Delete" @click="delZsetOne(scope.row.Member)"></el-button>
      </template>
    </el-table-column>
  </el-table>

</template>

<script setup>
import {AddZset, DelZsetOne, GetZset, SetZsetScore} from "../../../wailsjs/go/main/App";
import {defineProps, reactive, ref, toRefs} from "vue";
import {Delete, Edit, Select} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";
const props = defineProps({
//子组件接收父组件传递过来的值
  operateKey:""
})
//使用父组件传递过来的值
const {operateKey} =toRefs(props)
let tableData=reactive({
  key:operateKey.value,
  valueList:[],
  value:{Member:"",Score:null},
  addValue:{Member:"",Score:null},
  editValue:{Member:"",Score:null},
  editIndex:-1
})
function getZetValue(){
  GetZset(tableData).then(res=>{
    tableData.valueList=res.data.valueList
  })
}
getZetValue()


function addZset(){
  // console.log(tableData)
  if(tableData.addValue.Member===""){
    ElMessage.error("请输入内容")
    return
  }
  tableData.value=tableData.addValue
  AddZset(tableData).then(res=>{
    if(res.code===200){
      getZetValue()
    }else {
      ElMessage.error(res.data)
    }
  })
}
let oldScore=0
function editButton(scope){
  tableData.editIndex=scope.$index
  tableData.editValue=scope.row
  oldScore=tableData.editValue.Score
}

function editZset(index){
  // console.log(tableData)
  if(tableData.editValue.Score===""){
    ElMessage.error("请输入内容")
    return
  }
  tableData.editValue.Score=tableData.editValue.Score-oldScore
  tableData.value=tableData.editValue
  SetZsetScore(tableData).then(res=>{
    if(res.code===200){
      tableData.editIndex=-1
      getZetValue()
    }else {
      ElMessage.error(res.data)
    }
  })
}

function delZsetOne(Member){
  console.log(Member)
  DelZsetOne({key:tableData.key,value:{Member:Member}}).then(res=>{
    console.log(tableData,res)
    if(res.code===200){
      getZetValue()
    }else {
      ElMessage.error(res.data)
    }
  })
}
</script>