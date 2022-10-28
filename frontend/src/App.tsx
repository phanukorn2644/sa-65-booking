import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { styled, createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import MuiDrawer from "@mui/material/Drawer";
import Box from "@mui/material/Box";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import IconButton from "@mui/material/IconButton";
import Container from "@mui/material/Container";
import MenuIcon from "@mui/icons-material/Menu";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import Button from "@mui/material/Button";

import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import HomeIcon from "@mui/icons-material/Home";
import PeopleIcon from "@mui/icons-material/People";
import YouTubeIcon from "@mui/icons-material/YouTube";
import FolderIcon from '@mui/icons-material/Folder';
import ConstructionIcon from '@mui/icons-material/Construction';
import ReceiptLongRoundedIcon from '@mui/icons-material/ReceiptLongRounded';



import Home from "./components/Home"
import SignIn from "./components/SignIn";
import BookingCreate from "./components/BookingCreate";
import Booking from "./components/Booking";
import Employees from "./components/Employees";
import EmployeeCreate from "./components/EmployeeCreate";
import Room from "./components/Room";
import RoomCreate from "./components/RoomCreate"
import Repair from "./components/Repair";
import RepairCreate from "./components/RepairCreate";
import Students from "./components/Student";
import StudentCreate from "./components/StudentCreate";
import Payment_Bill from "./components/Payment_Bill";
import Payment_Bill_Create from "./components/Payment_Bill_Create";

const drawerWidth = 240;

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== "open",
})<AppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(["width", "margin"], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== "open",
})(({ theme, open }) => ({
  "& .MuiDrawer-paper": {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
    boxSizing: "border-box",
    ...(!open && {
      overflowX: "hidden",
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9),
      },
    }),
  },
}));

const mdTheme = createTheme();

const menu = [
  { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },
  { name: "จองห้องพัก", icon: <PeopleIcon />, path: "/Booking" , position: "Student"},
  { name: "พนักงาน", icon: <PeopleIcon />, path: "/employees" , position: "Admin"},
  { name: "บันทึกข้อมูลห้อง", icon: <FolderIcon />, path: "/Rooms" , position: "Admin"},
  { name: "แจ้งซ่อม", icon: <ConstructionIcon />, path: "/Repair" , position: "Student"},
  { name: "ลงทะเบียนนักศึกษา", icon: <PeopleIcon />, path: "/students" ,position: "Admin"},
  { name: "ใบเสร็จชำระค่าใช้จ่าย", icon: <ReceiptLongRoundedIcon />, path: "/payment_bills",position: "Admin" },
];

function App() {
  const [token, setToken] = useState<String>("");
  const [open, setOpen] = React.useState(true);
  const [position, setPosition] = useState<String | null>("");
  const toggleDrawer = () => {
    setOpen(!open);
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    const position = localStorage.getItem("position");
    if (token) {
      setToken(token);
      
      setPosition(position);
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  return (
    <Router>
      <ThemeProvider theme={mdTheme}>
        <Box sx={{ display: "flex" }}>
          <CssBaseline />
          <AppBar position="absolute" open={open}>
            <Toolbar
              sx={{
                pr: "24px", // keep right padding when drawer closed
              }}
            >
              <IconButton
                edge="start"
                color="inherit"
                aria-label="open drawer"
                onClick={toggleDrawer}
                sx={{
                  marginRight: "36px",
                  ...(open && { display: "none" }),
                }}
              >
                <MenuIcon />
              </IconButton>
              <Typography
                component="h1"
                variant="h6"
                color="inherit"
                noWrap
                sx={{ flexGrow: 1 }}
              >
                System Analysis and Design 1/65
              </Typography>
              <Button color="inherit" onClick={signout}>
                ออกจากระบบ
              </Button>
            </Toolbar>
          </AppBar>
          <Drawer variant="permanent" open={open}>
            <Toolbar
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "flex-end",
                px: [1],
              }}
            >
              <IconButton onClick={toggleDrawer}>
                <ChevronLeftIcon />
              </IconButton>
            </Toolbar>
            <Divider />
            <List>
              {menu.map((item, index) => 
              position === item.position &&(
                <Link
                  to={item.path}
                  key={item.name}
                  style={{ textDecoration: "none", color: "inherit" }}
                >
                  <ListItem button>
                    <ListItemIcon>{item.icon}</ListItemIcon>
                    <ListItemText primary={item.name} />
                  </ListItem>
                </Link>
              ))}
            </List>
          </Drawer>
          <Box
            component="main"
            sx={{
              backgroundColor: (theme) =>
                theme.palette.mode === "light"
                  ? theme.palette.grey[100]
                  : theme.palette.grey[900],
              flexGrow: 1,
              height: "100vh",
              overflow: "auto",
            }}
          >
            <Toolbar />
            <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/Booking" element={<Booking />} />
                <Route path="/Booking/create" element={<BookingCreate />} />
                <Route path="/employees" element={<Employees />} />
                <Route path="/employee/create" element={<EmployeeCreate />} />
                <Route path="/Rooms" element={<Room />} />
                <Route path="/Room/create" element={<RoomCreate/>}/>
                <Route path="/Repair" element={<Repair />} />
                <Route path="/Repair/create" element={<RepairCreate />} />
                <Route path="/students" element={<Students />} />
                <Route path="/student/create" element={<StudentCreate />} />
                <Route path="/payment_bills" element={<Payment_Bill />}/>
                <Route path="/payment_bills/create"element={<Payment_Bill_Create />}/>
                
              </Routes>
            </Container>
          </Box>
        </Box>
      </ThemeProvider>
    </Router>
  );
}

export default App;