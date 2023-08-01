<template>
  <div class="foot_div">
    <div>
      <el-divider content-position="center">@精弘网络·2023</el-divider>
    </div>
    <div class="content">
      <el-descriptions :column="4" size="small" >
        <el-descriptions-item label="QQ group"
        >885342047&ensp;&ensp;&ensp;</el-descriptions-item>
        <el-descriptions-item label="Github">
          <el-link href="http://github.com/Tianci-King/Summer-Lessons-2023">
            http://github.com/Tianci-King/Summer-Lessons-2023&ensp;&ensp;&ensp;
          </el-link>
        </el-descriptions-item>
        <el-descriptions-item label="Official website"
        ><el-link href="http://www.jh.zjut.edu.cn/index">http://www.jh.zjut.edu.cn&ensp;&ensp;&ensp;</el-link>
        </el-descriptions-item><el-descriptions-item label="Designed by"
      >Tiancy</el-descriptions-item>

      </el-descriptions>
    </div>
  </div>
</template>

<script setup lang="ts">

</script>

<style scoped>
.foot_div {
  position: absolute;
  left: 2%;
  right: 2%;
  bottom: 2%;
}
</style>