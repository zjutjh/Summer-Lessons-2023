import {defineStore} from "pinia";
import {reactive} from "vue";
import {PersonalInfo} from "../types/personalInfo.ts";

const userStore = defineStore(
	"Id",
	() => {
		const idInfo = reactive({
			id: -1,
			password: "",
		});

		const personalInfo = reactive({
			name: "未登录",
			sex: true,
			phoneNumber: "未登录",
			major: "未登录",
		});

		const setIdInfo = (info: {id:number , password:string}) => {
			idInfo.id = info.id;
			idInfo.password = info.password;
		};
		const setPersonalInfo = (info: PersonalInfo) => {
			personalInfo.name = info.name;
			personalInfo.sex = info.sex;
			personalInfo.phoneNumber = info.phoneNumber;
			personalInfo.major = info.major;
		};

		return {
			idInfo,
			personalInfo,

			setIdInfo,
			setPersonalInfo,
		};
	}
);

export default userStore;