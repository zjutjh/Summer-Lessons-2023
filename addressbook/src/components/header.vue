<script setup lang="ts">
import { ref } from "vue";
import router from "../routers";

const isLogin = localStorage.getItem("login");
console.log(isLogin);
const isName = localStorage.getItem("name");
const activeIndex = ref("1");
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

const pushToLogin = () => {
  router.push("/Login");
};

const pushToSign = () => {
  router.push("/Sign");
};

const pushToAdd = () => {
  router.push("/Add");
};

const pushToShow = () => {
  router.push("/Show");
};
</script>

<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo menu"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
  >
    <el-menu-item index="0"><div class="flex items-center">
      <span class="text-large font-600 mr-3"> 通讯录管理系统 </span>
      &ensp;&ensp;&ensp;
      <el-tag>精弘网络</el-tag>
    </div></el-menu-item>
    <div class="flex-grow" />
    <el-sub-menu index="2">
      <template #title>个人中心</template>
      <el-menu-item index="2-1" @click="pushToAdd">新建联系人</el-menu-item>
      <el-menu-item index="2-2" @click="pushToShow">查看联系人</el-menu-item>
    </el-sub-menu>
    <div class="button_div">
      <div class="button" v-if="!isLogin">
        <el-button type="primary" @click="pushToLogin" >登录</el-button>
        <el-button class="ml-2" @click="pushToSign" >注册</el-button>
      </div>
      <div v-else>
        <p>亲爱的{{isName}},欢迎回来</p>
      </div>
    </div>
  </el-menu>
</template>

<style scoped>
.flex-grow {
  flex-grow: 1;
}

.menu {
  position: fixed;
  left: 2%;
  right: 2%;
  top: 2%;
}

.button_div {
  position: relative;
}

.button {
  margin-top: 12px;
}
</style>