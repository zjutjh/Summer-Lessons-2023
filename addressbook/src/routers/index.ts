import {createRouter , createWebHistory } from "vue-router";
import Login from "../components/login.vue";
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
		path: "/Add" ,
		name: "Add" ,
		component:
	}
];

const router = createRouter (
	{
		history: createWebHistory() ,
		routes: routes
	}
);

export default router;