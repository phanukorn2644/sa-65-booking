import { EmployeeInterface } from "../interfaces/IEmployee";
import { GenderInterface } from "../interfaces/IGender";
import { ProvinceInterface } from "../interfaces/IProvince";
import { ProgramInterface } from "../interfaces/IProgram";
import { RoleInterface } from "../interfaces/IRole";

export interface StudentInterface {
  ID?: number;
  STUDENT_NUMBER?: string;
  STUDENT_NAME? : string;
  PERSONAL_ID? : string;
  Password?: string;
  GenderID?: number;
  Gender?: GenderInterface;
  ProvinceID?: number;
  Province?: ProvinceInterface;
  ProgramID?: number;
  Program?: ProgramInterface;
  RoleID?: number;
  Role?: RoleInterface;
  EmployeeID? : number;
  Employee? : EmployeeInterface;
}
