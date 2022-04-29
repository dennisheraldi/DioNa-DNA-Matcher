import * as React from 'react';
import { makeStyles } from '@material-ui/core'
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';

import MenuIcon from '@mui/icons-material/Menu';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';

import {ContentPasteSearchOutlined, HistoryOutlined, AddBoxOutlined} from '@mui/icons-material';

import { useHistory, useLocation } from 'react-router-dom'

const drawerWidth = 240;
const useStyles = makeStyles((theme) => {
  return {
    page: {
      // background: '#f9f9f9',
      width: '100%',
      padding: theme.spacing(3),
    },
    toolbar: theme.mixins.toolbar,
  }
})
function Layout(props) {
  const classes = useStyles();
  const { window, children } = props;
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const history = useHistory();

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const buttonList = [
    {
      name: 'DNA Checkup',
      icon: <ContentPasteSearchOutlined color="secondary"/>,
      path: '/',
    },
    {
      name: 'Add Penyakit',
      icon: <AddBoxOutlined color="secondary" />,
      path: '/add-penyakit',
    },
    {
      name: 'Checkup History',
      icon: <HistoryOutlined color="secondary" />,
      path: '/check-history',
    }
  ]

  const drawer = (
    <div>
      <Toolbar />
      <Divider />
      <List>
        {buttonList.map((elmt, index) => (
          <ListItem button key={elmt.name} onClick={() => {handleDrawerToggle(); history.push(elmt.path);}}>
            <ListItemIcon>
              {elmt.icon}
            </ListItemIcon>
            <ListItemText primary={elmt.name} />
          </ListItem>
        ))}
      </List>
    </div>
  );

  const container = window !== undefined ? () => window().document.body : undefined;

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        color="primary"
        sx={{
          width: { sm: `calc(100% - ${drawerWidth}px)` },
          ml: { sm: `${drawerWidth}px` }
        }}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="start"
            onClick={handleDrawerToggle}
            sx={{ mr: 2, display: { sm: 'none' }, textAlign: "center", height: 60 }}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h5" noWrap component="div" sx={{fontWeight: 'medium'}} >
            DioNA
          </Typography>
        </Toolbar>
      </AppBar>
      <Box
        component="nav"
        sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
        aria-label="Navigation Button"
      >
        <Drawer
          container={container}
          variant="temporary"
          open={mobileOpen}
          onClose={handleDrawerToggle}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
          sx={{
            display: { xs: 'block', sm: 'none' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
          }}
        >
          {drawer}
        </Drawer>
        <Drawer
          variant="permanent"
          sx={{
            display: { xs: 'none', sm: 'block' },
            '& .MuiDrawer-paper': { boxSizing: 'border-box', width: drawerWidth },
          }}
          open
        >
          {drawer}
        </Drawer>
      </Box>
      <Box
        component="main"
        sx={{ flexGrow: 1, p: 1, width: { sm: `calc(100% - ${drawerWidth}px)` } }}
        className="allbg">
        {/* <CssBaseline /> */}
        <div className={classes.page}>
          <div className={classes.toolbar}></div>
          { children }
        </div>
      </Box>
    </Box>
  );
}

// ResponsiveDrawer.propTypes = {
//   /**
//    * Injected by the documentation to work in an iframe.
//    * You won't need it on your project.
//    */
//   window: PropTypes.func,
// };

export default Layout;