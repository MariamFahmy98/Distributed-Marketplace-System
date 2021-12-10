import React from "react";
import { useDispatch, useSelector } from "react-redux";

import { makeStyles } from "@material-ui/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CardMedia from "@material-ui/core/CardMedia";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Link from "@material-ui/core/Link";
import ShoppingCartOutlinedIcon from "@material-ui/icons/ShoppingCartOutlined";
import Typography from "@material-ui/core/Typography";

import { decrementBalance } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { setForm } from "../state/features/forms";
import { showSnackbar } from "../state/features/snackbar";
import { useOrderProductMutation } from "../state/service";

const useStyles = makeStyles((theme) => ({
  paper: {
    margin: theme.spacing(2, 0, 2, 0),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  cover: {
    marginTop: theme.spacing(1),
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(2, 0),
  },
}));

function OrderProductDialog() {
  const classes = useStyles();

  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);
  const form = useSelector((state) => state.forms.order);

  const [orderProduct] = useOrderProductMutation();

  async function handleOrderProduct() {
    if (
      current_user.balance === null ||
      current_user.balance < parseFloat(form.price)
    ) {
      dispatch(
        showSnackbar({
          variant: "warning",
          message: "Not enough balance.",
        })
      );
      return;
    }

    try {
      const data = await orderProduct(form.id).unwrap();
      if (!data.success) {
        dispatch(
          showSnackbar({
            variant: "error",
            message: data.error
              ? data.error
              : "An error ocurred please try again later!",
          })
        );
      } else {
        dispatch(
          showSnackbar({
            variant: "success",
            message: "Order placed",
          })
        );
        dispatch(setDialog(null));
        dispatch(setForm({ form: "order", values: null }));
        localStorage.setItem(
          "user",
          JSON.stringify({
            ...current_user,
            balance: current_user.balance - parseFloat(form.price),
          })
        );
        dispatch(decrementBalance(parseFloat(form.price)));
      }
    } catch (err) {
      dispatch(
        showSnackbar({
          variant: "error",
          message: err.data.error && err.data ? err.data.error : err.status,
        })
      );
    }
  }

  return (
    <Container className={classes.paper} maxWidth="xs">
      <Avatar className={classes.avatar}>
        <ShoppingCartOutlinedIcon />
      </Avatar>
      <Typography component="h1" variant="h5">
        {form && form.name ? form.name : "Product Name"}
      </Typography>
      <CardMedia
        component="img"
        className={classes.cover}
        image={
          form && form.image_url
            ? form.image_url
            : "https://images.unsplash.com/photo-1638467611417-c5437577fcd4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80"
        }
        loading="auto"
        title={form && form.name ? form.name : "Product Name"}
      />
      <Grid className={classes.form} container spacing={2}>
        <Grid item xs={6}>
          <Typography style={{ textAlign: "center" }}>
            Balance: {current_user.balance !== null ? current_user.balance : 0}
          </Typography>
        </Grid>
        <Grid item xs={6}>
          <Typography style={{ textAlign: "center" }}>{`Price: ${
            form && form.price ? form.price : 0
          }`}</Typography>
        </Grid>
      </Grid>
      <Button
        fullWidth
        variant="outlined"
        color="primary"
        className={classes.submit}
        onClick={handleOrderProduct}
      >
        Order
      </Button>
      <Grid container justify="flex-end">
        <Grid item>
          <Link
            component="button"
            variant="body2"
            onClick={() => {
              dispatch(setDialog("deposit-money"));
            }}
          >
            Not enough balance? Top Up
          </Link>
        </Grid>
      </Grid>
    </Container>
  );
}

export default OrderProductDialog;
