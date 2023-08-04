import {defineStore} from "pinia";
import {reactive} from "vue";
import {userInfo} from "../types/personalInfo.ts";

const userStore = defineStore(
	"user",
	() => {
		const userSession = reactive({
			id: -1,
			username: "未登录",
			sex: "未登录",
			phone_num: "未登录",
			major: "未登录",
		});

		const setUserInfo = (info: userInfo) : void => {
			userSession.id = info.id;
			userSession.username = info.username;
			userSession.sex = info.sex;
			userSession.phone_num = info.phone_num;
			userSession.major = info.major;
		};

		return {
			userSession,
			setUserInfo,
		};
	},
	{
		persist: true
	}
);

export default userStore;