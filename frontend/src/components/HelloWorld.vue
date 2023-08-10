<script setup>
import {reactive} from 'vue'
import {SetString, Greet, RedisInit, DelString, GetString, AllKeyTypes,Del} from '../../wailsjs/go/main/App'
const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
  result:{},
  string:{key:"", value:"",expiration:-1,},
  stringResult:{},
  set:{key:"", value:"",expiration:-1,},
  redisInit:{}
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
    RedisInit({IP:"127.0.0.1",port:"6379"}).then(result => {
      data.redisInit = result
    })

  })
}
greet()
function showAll(){
  AllKeyTypes().then(result => {
    data.allKeyTypes = result
  })
}


function setString() {
  data.expiration=data.expiration+0
  SetString(data.string).then(result => {
    data.stringResult = result
  })
}

function getString() {
  GetString(data.string).then(result => {
    data.stringResult = result
  })
}

function delString() {
  Del(data.string.key).then(result => {
      data.stringResult = result
    })
}


</script>

<template>
  <main>
    <div id="result" class="result">redis连接：{{ data.redisInit }}</div>
    <div>{{data.allKeyTypes}}</div>
    <button class="btn" @click="showAll">显示所有</button>
    <div id="input" class="input-box">
      string:
      key<input id="name" v-model="data.string.key" autocomplete="off" class="input" type="text" />
      value<input id="name" v-model="data.string.value" autocomplete="off" class="input" type="text" />
      expiration<input id="name" v-model="data.string.expiration" autocomplete="off" class="input" type="number" />
      <button class="btn" @click="setString">创建</button>
      <button class="btn" @click="getString">获取</button>
      <button class="btn" @click="delString">删除</button>
      <div>{{data.stringResult}}</div>
    </div>


  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
