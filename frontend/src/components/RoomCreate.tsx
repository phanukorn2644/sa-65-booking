import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import SaveIcon from '@mui/icons-material/Save';
import { Room_typeInterface } from "../interfaces/IRoom_type";
import { Room_priceInterface } from "../interfaces/IRoom_price";
import { Set_of_furnitureInterface } from "../interfaces/ISet_of_furniture";
import { RoomInterface } from "../interfaces/IRoom";
import { FurnitureInterface } from "../interfaces/IFurniture";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import {
  GetRoom_type,
  GetRoom_price,
  GetSet_of_furniture,
  CreateRooms,
  GetFurniture,
  GetRoom,
  
} from "../services/HttpClientServiceR";
import FormHelperText from "@mui/material/FormHelperText";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function RoomCreate() {

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },


    {
      field: "Furniture_type",
      headerName: "เฟอนิเจอร์",
      width: 150,
      valueFormatter: (params) => params.value.ID,
    },
    {
      field: "Set_of_furniture_id",
      headerName: "เซทเฟอนิเจอร์",
      width: 150,
      valueFormatter: (params) => params.value.Set_of_furniture_title,
    },
  ];
  


   const [furnitures, setFurnitures] = useState<FurnitureInterface[]>([]);//
  const [Room_type, setRoom_type] = useState<Room_typeInterface[]>([]);
  const [Room_price, setRoom_price] = useState<Room_priceInterface[]>([]);
  const [Set_of_furniture, setSet_of_furniture] = useState<Set_of_furnitureInterface[]>([]);
  const [room, setRoom] = useState<RoomInterface>({});
 

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof RoomCreate;
    setRoom({
      ...room,
      [name]: event.target.value,
    });

    
  };


  

  const getRoom_type = async () => {
    let res = await GetRoom_type();
    if (res) {
      setRoom_type(res);
    }
  };

  const getRoom_price = async () => {
    let res = await GetRoom_price();
    if (res) {
      setRoom_price(res);
    }
  };

  const getSet_of_furniture = async () => {
    let res = await GetSet_of_furniture();
    if (res) {
      setSet_of_furniture(res);
    }
  };
  const Furnitures = async () => {
    let res = await GetFurniture();
    if (res) {
      setFurnitures(res);
    } 
  };

  //



  useEffect(() => {
    getRoom_type();
    getRoom_price();
    getSet_of_furniture();
    Furnitures();
   
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Room_type_id: convertType(room.Room_type_id),
      Room_price_id: convertType(room.Room_price_id),
      Set_of_furniture_id: convertType(room.Set_of_furniture_id),


    };

    //
    console.log(data);
    let res = await CreateRooms(data);
   
    console.log(res)
    if (res) { 
      setSuccess(true);
    } else {
      setError(true);
    }

  }

  return (
    <Container maxWidth="md"
    >
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          มึงขึ้นเหอะกูใหว้ละ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper
      sx ={{ bgcolor :"#E3E3E3"}}>
        
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูลห้อง
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined" >
              <p>ประเภทห้อง</p>
              <Select
                native
                value={room.Room_type_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Room_type_id",
                  
                  
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกประเภทห้อง
                </option>
                {Room_type.map((item: Room_typeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Room_type_name}
                  </option>
                ))}
              </Select>
              
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined" >
              <p>ราคา</p>
              <Select
                native
                value={room.Room_price_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Room_price_id",
                  
                  
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกราคา
                </option>
                {Room_price.map((item: Room_priceInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Price}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เซทเฟอนิเจอร์</p>
              <Select
                native
                value={room.Set_of_furniture_id + ""}
                onChange={handleChange}
                inputProps={{
                  name: "Set_of_furniture_id",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเซทเฟอนิเจอร์
                </option>
                {Set_of_furniture.map((item: Set_of_furnitureInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Set_of_furniture_title}
                  </option>
                ))}
                
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/Rooms"
              variant="contained"
              color="inherit"
              startIcon={< ArrowBackIcon />}


            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
              startIcon={<SaveIcon />}
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
      
      <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
          
            rows={furnitures }
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>

    </Container>
  );
}

export default RoomCreate;