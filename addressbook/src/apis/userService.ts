import request from "./request";
import {signInfo} from "../types/personalInfo.ts";

export default class userService {
	static async login(data: {
		phone_num: string,
		password: string
	}) {
		return request({
			"headers": {
				"Content-Type": "application/json",
			},
			method: "post",
			url: "/api/login",
			data: data
		});
	}

	static async sign(data: signInfo) {
		return request({
			"headers": {
				"Content-Type": "application/json",
			},
			method: "post",
			url: "/api/register",
			data: data
		});
	};
}