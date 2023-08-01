// Purpose: Export the IdentityInfo interface.

//store中存的个人信息
export interface userInfo {
	id: number,
	username: string,
	sex: string,
	phone_num: string,
	major: string
}

//新建联系人
export interface contacterInfo {
	owner_id: number,
	name: string,
	sex: string | null,
	student_id: string | null,
	phone_num: string,
	major: string | null,
	blacklist: boolean
}

export interface signInfo {
	confirm_password: string;
	major: string;
	password: string;
	phone_num: string;
	sex: string;
	username: string;
}