<template>
  <el-table
      ref="singleTableRef"
      :data="filterTableData"
      highlight-current-row
      style="width: 100%"
      @current-change="handleCurrentChange"
      height="90vh"
  >
    <el-table-column type="index" >
      <template #header>
        <el-button :icon="RefreshRight" circle size="small" type="primary" @click="showAll"></el-button>
<!--        <el-icon><RefreshRight /></el-icon>-->
      </template>
    </el-table-column>
    <el-table-column property="type" label="type" width="80" :filters="types"  :filter-method="filterTag">
      <template #default="scope">
        <el-tag size="small" :type="tagType[scope.row.type]">{{scope.row.type}}</el-tag>
      </template>
    </el-table-column>

    <el-table-column property="key" label="key" >
      <template #header>
        <el-input v-model="search"  placeholder="输入key搜索" size="small" />
      </template>
    </el-table-column>
  </el-table>
<!--  <div style="margin-top: 20px">-->
<!--    <el-button @click="setCurrent(tableData[1])">Select second row</el-button>-->
<!--    <el-button @click="setCurrent()">Clear selection</el-button>-->
<!--  </div>-->
</template>

<script  setup>
import {computed, ref} from 'vue'
import { ElTable } from 'element-plus'
import {AllKeyTypes} from "../../wailsjs/go/main/App";
import {reactive} from "vue";
import {Plus, RefreshRight} from "@element-plus/icons-vue";
import { getCurrentInstance } from 'vue'


const currentRow = ref()
const singleTableRef = ref()
const data = reactive({
  allKeyTypes:[],
  value:""
})
const tagType={
  string:"",
  list:"success",
  set:"info",
  zset:"warning",
  hash:"danger",
}
const types=[
  {text:"string",value:"string"},
  {text:"list",value:"list"},
  {text:"set",value:"set"},
  {text:"zset",value:"zset"},
  {text:"hash",value:"hash"},

]


// 监听表格点击
const handleCurrentChange = (val) => {
  currentRow.value = val
  console.log("点击",val)
  sendBrotherData(val)
}


// 兄弟组件传值
let { proxy } = getCurrentInstance()
function sendBrotherData(value) {
  // 通过暴露info  传递 brother 的值
  proxy.$mitt.emit('key', value)
}




// 标签筛选
const filterTag = (value, row) => {
  return row.type === value
}

// 搜索
const search = ref('')
const filterTableData = computed(() =>
    data.allKeyTypes.filter(
        (data) => !search.value || data.key.toLowerCase().includes(search.value.toLowerCase())
    )
)

function showAll(){
  AllKeyTypes().then(result => {
    data.allKeyTypes = result.data
  })
  console.log("请求全部key")
}
showAll()
// setInterval(showAll , 1000);


</script>