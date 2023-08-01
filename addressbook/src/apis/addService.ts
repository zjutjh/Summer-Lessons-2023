import request from "./request.ts";
import {contacterInfo} from "../types/personalInfo.ts";

export default class addService {
	static async add (data: contacterInfo) {
		return request({
			"headers": {
				"Content-Type": "application/json",
			},
			url:"/api/contact",
			method: "post",
			data: data
		});
	}

	static async show (data: {
		owner_id: number
	}) {
		return request({
			"headers": {
				"Content-Type": "application/json",
			},
			url:"/api/contact",
			method: "get",
			data:data
		});
	}
}