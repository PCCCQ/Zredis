<script setup>
import {defineProps, reactive, toRefs} from "vue";
import {DelHashOne, GetHash, SetHash} from "../../../wailsjs/go/main/App.js";
import {Delete, Edit} from "@element-plus/icons-vue";
import {ElMessage} from "element-plus";

const props = defineProps({
//子组件接收父组件传递过来的值
  operateKey:""
})
//使用父组件传递过来的值
const {operateKey} =toRefs(props)
let data=reactive({
  key:operateKey.value,
  valueList:[],
  dialogFormVisible:false,
  editValue:{key:"",value:""},
  addValue:{key:"",value:""},
})
function getHAshFunc(){
  GetHash({key:data.key}).then(res=>{
    console.log(res)
    data.valueList=res.data.valueList
  })
}
getHAshFunc()

function addValueFunc(){
  if(data.addValue.key===""||data.addValue.value===""){
    ElMessage.error("请输入key和value")
    return
  }
SetHash({key:data.key,valueList:[data.addValue]}).then(res=>{
  if (res.code===200){
    ElMessage.success("添加成功")
    getHAshFunc()
    data.addValue.key=""
    data.addValue.value=""
  }else {
    ElMessage.error(res.data)
  }
})

}
function editValueFunc(){
  if(data.editValue.key===""||data.editValue.value===""){
    ElMessage.error("请输入key和value")
    return
  }
  SetHash({key:data.key,valueList:[data.editValue]}).then(res=>{
    if (res.code===200){
      ElMessage.success("修改成功")
      getHAshFunc()
      // data.editValue.key=""
      // data.editValue.value=""
      data.dialogFormVisible=false
    }else {
      ElMessage.error(res.data)
    }
  })
}
function delValueFunc(delValue){
  DelHashOne({key:data.key,valueList:[delValue]}).then(res=>{
    if(res.code===200){
      ElMessage.success("删除成功")
      getHAshFunc()
    }else {
      ElMessage.error(res.data)
    }

  })
}
</script>

<template>
hash
<!--  {{data.valueList}}-->
  <el-table :data="data.valueList">
<!--    <el-table-column type="index" width="50"/>-->
    <el-table-column prop="key" label="key" >
      <template #header>
        <el-input
            v-model="data.addValue.key"
            placeholder="Key"
        />
      </template>
    </el-table-column>
    <el-table-column prop="value" label="value" >
      <template #header>
        <el-input
            v-model="data.addValue.value"
            placeholder="Value"
        />
      </template>
    </el-table-column>
    <el-table-column label="操作" width="80px">
      <template #header>
        <el-button type="text" @click="addValueFunc">插入</el-button>
      </template>
      <template #default="scope">
        <el-button type="text"  :icon="Edit" @click="data.dialogFormVisible=true;data.editValue=scope.row" ></el-button>
        <el-button type="text"  :icon="Delete" @click="delValueFunc(scope.row)"></el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-dialog v-model="data.dialogFormVisible" title="修改数据" :show-close="false" :close-on-click-modal="false">
    <el-form >
      <el-form-item label="key">
<!--        <el-input v-model="data.editValue.key" autocomplete="off" />-->
        {{data.editValue.key}}
      </el-form-item>
      <el-form-item label="value">
        <el-input v-model="data.editValue.value" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="data.dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="editValueFunc">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>