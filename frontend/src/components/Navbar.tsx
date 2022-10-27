import React from "react";
// import { makeStyles } from "@mui/material/styles";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import { Link } from "react-router-dom";

// const useStyles = makeStyles((theme) => ({
//   root: {
//     flexGrow: 1,
//   },
//   menuButton: {
//     marginRight: theme.spacing(2),
//   },
//   title: {
//     flexGrow: 1,
//   },
//   navlink: {
//     color: "white",
//     textDecoration: "none",
//   },
// }));

function Navbar() {
  // const classes = useStyles();
  return (
    <div  style={{flexGrow: 1}}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            sx={{marginRight: 2}}
            color="inherit"
            aria-label="menu"
          >
            <MenuIcon />
          </IconButton>
          <Link style={{color: "white",textDecoration: "none",marginRight:"auto"}} to="/">
            <Typography variant="h6" sx={{ flexGrow: 1 }}>
            ระบบขี้นทะเบียนผู้ป่วยโรงพยาบาลมหาวิทยาลัยเทคโนโลยีสุรนารี
            </Typography>
          </Link>
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default Navbar;