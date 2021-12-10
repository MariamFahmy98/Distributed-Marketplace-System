import React, { useEffect } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate, useLocation } from "react-router-dom";
import clsx from "clsx";

import { makeStyles, useTheme } from "@material-ui/core/styles";
import AccountCircleTwoTone from "@material-ui/icons/AccountCircleTwoTone";
import Avatar from "@material-ui/core/Avatar";
import AppBar from "@material-ui/core/AppBar";
import Container from "@material-ui/core/Container";
import Dialog from "@material-ui/core/Dialog";
import Divider from "@material-ui/core/Divider";
import Drawer from "@material-ui/core/Drawer";
import FaceIcon from "@material-ui/icons/Face";
import Hidden from "@material-ui/core/Hidden";
import HomeIcon from "@material-ui/icons/Home";
import IconButton from "@material-ui/core/IconButton";
import KeyboardArrowDownRoundedIcon from "@material-ui/icons/KeyboardArrowDownRounded";
import KeyboardArrowLeftRoundedIcon from "@material-ui/icons/KeyboardArrowLeftRounded";
import KeyboardArrowRightRoundedIcon from "@material-ui/icons/KeyboardArrowRightRounded";
import Link from "@material-ui/core/Link";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import Menu from "@material-ui/core/Menu";
import MenuIcon from "@material-ui/icons/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import Snackbar from "@material-ui/core/Snackbar";
import StorefrontIcon from "@material-ui/icons/Storefront";
import SwipeableDrawer from "@material-ui/core/SwipeableDrawer";
import Toolbar from "@material-ui/core/Toolbar";
import Tooltip from "@material-ui/core/Tooltip";
import Typography from "@material-ui/core/Typography";
import useMediaQuery from "@material-ui/core/useMediaQuery";
import useScrollTrigger from "@material-ui/core/useScrollTrigger";

import AddProductDialog from "../dialogs/AddProductDialog";
import DepositMoneyDialog from "../dialogs/DepositMoneyDialog";
import Logo from "../assets/images/logo192.png";
import OrderProductDialog from "../dialogs/OrderProductDialog";
import SearchBar from "./SearchBar";
import SignUpDialog from "../dialogs/SignUpDialog";
import SignInDialog from "../dialogs/SignInDialog";
import SnackbarContentWrapper from "./SnackbarContentWrapper";

