import { h } from "vue";
import {createRouter , createWebHistory } from "vue-router";
import Login from "../pages/Login.vue";
import Add from "../pages/Add.vue";
import Show from "../pages/Show.vue";
// 分析原因：因为在main.ts中，注册router总比pinia先，所以不能使用到store/index.js文件中createPinia方法，只能在router文件中再createPinia一次，才能使用到pinia。
import Sign from "../pages/Sign.vue";
import {ElNotification} from "element-plus";

const routes = [
	{
		path: "/" ,
		redirect: "/Login"
	},
	{
		path: "/Login" ,
		name: "Login" ,
		component: Login
	},
	{
		path: "/Sign" ,
		name: "Sign" ,
		component: Sign
	},
	{
		path: "/Add" ,
		name: "Add" ,
		component: Add
	},
	{
		path: "/Show" ,
		name: "Show" ,
		component: Show
	},

];

const router = createRouter (
	{
		history: createWebHistory() ,
		routes: routes
	}
);

// 路由守卫
// router.beforeEach((to, from,next) => {
//   // ...
//   // 返回 false 以取消导航
//   return false;
// });
router.beforeEach((to, _, next) => {
	//loginSession 是登录的标记，true表示已经登录，false表示未登录
	if (localStorage.getItem("login") === "false") {
		//解决无限重定向的问题
		if (to.path === "/Login") {
			next();
		}
		else {
			ElNotification({
				title: "失败",
				message: h("i", { style: "color: teal" }, "请您先登录！"),
			});
			next("/Login");
		}
	}
	else {
		next();
	}
});


export default router;