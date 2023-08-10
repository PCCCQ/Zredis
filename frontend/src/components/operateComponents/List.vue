<script setup>
import {Delete, Edit, Plus} from "@element-plus/icons-vue";
import {DelListOne, DelSetOne, GetList, GetSet, SetList, SetListSet, SetSet} from "../../../wailsjs/go/main/App.js";
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
  index:0,
  addTab:"before",
  editValue:"",
  isEditValue:false
})
let dialogFormVisible = ref(false)
function getList(){GetList({key:data.key}).then(res=>{
  console.log("执行")
  if(res.code===200){
    data.value=res.data['value']
    data.length=data.value.length
  }else {
    ElMessage.error(res.data)
  }
})}
getList()
const textarea = ref('')

function  addListOne(){
  if(textarea.value===""){
    ElMessage.error("请输入内容")
    return
  }

    SetList({key:data.key,value:textarea.value,addTab:data.addTab}).then(res=>{
      console.log("res",res)
      if(res.code===200){
        getList()
      }else {
        ElMessage.error(res.data)
      }
    })
  // SetList({key:data.key},data.value[data.index],textarea.value).then(res=>{
  //   console.log("res",res)
  //   if(res.code===200){
  //     data.index=0
  //     getList()
  //   }else {
  //     ElMessage.error(res.data)
  //   }
  // })
}




function delListOne(value){
  DelListOne({key:data.key,value:value}).then(res=>{
    if(res.code===200){
      getList()
      // ElMessage.success(res.data)
    }else {
      ElMessage.error(res.data)
    }
  })
}
const editTextValue = ref('')
let editTextValueIndex=0
function editListOne(){
if (editTextValue.value===""){
  ElMessage.error("请输入内容")
  return
}
  SetListSet({key:data.key,value:editTextValue.value,index:editTextValueIndex}).then(res=>{
    if(res.code===200){

      getList()
    }else {
      ElMessage.error(res.data)
    }
  })
  dialogFormVisible =ref(false)
}



</script>

<template>
<!--    <el-text> length: {{data.length}}</el-text>-->
  <el-table :data="data.value"  stripe style="width: 100%">
    <el-table-column type="index"  width="50" >
      <template #header>
        <el-text>{{data.length}}</el-text>
      </template>
    </el-table-column>
    <el-table-column  label="Date" >
      <template #header>

        <el-input
            v-model="textarea"
            placeholder="新增数据"
        >
          <template #prepend>

            <el-select v-model="data.addTab" style="width: 70px">
<!--              <el-option v-for="(_,i) in data.length" :label="i+1+'->'" :value="i"/>-->
              <el-option label="前" value="before"/>
              <el-option label="后" value="after"/>
            </el-select>

          </template>
        </el-input>

      </template>
      <template #default="scope">
        <template v-for="item in scope.row">{{item}}</template>
      </template>
    </el-table-column>
    <el-table-column label="操作" width="70">

      <template #header>
        <el-button type="text" @click="addListOne">插入</el-button>
      </template>
      <template #default="scope">
        <el-button type="text"  :icon="Edit" @click="editTextValueIndex=scope.$index;editTextValue=scope.row;dialogFormVisible = true"></el-button>
        <el-button type="text"  :icon="Delete" @click="delListOne(scope.row);"></el-button>
      </template>

    </el-table-column>
  </el-table>
  <el-dialog v-model="dialogFormVisible" title="修改数据" :show-close="false" :close-on-click-modal="false">
    <el-input
        v-model="editTextValue"
        autosize
        type="textarea"
        placeholder="修改数据"
    >
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="editListOne();">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>