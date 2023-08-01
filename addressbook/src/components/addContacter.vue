<script setup lang="ts">
import { reactive , h } from "vue";
import addService from "../apis/addService.ts";
import userStore from "../stores/userStore.ts";
import { ElNotification } from "element-plus";

const newUserStore = userStore();
// do not use same name with ref
const form = reactive({
  owner_id: newUserStore.userSession.id,
  name: "",
  sex: "",
  student_id: "",
  phone_num: "",
  major: "",
  blacklist: false
});

const onSubmit = async () => {
  if (form.name !== "" && form.phone_num !== "") {
    const res = await addService.add(form);

    if (res.data.code === 200 && res.data.msg === "OK") {
      console.log(form);
      ElNotification({
        title: "成功",
        message: h("i", { style: "color: teal" }, "新建联系人成功！"),
      });
      clear();
    }
    else {
      ElNotification({
        title: "失败",
        message: h("i", { style: "color: teal" }, "新建联系人失败，请检查网络！"),
      });
    }
  }
  else {
    ElNotification({
      title: "失败",
      message: h("i", { style: "color: teal" }, "新建联系人失败，请至少输入完整的姓名和电话号码！"),
    });
  }
};

const clear = () => {
  form.name = "";
  form.sex = "";
  form.student_id = "";
  form.phone_num = "";
  form.major = "";
  form.blacklist = false;
};
</script>

<template>
  <div class="div">
    <div style="text-align: center ; margin-left: 120px">
      <h3>添加联系人</h3>
    </div>
    <el-form :model="form" label-width="120px">
      <el-form-item label="姓名">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="性别">
        <el-select v-model="form.sex" placeholder="请选择TA的性别">
          <el-option label="男" value= "男" />
          <el-option label="女" value= "女" />
        </el-select>
      </el-form-item>
      <el-form-item label="学号">
        <el-input v-model="form.student_id" />
      </el-form-item>
      <el-form-item label="电话号码">
        <el-input v-model="form.phone_num" />
      </el-form-item>
      <el-form-item label="就读专业">
        <el-input v-model="form.major" type="textarea" />
      </el-form-item>
      <el-form-item label="拉黑名单">
        <el-switch v-model="form.blacklist" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit" class="button">新建</el-button>
        <el-button class="button" @click="clear">清空</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped>
.div {
  width: 400px;
  height: 300px;
  margin: 40px 130px 0 auto;
}

.button {
  margin: 20px;
  width: 100px;
}
</style>