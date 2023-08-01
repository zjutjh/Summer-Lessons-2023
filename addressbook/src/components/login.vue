<template>
  <div class="div">
    <el-container>
      <label class="label">Account&nbsp;&nbsp;&nbsp;</label>
      <el-input v-model="input" placeholder="请输入您的账号" class="input"/>
    </el-container>
    <br />
    <el-container>
      <label class="label">Password</label>
      <el-input
          v-model="password"
          type="password"
          placeholder="请输入您的密码"
          show-password
          class="input"
      />
    </el-container>
    <br />
    <el-button type="primary" @click="login" class="button">登录</el-button>
    <el-button @click="clear" class="button">清空</el-button>
  </div>
</template>

<script setup lang="ts">
import { ref , h } from "vue";
import userService from "../apis/userService.ts";
import loginStore from "../stores/loginStore.ts";
import userStore from "../stores/userStore.ts";
import {ElNotification} from "element-plus";
import router from "../routers";

const input = ref("");
const password = ref("");
const newLoginStore = loginStore();
const newUserStore = userStore();

const login = async () => {
  const loginInfo = ref({
    phone_num: input.value,
    password: password.value,
  });

  const res = await userService.login(loginInfo.value);

if (res.data.msg === "OK" && res.data.code === 200) {
  const responseData = res.data.data;
  const message = "亲爱的" + responseData.username + ",欢迎回来！";
  ElNotification({
    title: "登陆成功！",
    message: h("i", { style: "color: teal" }, message),
  });

  newLoginStore.setLogin(true);
  console.log("登录状态"+newLoginStore.loginSession);
  localStorage.setItem("login", String(true));
  localStorage.setItem("name", String(responseData.username));
  newUserStore.setUserInfo({
    id: responseData.id,
    username: responseData.username,
    sex: responseData.sex,
    phone_num: responseData.phone_num,
    major: responseData.major
  });
  router.push("/Add");
}
else {
  alert("sorry!");
}

};
const clear = () => {
  input.value = "";
  password.value = "";
};
</script>

<style scoped>
.div {
  width: 400px;
  height: 300px;
  margin: 200px 130px 0 auto;
}

.label {
  margin-right: 10px;
  font-weight: bolder;
}

.button {
  margin: 20px;
  width: 100px;
}
</style>
