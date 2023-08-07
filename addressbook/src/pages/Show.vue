<script setup lang="ts">
import showContacts from "../components/showContacter.vue";
</script>

<template>
  <el-container>
    <el-col :span="8">
      <div>
        <el-image class="img" src="src/assets/Add.png">
        </el-image>
      </div>
    </el-col>
    <el-col :span="16">
      <showContacts />
    </el-col>
  </el-container>
</template>

<style scoped>
.img {
  position: relative;
  left: 60px;
  top: -50px;
}
@media (max-width: 1010px) {
  .img {
    display: none;
  }
}
</style>