// Purpose: Export the IdentityInfo interface.
import {PersonalInfo} from "./personalInfo";

export interface IdentityInfo {
	id: number,
	password: string,
	personInfo: PersonalInfo
}