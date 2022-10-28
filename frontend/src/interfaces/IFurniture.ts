import { Set_of_furnitureInterface } from "./ISet_of_furniture";
export interface FurnitureInterface {

    ID: number,
	Furniture_type: string
   
    Set_of_furniture_id?: number;
    Set_of_furniture?: Set_of_furnitureInterface;
   
   }