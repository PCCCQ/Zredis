<script setup>
import {defineProps, reactive, toRefs} from "vue";
import {GetHash} from "../../../wailsjs/go/main/App.js";

const props = defineProps({
//子组件接收父组件传递过来的值
  operateKey:""
})
//使用父组件传递过来的值
const {operateKey} =toRefs(props)
let data=reactive({
  key:operateKey.value,
  value:""
})
GetHash({key:"hash"}).then(res=>{
  console.log(res)
  data.value=res.data.valueList
})
</script>

<template>
hash
  {{data.value}}
</template>

<style scoped>

</style>