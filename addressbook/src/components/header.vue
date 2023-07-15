<script setup lang="ts">

</script>

<template>
  <el-header>
    <el-row>
      <el-col :span="8">
        <div class="logo">
          <img src="../assets/JH.png" alt="logo" />
        </div>
      </el-col>
      <el-col :span="16">
        <div class="user-info">
          <span>欢迎您</span>
          <el-button type="primary" @click="logout">退出</el-button>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>

<style scoped>

</style>