import { setUser } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { setMenuAnchor } from "../state/features/menus";
import { setSnackbar, showSnackbar } from "../state/features/snackbar";

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
  },
  grow: {
    flexGrow: 1,
  },
  appBar: {
    backgroundColor: theme.palette.background.default,
    [theme.breakpoints.up("sm")]: {
      // zIndex: theme.zIndex.drawer + 1,
      marginLeft: theme.spacing(9) + 1,
      width: `calc(100% - ${theme.spacing(9) + 1}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  hide: {
    display: "none",
  },
  logoWrapper: {
    display: "flex", // contents is buggy with some major browsers
  },
  logo: {
    width: theme.spacing(3.5),
    margin: theme.spacing(0, 1.75, 0, 0),
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  logoHide: {
    width: 0,
    margin: 0,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
    whiteSpace: "nowrap",
  },
  drawerHeader: {
    marginTop: theme.spacing(1),
  },
  swipableDrawer: {
    borderBottomLeftRadius: 0,
    borderBottomRightRadius: 0,
  },
  drawerOpen: {
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerClose: {
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    overflowX: "hidden",
    width: theme.spacing(9) + 1,
  },
  toolbar: {
    padding: theme.spacing(0, 0, 0, 2),
    ...theme.mixins.toolbar,
  },
  inline: {
    display: "inline",
  },
  content: {
    flexGrow: 1,
    display: "flex",
    flexDirection: "column",
    minHeight: "100vh",
  },
  main: {
    padding: theme.spacing(2),
  },
  footer: {
    padding: theme.spacing(2),
    marginTop: "auto",
    borderTop: `1px solid ${theme.palette.divider}`,
    backgroundColor: theme.palette.background.paper,
  },
  facebookIcon: {
    verticalAlign: "middle",
    marginLeft: theme.spacing(1),
  },
  sponsorIcon: {
    verticalAlign: "middle",
    height: theme.spacing(6),
  },
}));

function ElevationScroll(props) {
  const { children } = props;
  const trigger = useScrollTrigger({
    disableHysteresis: true,
    threshold: 0,
  });

  return React.cloneElement(children, {
    elevation: trigger ? 4 : 0,
  });
}

ElevationScroll.propTypes = {
  children: PropTypes.element.isRequired,
};

function Dialogs() {
  const dispatch = useDispatch();
  const dialog = useSelector((state) => state.dialog.value);

  const handleDialogClose = () => {
    dispatch(setDialog(null));
  };

  return (
    <React.Fragment>
      <Dialog
        open={dialog === "sign-in"}
        onClose={handleDialogClose}
        aria-labelledby="sign-in-dialog"
      >
        <SignInDialog />
      </Dialog>
      <Dialog
        open={dialog === "sign-up"}
        onClose={handleDialogClose}
        aria-labelledby="sign-up-dialog"
      >
        <SignUpDialog />
      </Dialog>
      <Dialog
        open={dialog === "add-product"}
        onClose={handleDialogClose}
        aria-labelledby="add-product-dialog"
      >
        <AddProductDialog />
      </Dialog>
      <Dialog
        open={dialog === "deposit-money"}
        onClose={handleDialogClose}
        aria-labelledby="deposit-money-dialog"
      >
        <DepositMoneyDialog />
      </Dialog>
      <Dialog
        open={dialog === "order-product"}
        onClose={handleDialogClose}
        aria-labelledby="order-product-dialog"
      >
        <OrderProductDialog />
      </Dialog>
    </React.Fragment>
  );
}

const RootLevelDialogs = Dialogs;

function Snackbars() {
  const dispatch = useDispatch();
  const snackbar = useSelector((state) => state.snackbar);

  const handleSnackbarClose = (event_, reason) => {
    if (reason === "clickaway") return;
    dispatch(setSnackbar({ open: false }));
  };

  const processQueue = () => {
    if (snackbar.queue.length > 0) {
      setSnackbar(true, snackbar.queue.shift());
    }
  };

  const handleSnackbarExited = () => {
    processQueue();
  };

  return (
    <Snackbar
      key={snackbar.messageInfo ? snackbar.messageInfo.key : undefined}
      anchorOrigin={{
        vertical: "bottom",
        horizontal: "center",
      }}
      open={snackbar.open}
      autoHideDuration={6000}
      onClose={handleSnackbarClose}
      onExited={handleSnackbarExited}
      ContentProps={{
        "aria-describedby": "message-id",
      }}
    >
      <SnackbarContentWrapper
        onClose={handleSnackbarClose}
        variant={
          snackbar.messageInfo ? snackbar.messageInfo.variant : undefined
        }
        message={
          snackbar.messageInfo ? snackbar.messageInfo.message : undefined
        }
        actionLabel={
          snackbar.messageInfo ? snackbar.messageInfo.actionLabel : undefined
        }
        onActionClick={
          snackbar.messageInfo ? snackbar.messageInfo.action : undefined
        }
      />
    </Snackbar>
  );
}

const RootLevelSnackbars = Snackbars;

function DrawerInfo() {
  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);

  useEffect(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      dispatch(setUser(JSON.parse(storedUser)));
    }

    return () => {
      if (current_user.id !== null)
        localStorage.setItem("user", JSON.stringify(current_user));
    };
    // eslint-disable-next-line
  }, []);

  return (
    <React.Fragment>
      <ListItemAvatar>
        {current_user.image_url ? (
          <Avatar alt="user avatar" src={current_user.image_url} />
        ) : (
          <AccountCircleTwoTone fontSize="large" />
        )}
      </ListItemAvatar>
      <ListItemText
        primary="Welcome!"
        secondary={current_user.name ? current_user.name : "Anon"}
      />
    </React.Fragment>
  );
}

const DrawerHeader = DrawerInfo;

const accountMenuId = "account-menu";

function MainMenu() {
  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);
  const accountMenu = useSelector((state) => state.menus.account);

  const isAccountMenuOpen = Boolean(accountMenu);

  const handleMenuClose = (menu) => () => {
    dispatch(setMenuAnchor({ menu, anchor: null }));
  };

  const handleAccountClick = () => {
    if (current_user.id === null) {
      dispatch(setDialog("sign-in"));
    } else {
      localStorage.removeItem("user");
      dispatch(setUser(null));
      // TODO invalidate cache
      dispatch(showSnackbar({ variant: "success", message: "Signed out" }));
    }
    handleMenuClose("account")();
  };

  return (
    <Menu
      anchorEl={accountMenu}
      anchorOrigin={{ vertical: "top", horizontal: "right" }}
      id={accountMenuId}
      keepMounted
      transformOrigin={{ vertical: "top", horizontal: "right" }}
      open={isAccountMenuOpen}
      onClose={handleMenuClose("account")}
    >
      <MenuItem onClick={handleAccountClick}>
        {current_user.id === null ? "Sign in" : "Sign out"}
      </MenuItem>
    </Menu>
  );
}

const AccountMenu = MainMenu;

const accountMenuStyles = makeStyles((theme) => ({
  avatar: {
    width: theme.spacing(3),
    height: theme.spacing(3),
  },
}));

function MainMenuTrigger() {
  const classes = accountMenuStyles();

  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);

  const handleAccountMenuOpen = (event) => {
    dispatch(setMenuAnchor({ menu: "account", anchor: event.currentTarget }));
  };

  return (
    <IconButton
      aria-label="account of current user"
      aria-controls={accountMenuId}
      aria-haspopup="true"
      onClick={handleAccountMenuOpen}
      color="inherit"
    >
      {current_user.image_url ? (
        <Avatar
          className={classes.avatar}
          alt="user avatar"
          src={current_user.image_url}
        />
      ) : (
        <AccountCircleTwoTone />
      )}
    </IconButton>
  );
}

const AccountMenuTrigger = MainMenuTrigger;

const drawerItemsStyles = makeStyles((theme) => ({
  listItem: {
    margin: theme.spacing(1),
    width: "auto",
    borderRadius: theme.spacing(1),
  },
}));

function DrawerItems() {
  const classes = drawerItemsStyles();

  const navigate = useNavigate();
  const location = useLocation();

  const items = [
    {
      id: "/",
      icon: <HomeIcon />,
      text: "Home",
    },
    {
      id: "/profile",
      icon: <FaceIcon />,
      text: "Profile",
    },
    {
      id: "/store",
      icon: <StorefrontIcon />,
      text: "My Store",
    },
  ];

  return (
    <List>
      {items.map((item) => (
        <Tooltip key={item.id} title={item.text} arrow placement="right">
          <ListItem
            className={classes.listItem}
            selected={
              item.id === "/"
                ? location.pathname === item.id
                : location.pathname.startsWith(item.id)
            }
            button
            onClick={() => navigate(item.id)}
          >
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItem>
        </Tooltip>
      ))}
    </List>
  );
}

export default function MainNav(props) {
  const classes = useStyles();

  const theme = useTheme();

  const [mobileOpen, setMobileOpen] = React.useState(false);
  const [open, setOpen] = React.useState(false);
  const [searchFocus, setSearchFocus] = React.useState(false);

  const isDesktop = useMediaQuery(theme.breakpoints.up("sm"));
  const iOS = process.browser && /iPad|iPhone|iPod/.test(navigator.userAgent);

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const handleDrawerOn = () => {
    if (isDesktop) {
      handleDrawerOpen();
    } else {
      handleDrawerToggle();
    }
  };

  const handleDrawerOff = () => {
    if (isDesktop) {
      handleDrawerClose();
    } else {
      handleDrawerToggle();
    }
  };

  const handleSearchFocus = (focused) => () => {
    setSearchFocus(focused);
  };

  const handleProductSelected = () => {
    return;
  };

  const drawer = (
    <React.Fragment>
      <ListItem
        className={classes.drawerHeader}
        ContainerComponent="div"
        alignItems="flex-start"
      >
        <DrawerHeader />
        {(open || mobileOpen) && (
          <ListItemSecondaryAction>
            <IconButton
              edge="end"
              aria-label="close drawer"
              onClick={handleDrawerOff}
            >
              {isDesktop ? (
                theme.direction === "rtl" ? (
                  <KeyboardArrowRightRoundedIcon />
                ) : (
                  <KeyboardArrowLeftRoundedIcon />
                )
              ) : (
                <KeyboardArrowDownRoundedIcon />
              )}
            </IconButton>
          </ListItemSecondaryAction>
        )}
      </ListItem>
      <Divider variant="middle" />
      <DrawerItems />
    </React.Fragment>
  );

  function Copyright() {
    return (
      <Typography variant="body2" color="textSecondary" align="center">
        {"Copyright © "}
        <Link color="inherit" href="/">
          Distributed Marketplace
        </Link>{" "}
        {new Date().getFullYear()}
      </Typography>
    );
  }

  return (
    <React.Fragment>
      <div className={clsx(classes.root, classes.grow)}>
        <ElevationScroll>
          <AppBar
            position="fixed"
            className={clsx(classes.appBar, {
              [classes.appBarShift]: open,
            })}
            color="inherit"
          >
            <Toolbar className={classes.toolbar}>
              <IconButton
                color="inherit"
                aria-label="open drawer"
                onClick={handleDrawerOn}
                edge="start"
                className={clsx({
                  [classes.hide]: open,
                })}
              >
                <MenuIcon />
              </IconButton>
              <Link className={classes.logoWrapper} href="/">
                <img
                  className={clsx(classes.logo, {
                    [classes.logoHide]: searchFocus,
                  })}
                  src={Logo}
                  alt="logo"
                />
              </Link>
              <SearchBar
                id="products-search"
                placeholder="Search products"
                onFocus={handleSearchFocus(true)}
                onBlur={handleSearchFocus(false)}
                handleSuggestionSelected={handleProductSelected}
              />
              <AccountMenuTrigger />
            </Toolbar>
          </AppBar>
        </ElevationScroll>
        <AccountMenu />
        <nav aria-label="site pages">
          <Hidden smUp implementation="css">
            <SwipeableDrawer
              anchor="bottom"
              open={mobileOpen}
              onClose={handleDrawerToggle}
              ModalProps={{
                keepMounted: true, // Better open performance on mobile.
              }}
              PaperProps={{
                square: false,
              }}
              classes={{
                paper: classes.swipableDrawer,
              }}
              disableBackdropTransition={!iOS}
              disableSwipeToOpen={false}
              onOpen={handleDrawerToggle}
            >
              {drawer}
            </SwipeableDrawer>
          </Hidden>
          <Hidden xsDown implementation="css">
            <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open,
                  [classes.drawerClose]: !open,
                }),
              }}
              open={open}
            >
              {drawer}
            </Drawer>
          </Hidden>
        </nav>
        <div className={classes.content}>
          <div className={classes.toolbar} />
          <Container className={classes.main} component="main" maxWidth="md">
            {props.children}
          </Container>
          <footer className={classes.footer}>
            <Copyright />
          </footer>
        </div>
      </div>
      <RootLevelDialogs />
      <RootLevelSnackbars />
    </React.Fragment>
  );
}

MainNav.propTypes = {
  children: PropTypes.node.isRequired,
};
