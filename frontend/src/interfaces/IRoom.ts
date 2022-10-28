import { Room_typeInterface } from "./IRoom_type";
import { Room_priceInterface} from "./IRoom_price";
import { Set_of_furnitureInterface } from "./ISet_of_furniture";

export interface RoomInterface {
  ID?: number;

  Room_type_id?: number;
 Room_type?: Room_typeInterface;
  Room_price_id?: number;
 Room_price?: Room_priceInterface;
  Set_of_furniture_id?: number;
  Set_of_furniture?: Set_of_furnitureInterface;
}