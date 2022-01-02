import React, { useEffect, useRef } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";

import { makeStyles, useTheme, withStyles } from "@material-ui/core/styles";
import AccountBalanceIcon from "@material-ui/icons/AccountBalance";
import Box from "@material-ui/core/Box";
import Card from "@material-ui/core/Card";
import CardMedia from "@material-ui/core/CardMedia";
import CardContent from "@material-ui/core/CardContent";
import Chip from "@material-ui/core/Chip";
import EditIcon from "@material-ui/icons/Edit";
import IconButton from "@material-ui/core/IconButton";
import Tab from "@material-ui/core/Tab";
import Tabs from "@material-ui/core/Tabs";
import Typography from "@material-ui/core/Typography";
import useMediaQuery from "@material-ui/core/useMediaQuery";

import optimizeImage from "../utils/image";
import Avatar from "../assets/images/avatar-0.png";
import StyledButton from "../components/Button";
import Products from "../components/Products";
import Report from "../components/Report";

import { changePicture } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { showSnackbar } from "../state/features/snackbar";
import {
  useGetUserBalanceQuery,
  useGetUserProductsQuery,
  useGetUserPurchasedProductsQuery,
  useGetUserSoldProductsQuery,
  useGetUserTransactionsQuery,
  useGetUserOrdersQuery,
  useEditUserMutation,
} from "../state/service";

function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`vertical-tabpanel-${index}`}
      aria-labelledby={`vertical-tab-${index}`}
      {...other}
    >
      {value === index && <Box m={2}>{children}</Box>}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired,
};

const AntTabs = withStyles({
  root: {
    borderBottom: "1px solid #e8e8e8",
  },
  indicator: {
    backgroundColor: "#1890ff",
  },
})(Tabs);

const AntTab = withStyles((theme) => ({
  root: {
    textTransform: "none",
    minWidth: 72,
    fontWeight: theme.typography.fontWeightRegular,
    marginRight: theme.spacing(4),
    fontFamily: [
      "-apple-system",
      "BlinkMacSystemFont",
      '"Segoe UI"',
      "Roboto",
      '"Helvetica Neue"',
      "Arial",
      "sans-serif",
      '"Apple Color Emoji"',
      '"Segoe UI Emoji"',
      '"Segoe UI Symbol"',
    ].join(","),
    "&:hover": {
      color: "#40a9ff",
      opacity: 1,
    },
    "&$selected": {
      color: "#1890ff",
      fontWeight: theme.typography.fontWeightMedium,
    },
    "&:focus": {
      color: "#40a9ff",
    },
  },
  selected: {},
}))((props) => <Tab disableRipple {...props} />);

function a11yProps(index) {
  return {
    id: `vertical-tab-${index}`,
    "aria-controls": `vertical-tabpanel-${index}`,
  };
}

const useStyles = makeStyles(({ breakpoints, palette, spacing }) => ({
  root: {
    display: "flex",
    flexGrow: 1,
  },
  tabs: {
    borderRight: `1px solid ${palette.divider}`,
  },
  relative: { position: "relative" },
  card: {
    margin: "auto",
    overflow: "initial",
    position: "relative",
    boxShadow: "none",
    borderRadius: 0,
    "&:hover": {
      "& .MuiTypography--explore": {
        transform: "scale(1.2)",
      },
    },
  },
  cardCover: {
    [breakpoints.only("xs")]: { height: 240 },
    [breakpoints.only("sm")]: { height: 360 },
    [breakpoints.up("md")]: { height: 480 },
    objectFit: "cover",
  },
  cardContent: {
    boxShadow: "0 16px 40px -12.125px rgba(0,0,0,0.3)",
    borderRadius: spacing(0.5),
    backgroundColor: "#ffffff",
    position: "absolute",
    top: "70%",
    left: "2%",
    width: "96%",
    padding: `${spacing(3, 3, 1, 3)} !important`,
    textAlign: "left",
  },
  innerCard: {
    display: "flex",
  },
  details: {
    display: "flex",
    flexDirection: "column",
    width: "100%",
  },
  content: {
    flex: "1 0 auto",
  },
  photo: {
    [breakpoints.only("xs")]: { width: spacing(15), height: spacing(15) },
    [breakpoints.up("sm")]: { width: spacing(20), height: spacing(20) },
    objectFit: "cover",
    borderRadius: "50%",
  },
  photoEdit: {
    position: "absolute",
    bottom: 0,
    right: 0,
  },
  photoEditIcon: {
    background: "GhostWhite",
    margin: spacing(1),
  },
  controls: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    width: "100%",
    paddingLeft: spacing(1),
    paddingBottom: spacing(1),
  },
  input: {
    display: "none",
  },
}));

