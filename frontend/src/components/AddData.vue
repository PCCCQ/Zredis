<template>
  <el-dialog v-model="isShow" title="新增数据">
    <el-form>
      <el-form-item label="Type">
        <el-select v-model="data.type" placeholder="请选择数据类型">
          <el-option v-for="item in types" :label="item.text" :value="item.value"/>
        </el-select>
      </el-form-item>
      <el-form-item label="Key">
        <el-input v-model="data.key"/>
      </el-form-item>


      <el-form-item label="Value" v-if="data.type==='string'">
        <el-input v-model="data.value"/>
      </el-form-item>
      <el-form-item label="Value" v-if="data.type==='list'">
        <el-input v-model="data.value"/>
      </el-form-item>
      <el-form-item label="Value" v-if="data.type==='set'">
        <el-input v-model="data.value"/>
      </el-form-item>
      <template v-if="data.type==='zset'">
        <el-form-item label="Score">
          <el-input-number v-model.number="data.objectValue.Score"/>
        </el-form-item>
        <el-form-item label="Member">
          <el-input v-model="data.objectValue.Member"/>
        </el-form-item>
      </template>
      <template v-if="data.type==='hash'">
        <el-form-item label="key">
          <el-input v-model="data.objectValue.key"/>
        </el-form-item>
        <el-form-item label="value">
          <el-input v-model="data.objectValue.value"/>
        </el-form-item>
      </template>


    </el-form>


    <template #footer>
      <span class="dialog-footer">
        <el-button @click="isShow = false">取消</el-button>
        <el-button type="primary" @click="addDataFunc">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import {computed, reactive, ref} from 'vue'
import {ElMessage} from "element-plus";
import {AddZset, SetHash, SetList, SetSet, SetString} from "../../wailsjs/go/main/App.js";
const isShow = ref(false)
const types = [
  {text: "string", value: "string"},
  {text: "list", value: "list"},
  {text: "set", value: "set"},
  {text: "zset", value: "zset"},
  {text: "hash", value: "hash"},

]
const data = reactive({
  type: "",
  key: "",
  value: "",
  objectValue: {}
})
defineExpose({isShow})


function addString() {
  SetString(data).then(res => {
    if (res.code === 200) {
      ElMessage.success(res.data)

    } else {
      ElMessage.error(res.data)
    }
  })
}

function addList() {
  SetList(data).then(res => {
    if (res.code === 200) {
      ElMessage.success(res.data)
    } else {
      ElMessage.error(res.data)
    }
  })
}

function addSet() {
  SetSet(data).then(res => {
    if (res.code === 200) {
      ElMessage.success(res.data)
    } else {
      ElMessage.error(res.data)
    }
  })
}

function addZset() {
  AddZset({key:data.key,value:data.objectValue}).then(res => {
    if (res.code === 200) {
      ElMessage.success(res.data)
    } else {
      ElMessage.error(res.data)
    }
  })
}

function addHash() {
  SetHash({key:data.key,valueList:[data.objectValue]}).then(res => {
    if (res.code === 200) {
      ElMessage.success(res.data)
    } else {
      ElMessage.error(res.data)
    }
  })
}
const parent = computed(() => {
  return this.$parent;
});

function addDataFunc() {
  if(data.key=== ""||data.value=== ""||data.type=== ""||data.objectValue=== {}){
    ElMessage.error("请填写完整")
    return
  }
  switch (data.type) {
    case "string":
      addString()
      break;
    case "list":
      addList()
      break;
    case "set":
      addSet()
      break;
    case "zset":
      addZset()
      break;
    case "hash":
      addHash()
      break;
    default:
      ElMessage.error("错误")
      break;
  }
  console.log("执行到这")
  isShow.value=false
  parent.changeTime()

}


</script>

<style scoped>

</style>