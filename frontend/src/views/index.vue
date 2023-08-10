<script setup>

import HelloWorld from "../components/HelloWorld.vue";
import {Pointer} from "@element-plus/icons-vue";
import DataList from "../components/DataList.vue";
import DisplayEdit from "../components/DisplayEdit.vue";
</script>

<template>
  <el-container>
    <el-header>
<!--      Header-->
      <div class="header-text">
        <router-link to="/"><el-link  :icon="Back" :underline="false">返回</el-link></router-link>
</div>
    </el-header>

    <el-container>
      <el-aside width="300px"><data-list/></el-aside>
      <el-main>
<!--        <HelloWorld/>-->
      <display-edit></display-edit>
      </el-main>
    </el-container>
    <el-footer>
      <el-link href="https://github.com/PCCCQ" target="_blank" :underline="false">version:1.0.0 @PCCCQ</el-link>

    </el-footer>
  </el-container>
</template>
<script>
import { Back} from '@element-plus/icons-vue'

</script>
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