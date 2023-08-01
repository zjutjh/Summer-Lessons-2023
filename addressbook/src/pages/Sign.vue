<template>
  <el-row>
    <el-col :span="8">
      <div>
        <el-image class="img" src="src/assets/JHcartoon.png">
        </el-image>
      </div>
    </el-col>
    <el-col :span="16">
      <sign />
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import sign from "../components/sign.vue";
</script>

<style scoped>
.img {
  position: relative;
  left: 100px;
  top: -30px;
}
@media (max-width: 1010px) {
  .img {
    display: none;
  }
}
</style>