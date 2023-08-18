<script setup>

import DataList from "../components/DataList.vue";
import DisplayEdit from "../components/DisplayEdit.vue";
import {AnyCmd,GetDBList,ChangeDB} from "../../wailsjs/go/main/App.js";
import {reactive, ref, watch} from "vue";
import Cmd from "../components/Cmd.vue";
import {Back, Postcard} from "@element-plus/icons-vue";
import AddData from "../components/AddData.vue";
const cmd= ref('')
const showTerminal= ref(false)
const showData=ref("")
const data=reactive({
  DBList:[{DBName:"db0",DBNumber:0}],
  value:0,
  time:0
})

const childCmd=ref(null)
function openTerminal(){
  childCmd.value.drawer=true
}
const childAddData=ref(null)
function addDataFunc(){
  childAddData.value.isShow=true
}

function getDB(){
  GetDBList().then(res=>{
    if(res.code===200){
      data.DBList=res.data
      console.log(res.data)
    }
  })
}
getDB()
function valueChange(val){
  ChangeDB(val).then(res=> {
    console.log(res)
    changeTime()
  })
  }
  function changeTime(){
    data.time=new Date().getTime()
  }
</script>

<template>
  <el-container>
    <el-header>
<!--      Header-->
      <div class="header-text">
        <router-link to="/"><el-link  :icon="Back" :underline="false">返回</el-link></router-link>
        <span style="margin-left: 20px"></span>
        <el-button type="primary" size="small" @click="addDataFunc">新增数据</el-button>
</div>
    </el-header>

    <el-container>
      <el-aside width="300px"><data-list :key="data.time"/></el-aside>
      <el-main>
<!--        <HelloWorld/>-->
      <display-edit :key="data.time"></display-edit>
      </el-main>
    </el-container>
    <el-footer>
<!--      <span><el-input v-model="cmd"></el-input>-->
<!--      <el-button @click="exec">执行</el-button>{{showData}}-->
<!--      </span>-->
      <el-select v-model="data.value" placeholder="Select" size="small" @change="valueChange">
        <el-option
            v-for="item in data.DBList"
            :label="item.DBName"
            :value="item.DBNumber"
        />
      </el-select>
      <span style="margin-left: 100px"></span>
      <el-link @click="openTerminal" :underline="false" :icon="Postcard">执行redis命令</el-link>
      <span style="margin-left: 100px"></span>
      <el-link href="https://github.com/PCCCQ" target="_blank" :underline="false">version:1.0.0 @PCCCQ</el-link>

    </el-footer>
  </el-container>
  <cmd ref="childCmd"></cmd>
<add-data ref="childAddData" :time="data.time"></add-data>
</template>
<style scoped>

.el-header{
  background-color: #c8cdd9;
  color: #333;
  //text-align: center;
  height: 5vh;
}
.el-footer{
  background-color: #c8cdd9;
  color: #333;
text-align: right;
  height: 5vh;
  line-height: 4.5vh;
}
.header-text{
  line-height: 5vh;
}
.el-aside {
  background-color: #ffffff;
  //color: #333;
  //text-align: center;
  //line-height: 300px;
}

.el-main {
  background-color: #f5f5f5;
  color: #333;
  border-left:#333;
  //text-align: center;
//line-height: 160px;
  height: 90vh;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>