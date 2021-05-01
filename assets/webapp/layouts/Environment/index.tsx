import React from 'react';
import styled from 'styled-components';
import CssBaseline from '@material-ui/core/CssBaseline';
import Divider from '@material-ui/core/Divider';
import Drawer from '@material-ui/core/Drawer';
import Hidden from '@material-ui/core/Hidden';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import { makeStyles, useTheme, Theme, createStyles } from '@material-ui/core/styles';
import { Repeat, Settings, Speed, Web } from '@material-ui/icons';
import { Link, useParams } from 'react-router-dom';
import useEnvironment from '../../hooks/useEnvironment';
import Layers from '@material-ui/icons/Layers';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import IconButton from '@material-ui/core/IconButton';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
    },
    drawer: {
      [theme.breakpoints.up('sm')]: {
        width: drawerWidth,
        flexShrink: 0,
      },
      backgroundColor: '#0C2461',
    },
    appBar: {
      [theme.breakpoints.up('sm')]: {
        width: `calc(100% - ${drawerWidth}px)`,
        marginLeft: drawerWidth,
      },
    },
    menuButton: {
      marginRight: theme.spacing(2),
      [theme.breakpoints.up('sm')]: {
        display: 'none',
      },
    },
    // necessary for content to be below app bar
    toolbar: theme.mixins.toolbar,
    drawerPaper: {
      width: drawerWidth,
      backgroundColor: "#0C2461"
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
  }),
);

interface EnvironmentLayoutProps {
  children: React.ReactNode;
}

export default function EnvironmentLayout(props: EnvironmentLayoutProps) {
  const { children } = props;
  const classes = useStyles();
  const theme = useTheme();
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const { envId } = useParams();
  const { data: environment } = useEnvironment(envId);

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const StyledListItemIcon = styled(ListItemIcon)`
    color: white;
  `;
  const StyledListItemText = styled(ListItemText)`
    color: white;
  `;

  const Logo = styled.img`
    margin: 24px;
    height: 20px;
  `;

  const envLink = (path) => `/environments/${envId}${path}`;

  console.log(environment);
  const drawer = (
    <div>
      <div className={classes.toolbar}>
        <Logo src="/public/img/logo-dark.svg" />
      </div>
      <Divider />
      <List style={{ backgroundColor: 'rgba(0, 0, 0, 0.2)'}}>
        <ListItem>
          <StyledListItemText primary={environment?.name} />
          <ListItemSecondaryAction>
            <IconButton component={Link} to="/select-environment">
              <Repeat style={{ color: '#fff' }} />
            </IconButton>
          </ListItemSecondaryAction>
        </ListItem>
      </List>
      <Divider />
      <List>
        <ListItem button component={Link} to={envLink('/applications')}>
          <StyledListItemIcon><Web /></StyledListItemIcon>
          <StyledListItemText>Applications</StyledListItemText>
        </ListItem>
        {/* TODO: Add status page */}
        {/* <ListItem button>
          <StyledListItemIcon><Speed /></StyledListItemIcon>
          <StyledListItemText>Status</StyledListItemText>
        </ListItem> */}
        <ListItem button component={Link} to={envLink('/settings')}>
          <StyledListItemIcon><Settings /></StyledListItemIcon>
          <StyledListItemText>Settings</StyledListItemText>
        </ListItem>
      </List>
    </div>
  );
  return (
    <div className={classes.root}>
      <CssBaseline />
      <nav className={classes.drawer} aria-label="mailbox folders">
        <Hidden smUp implementation="css">
          <Drawer
            variant="temporary"
            anchor={theme.direction === 'rtl' ? 'right' : 'left'}
            open={mobileOpen}
            onClose={handleDrawerToggle}
            classes={{
              paper: classes.drawerPaper,
            }}
            ModalProps={{
              keepMounted: true,
            }}
            style={{ backgroundColor: "#0C2461" }}
          >
            {drawer}
          </Drawer>
        </Hidden>
        <Hidden xsDown implementation="css">
          <Drawer
            classes={{
              paper: classes.drawerPaper,
            }}
            variant="permanent"
            open
            style={{ backgroundColor: "#0C2461" }}
          >
            {drawer}
          </Drawer>
        </Hidden>
      </nav>
      <main className={classes.content}>
        {children}
      </main>
    </div>
  );
}
