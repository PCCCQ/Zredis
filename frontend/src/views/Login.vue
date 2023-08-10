<template>
<div class="father">
  <div class="form">
    <v-form-render :form-json="formJson" :form-data="formData" :option-data="optionData" ref="vFormRef">
    </v-form-render>
    <div style="text-align: center">
<!--    <el-button type="primary" @click="submitForm">确定</el-button>-->
    <el-button type="primary" @click="submitForm">测试连接</el-button>
    <el-button type="primary" @click="router.push('/')">返回主页面</el-button>
    </div>
  </div>
</div>
</template>




<style>
body {
  margin: 0; /* 如果页面出现垂直滚动条，则加入此行CSS以消除之 */
}
.father{
  display: flex;
  flex-direction: column;
  justify-content: center;/*新版本*/
  align-items: center; /* 添加这一行来垂直居中内容 */
  height: 100vh; /* 添加这一行以确保容器占据整个视口高度 */
}
.form {
  min-width: 400px;
  width: 50%;
  margin:auto;

}
</style>

<script setup>
import {reactive, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {RedisInit} from '../../wailsjs/go/main/App.js'
import router from "../router/index.js";

const formJson = reactive({
  "widgetList": [{
    "key": 52468,
    "type": "input",
    "icon": "text-field",
    "formItemFlag": true,
    "options": {
      "name": "name",
      "label": "连接名称",
      "labelAlign": "",
      "type": "text",
      "defaultValue": "",
      "placeholder": "请为此连接命名",
      "columnWidth": "200px",
      "size": "",
      "labelWidth": null,
      "labelHidden": false,
      "readonly": false,
      "disabled": false,
      "hidden": false,
      "clearable": true,
      "showPassword": false,
      "required": true,
      "requiredHint": "连接标识符不可为空",
      "validation": "",
      "validationHint": "",
      "customClass": [],
      "labelIconClass": null,
      "labelIconPosition": "rear",
      "labelTooltip": null,
      "minLength": null,
      "maxLength": null,
      "showWordLimit": false,
      "prefixIcon": "",
      "suffixIcon": "",
      "appendButton": false,
      "appendButtonDisabled": false,
      "buttonIcon": "custom-search",
      "onCreated": "",
      "onMounted": "",
      "onInput": "",
      "onChange": "",
      "onFocus": "",
      "onBlur": "",
      "onValidate": "",
      "onAppendButtonClick": ""
    },
    "id": "input57123"
  }, {
    "type": "input",
    "icon": "text-field",
    "formItemFlag": true,
    "options": {
      "name": "IP",
      "label": "IP地址",
      "labelAlign": "",
      "type": "text",
      "defaultValue": "",
      "placeholder": "默认为localhost",
      "columnWidth": "200px",
      "size": "",
      "labelWidth": null,
      "labelHidden": false,
      "readonly": false,
      "disabled": false,
      "hidden": false,
      "clearable": true,
      "showPassword": false,
      "required": false,
      "validation": "",
      "validationHint": "",
      "customClass": [],
      "labelIconClass": null,
      "labelIconPosition": "rear",
      "labelTooltip": null,
      "minLength": null,
      "maxLength": null,
      "showWordLimit": false,
      "prefixIcon": "",
      "suffixIcon": "",
      "appendButton": false,
      "appendButtonDisabled": false,
      "buttonIcon": "el-icon-search",
      "onCreated": "",
      "onMounted": "",
      "onInput": "",
      "onChange": "",
      "onFocus": "",
      "onBlur": "",
      "onValidate": ""
    },
    "id": "input12931"
  }, {
    "type": "input",
    "icon": "text-field",
    "formItemFlag": true,
    "options": {
      "name": "port",
      "label": "端口号",
      "labelAlign": "",
      "type": "text",
      "defaultValue": "",
      "placeholder": "默认为6379",
      "columnWidth": "200px",
      "size": "",
      "labelWidth": null,
      "labelHidden": false,
      "readonly": false,
      "disabled": false,
      "hidden": false,
      "clearable": true,
      "showPassword": false,
      "required": false,
      "validation": "number",
      "validationHint": "端口号格式错误",
      "customClass": [],
      "labelIconClass": null,
      "labelIconPosition": "rear",
      "labelTooltip": null,
      "minLength": null,
      "maxLength": null,
      "showWordLimit": false,
      "prefixIcon": "",
      "suffixIcon": "",
      "appendButton": false,
      "appendButtonDisabled": false,
      "buttonIcon": "el-icon-search",
      "onCreated": "",
      "onMounted": "",
      "onInput": "",
      "onChange": "",
      "onFocus": "",
      "onBlur": "",
      "onValidate": ""
    },
    "id": "input23031"
  }, {
    "type": "input",
    "icon": "text-field",
    "formItemFlag": true,
    "options": {
      "name": "userName",
      "label": "用户名",
      "labelAlign": "",
      "type": "text",
      "defaultValue": "",
      "placeholder": "",
      "columnWidth": "200px",
      "size": "",
      "labelWidth": null,
      "labelHidden": false,
      "readonly": false,
      "disabled": false,
      "hidden": false,
      "clearable": true,
      "showPassword": false,
      "required": false,
      "validation": "",
      "validationHint": "",
      "customClass": [],
      "labelIconClass": null,
      "labelIconPosition": "rear",
      "labelTooltip": null,
      "minLength": null,
      "maxLength": null,
      "showWordLimit": false,
      "prefixIcon": "",
      "suffixIcon": "",
      "appendButton": false,
      "appendButtonDisabled": false,
      "buttonIcon": "el-icon-search",
      "onCreated": "",
      "onMounted": "",
      "onInput": "",
      "onChange": "",
      "onFocus": "",
      "onBlur": "",
      "onValidate": ""
    },
    "id": "input113152"
  }, {
    "type": "input",
    "icon": "text-field",
    "formItemFlag": true,
    "options": {
      "name": "password",
      "label": "密码",
      "labelAlign": "",
      "type": "text",
      "defaultValue": "",
      "placeholder": "",
      "columnWidth": "200px",
      "size": "",
      "labelWidth": null,
      "labelHidden": false,
      "readonly": false,
      "disabled": false,
      "hidden": false,
      "clearable": true,
      "showPassword": false,
      "required": false,
      "validation": "",
      "validationHint": "",
      "customClass": [],
      "labelIconClass": null,
      "labelIconPosition": "rear",
      "labelTooltip": null,
      "minLength": null,
      "maxLength": null,
      "showWordLimit": false,
      "prefixIcon": "",
      "suffixIcon": "",
      "appendButton": false,
      "appendButtonDisabled": false,
      "buttonIcon": "el-icon-search",
      "onCreated": "",
      "onMounted": "",
      "onInput": "",
      "onChange": "",
      "onFocus": "",
      "onBlur": "",
      "onValidate": ""
    },
    "id": "input40240"
  }],
  "formConfig": {
    "modelName": "formData",
    "refName": "vForm",
    "rulesName": "rules",
    "labelWidth": 87,
    "labelPosition": "left",
    "size": "",
    "labelAlign": "label-right-align",
    "cssCode": "",
    "customClass": [],
    "functions": "",
    "layoutType": "PC",
    "jsonVersion": 3,
    "onFormCreated": "",
    "onFormMounted": "",
    "onFormDataChange": "",
    "onFormValidate": ""
  }
})
const formData = reactive({})
const optionData = reactive({})
const vFormRef = ref(null)
const data = reactive({

})
const submitForm = () => {
  vFormRef.value.getFormData().then(formData => {
    // Form Validation OK
    if (!formData.IP) {
      formData.IP = "localhost"
    }
    if (!formData.port) {
      formData.port = "6379"
    }
    // console.log("连接测试1",formData)
    RedisInit(formData).then(result => {
      console.log("连接测试2")
      if(result.code===200){
        ElMessage.success(result.data)
        let tmp=[]
        if (window.localStorage.getItem("redisList")!=null){
          tmp= eval('(' + window.localStorage.getItem("redisList") + ')');
        }
        tmp.push({name:formData.name,IP:formData.IP,port:formData.port,userName:formData.userName,password:formData.password})
        // console.log("formData",tmp.push(formData),tmp)
        console.log("formData",JSON.stringify(tmp))
        window.localStorage.setItem("redisList",JSON.stringify(tmp))
        router.push("/index")
      }else {
        ElMessage.error(result.data)
      }

    })
  }).catch(error => {
    // Form Validation failed
    ElMessage.error(error)
  })
}


</script>