function Profile() {
  const classes = useStyles();
  const theme = useTheme();

  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);

  const cardContent = useRef(null);
  const [step, setStep] = React.useState(0);
  const [cardContentHeight, setCardContentHeight] = React.useState(0);
  const [pictureState, setPictureState] = React.useState(
    current_user.image_url
  );

  const matchesSM = useMediaQuery(theme.breakpoints.down("sm"));
  const matchesXS = useMediaQuery(theme.breakpoints.down("xs"));

  let heightOffset = Math.round(0.3 * 480);
  if (matchesSM) heightOffset = Math.round(0.3 * 360);
  if (matchesXS) heightOffset = Math.round(0.3 * 240);

  const {
    data: userBalanceData,
    error: userBalanceError,
    isLoading: userBalanceIsLoading,
  } = useGetUserBalanceQuery(undefined, {
    skip: current_user.id === null,
  });

  const {
    data: userProductsData,
    error: userProductsError,
    isLoading: userProductsIsLoading,
  } = useGetUserProductsQuery(current_user.id, {
    skip: current_user.id === null,
  });
  const {
    data: userPurchasedProductsData,
    error: userPurchasedProductsError,
    isLoading: userPurchasedProductsIsLoading,
  } = useGetUserPurchasedProductsQuery(undefined, {
    skip: current_user.id === null,
  });
  const {
    data: userSoldProductsData,
    error: userSoldProductsError,
    isLoading: userSoldProductsIsLoading,
  } = useGetUserSoldProductsQuery(undefined, {
    skip: current_user.id === null,
  });
  const {
    data: userTransactionsData,
    error: userTransactionsError,
    isLoading: userTransactionsIsLoading,
  } = useGetUserTransactionsQuery(undefined, {
    skip: current_user.id === null,
  });
  const {
    data: userOrdersData,
    error: userOrdersError,
    isLoading: userOrdersIsLoading,
  } = useGetUserOrdersQuery(undefined, {
    skip: current_user.id === null,
  });

  const [editUser] = useEditUserMutation();

  useEffect(() => {
    setCardContentHeight(cardContent.current.offsetHeight);
  }, [cardContent]);

  const handleStepChange = (_event, newValue) => {
    setStep(newValue);
  };

  const handlePictureChange = async (event) => {
    const files = event.target.files;
    if (files.length > 0) {
      setPictureState(URL.createObjectURL(files[0]));

      const optimizedImage = await optimizeImage(files[0]);
      try {
        const data = await editUser({
          image_url: optimizedImage,
        }).unwrap();
        dispatch(changePicture(data.data.image_url));
        localStorage.setItem(
          "user",
          JSON.stringify({ ...current_user, image_url: data.data.image_url })
        );
        dispatch(
          showSnackbar({
            variant: "success",
            message: "Profile Picture changed successfully",
          })
        );
      } catch (err) {
        dispatch(
          showSnackbar({
            variant: "error",
            message: err.data.error && err.data ? err.data.error : err.status,
          })
        );
      }
    }
  };

  const handleDeposit = () => {
    dispatch(setDialog("deposit-money"));
  };

  const tabs = [
    {
      name: "Inventory",
      children: userProductsIsLoading ? (
        <div>Loading...</div>
      ) : userProductsError ? (
        <div>
          {userProductsError.data && userProductsError.data.error
            ? userProductsError.data.error
            : userProductsError.status}
        </div>
      ) : userProductsData &&
        userProductsData.data &&
        userProductsData.data.length > 0 ? (
        <Products products={userProductsData.data} />
      ) : (
        <div>No data.</div>
      ),
    },
    {
      name: "Purchased Products",
      children: userPurchasedProductsIsLoading ? (
        <div>Loading...</div>
      ) : userPurchasedProductsError ? (
        <div>
          {userPurchasedProductsError.data &&
          userPurchasedProductsError.data.error
            ? userPurchasedProductsError.data.error
            : userPurchasedProductsError.status}
        </div>
      ) : userPurchasedProductsData &&
        userPurchasedProductsData.data &&
        userPurchasedProductsData.data.length > 0 ? (
        <Products products={userPurchasedProductsData.data} />
      ) : (
        <div>No data.</div>
      ),
    },
    {
      name: "Sold Products",
      children: userSoldProductsIsLoading ? (
        <div>Loading...</div>
      ) : userSoldProductsError ? (
        <div>
          {userSoldProductsError.data && userSoldProductsError.data.error
            ? userSoldProductsError.data.error
            : userSoldProductsError.status}
        </div>
      ) : userSoldProductsData &&
        userSoldProductsData.data &&
        userSoldProductsData.data.length > 0 ? (
        <Products products={userSoldProductsData.data} />
      ) : (
        <div>No data.</div>
      ),
    },
    {
      name: "Transactions Report",
      children: userTransactionsIsLoading ? (
        <div>Loading...</div>
      ) : userTransactionsError ? (
        <div>
          {userTransactionsError.data && userTransactionsError.data.error
            ? userTransactionsError.data.error
            : userTransactionsError.status}
        </div>
      ) : userTransactionsData &&
        userTransactionsData.data &&
        userTransactionsData.data.length > 0 ? (
        <Report
          title="Transactions"
          headers={["Amount", "Timestamp"]}
          rows={userTransactionsData.data.map((el) => [
            el.amount,
            new Date(el.created_at).toLocaleString(),
          ])}
        />
      ) : (
        <div>No data.</div>
      ),
    },
    {
      name: "Orders Report",
      children: userOrdersIsLoading ? (
        <div>Loading...</div>
      ) : userOrdersError ? (
        <div>
          {userOrdersError.data && userOrdersError.data.error
            ? userOrdersError.data.error
            : userOrdersError.status}
        </div>
      ) : userOrdersData &&
        userOrdersData.data &&
        userOrdersData.data.length > 0 ? (
        <Report
          title="Orders"
          headers={[
            "Buyer Id",
            "Buyer Name",
            "Seller Id",
            "Seller Name",
            "Product",
            "Price",
            "Date",
          ]}
          rows={userOrdersData.data.map((el) => [
            el.buyer.id,
            el.buyer.name,
            el.seller.id,
            el.seller.name,
            el.product.title,
            el.product.price,
            new Date(el.created_at).toLocaleString(),
          ])}
        />
      ) : (
        <div>No data.</div>
      ),
    },
  ];

  return (
    <React.Fragment>
      <Card
        style={{
          marginBottom: cardContentHeight - heightOffset + theme.spacing(2.5),
        }}
        className={classes.card}
      >
        <CardMedia
          component="img"
          className={classes.cardCover}
          src="https://images.unsplash.com/photo-1638467611417-c5437577fcd4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80"
          alt="Cover"
        />
        <CardContent ref={cardContent} className={classes.cardContent}>
          <Card className={classes.innerCard} elevation={0}>
            <div className={classes.relative}>
              <CardMedia
                component="img"
                className={classes.photo}
                image={
                  pictureState
                    ? pictureState
                    : current_user.image_url
                    ? current_user.image_url
                    : Avatar
                }
                loading="auto"
                title={current_user.name ? current_user.name : "Profile"}
              />
              {current_user.id !== null && (
                <React.Fragment>
                  <input
                    accept="image/*"
                    className={classes.input}
                    id="photo-edit-button"
                    type="file"
                    onChange={handlePictureChange}
                  />
                  <label
                    className={classes.photoEdit}
                    htmlFor="photo-edit-button"
                  >
                    <IconButton
                      className={classes.photoEditIcon}
                      aria-label="edit-photo"
                      component="span"
                    >
                      <EditIcon />
                    </IconButton>
                  </label>
                </React.Fragment>
              )}
            </div>
            <div className={classes.details}>
              <CardContent className={classes.content}>
                <Typography component="h1" variant="h5">
                  {current_user.name ? current_user.name : "Profile"}
                </Typography>
                <Chip
                  variant="outlined"
                  size="small"
                  icon={<AccountBalanceIcon />}
                  label={`Balance: ${
                    !userBalanceIsLoading &&
                    !userBalanceError &&
                    userBalanceData &&
                    userBalanceData.data
                      ? userBalanceData.data
                      : 0
                  }`}
                  color="primary"
                />
              </CardContent>
              <div className={classes.controls}>
                <StyledButton color="red" onClick={handleDeposit}>
                  Deposit
                </StyledButton>
              </div>
            </div>
          </Card>
          <AntTabs
            style={{ width: "100%" }}
            value={step}
            onChange={handleStepChange}
            aria-label="profile sections"
            variant={matchesSM ? "scrollable" : "standard"}
            scrollButtons="on"
            centered={!matchesSM}
          >
            {tabs.map((tab, index) => (
              <AntTab key={tab.name} label={tab.name} {...a11yProps(index)} />
            ))}
          </AntTabs>
        </CardContent>
      </Card>
      <div className={classes.root}>
        {tabs.map((tab, index) => (
          <TabPanel
            key={tab.name}
            style={{ width: "100%" }}
            value={step}
            index={index}
          >
            {tab.children}
          </TabPanel>
        ))}
      </div>
    </React.Fragment>
  );
}

export default Profile;
