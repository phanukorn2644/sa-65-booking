import { EmployeeInterface } from "../interfaces/IEmployee";
import { GendersInterface } from "../interfaces/IGender";
import { ProvincesInterface } from "../interfaces/IProvince";
import { ProgramInterface } from "../interfaces/IProgram";
import { RoleInterface } from "../interfaces/IRole";

export interface StudentInterface {
  ID?: number;
  STUDENT_NUMBER?: string;
  STUDENT_NAME? : string;
  PERSONAL_ID? : string;
  Password?: string;
  GenderID?: number;
  Gender?: GendersInterface;
  ProvinceID?: number;
  Province?: ProvincesInterface;
  ProgramID?: number;
  Program?: ProgramInterface;
  RoleID?: number;
  Role?: RoleInterface;
  EmployeeID? : number;
  Employee? : EmployeeInterface;
}
