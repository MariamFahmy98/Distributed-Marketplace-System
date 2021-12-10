import React, { useState } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import NumberFormat from "react-number-format";

import { makeStyles } from "@material-ui/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import FormControl from "@material-ui/core/FormControl";
import Grid from "@material-ui/core/Grid";
import InputLabel from "@material-ui/core/InputLabel";
import InputAdornment from "@material-ui/core/InputAdornment";
import OutlinedInput from "@material-ui/core/OutlinedInput";
import PaymentOutlinedIcon from "@material-ui/icons/PaymentOutlined";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";

import { incrementBalance } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { showSnackbar } from "../state/features/snackbar";
import { useDepositMoneyMutation } from "../state/service";

function limit(val, max) {
  if (val.length === 1 && val[0] > max[0]) {
    val = "0" + val;
  }

  if (val.length === 2) {
    if (Number(val) === 0) {
      val = "01";

      //this can happen when user paste number
    } else if (val > max) {
      val = max;
    }
  }

  return val;
}

const formatCustomPropTypes = {
  inputRef: PropTypes.func.isRequired,
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func.isRequired,
};

function CreditCardFormatCustom(props) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      format="####-####-####-####"
    />
  );
}

CreditCardFormatCustom.propTypes = formatCustomPropTypes;

function CurrencyFormatCustom(props) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      thousandSeparator
      isNumericString
    />
  );
}

CurrencyFormatCustom.propTypes = formatCustomPropTypes;

function MonthFormatCustom(props) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      format={(val) => limit(val.substring(0, 2), "12")}
    />
  );
}

MonthFormatCustom.propTypes = formatCustomPropTypes;

function YearFormatCustom(props) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      format={(val) => limit(val.substring(0, 2), "99")}
    />
  );
}

YearFormatCustom.propTypes = formatCustomPropTypes;

function CvvFormatCustom(props) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      format={(val) => limit(val.substring(0, 3), "999")}
    />
  );
}

CvvFormatCustom.propTypes = formatCustomPropTypes;

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
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(4),
  },
  google: {
    backgroundColor: "#4285f4",
    color: theme.palette.getContrastText("#4285f4"),
    marginBottom: theme.spacing(2),
    [theme.breakpoints.down("xs")]: {
      marginBottom: 0,
    },
  },
  submit: {
    margin: theme.spacing(2, 0),
  },
}));

function DepositMoneyDialog() {
  const classes = useStyles();

  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);
  const orderForm = useSelector((state) => state.forms.order);

  const [dirty, setDirty] = useState(false);
  const [amount, setAmount] = useState("0");
  const [number, setNumber] = useState("");
  const [expirationMonth, setExpirationMonth] = useState("");
  const [expirationYear, setExpirationYear] = useState("");
  const [holderName, setHolderName] = useState("");
  const [cvv, setCvv] = useState("");

  const [depositMoney] = useDepositMoneyMutation();

  async function handleDeposit() {
    if (!dirty) setDirty(true);

    if (
      amount.length === 0 ||
      parseFloat(amount) <= 0 ||
      number.length !== 16 ||
      expirationMonth.length !== 2 ||
      expirationYear.length !== 2 ||
      holderName.length === 0 ||
      cvv.length !== 3
    ) {
      dispatch(
        showSnackbar({
          variant: "warning",
          message: "Please fill in the data correctly",
        })
      );
      return;
    }

    try {
      const data = await depositMoney(parseFloat(amount)).unwrap();
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
            message: `Your balance now is: ${data.balance}`,
          })
        );
        if (orderForm) dispatch(setDialog("order-product"));
        else dispatch(setDialog(null));
        localStorage.setItem(
          "user",
          JSON.stringify({
            ...current_user,
            balance: data.balance,
          })
        );
        dispatch(incrementBalance(parseFloat(amount)));
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
    <Container maxWidth="xs">
      <form className={classes.paper}>
        <Avatar className={classes.avatar}>
          <PaymentOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Deposit Money
        </Typography>
        <Grid className={classes.form} container spacing={2}>
          <Grid item xs={12}>
            <FormControl
              fullWidth
              required
              variant="outlined"
              error={dirty && (amount.length === 0 || parseFloat(amount) <= 0)}
            >
              <InputLabel htmlFor="deposit-amount">Amount</InputLabel>
              <OutlinedInput
                id="deposit-amount"
                name="deposit-amount"
                fullWidth
                startAdornment={
                  <InputAdornment position="start">$</InputAdornment>
                }
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                labelWidth={70}
                inputComponent={CurrencyFormatCustom}
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="card-number"
              label="Card Number"
              name="card-number"
              value={number}
              onChange={(e) => setNumber(e.target.value)}
              InputProps={{
                inputComponent: CreditCardFormatCustom,
              }}
              error={dirty && number.length !== 16}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="card-expiration-month"
              label="Expiration Month"
              name="card-expiration-month"
              autoComplete="bday-month"
              value={expirationMonth}
              onChange={(e) => setExpirationMonth(e.target.value)}
              InputProps={{
                inputComponent: MonthFormatCustom,
              }}
              error={dirty && expirationMonth.length !== 2}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="card-expiration-year"
              label="Expiration Year"
              name="card-expiration-year"
              autoComplete="bday-year"
              value={expirationYear}
              onChange={(e) => setExpirationYear(e.target.value)}
              InputProps={{
                inputComponent: YearFormatCustom,
              }}
              error={dirty && expirationYear.length !== 2}
            />
          </Grid>
          <Grid item xs={12} sm={9}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="card-holder-name"
              label="Holder Name"
              name="card-holder-name"
              autoComplete="name"
              value={holderName}
              onChange={(e) => setHolderName(e.target.value)}
              error={dirty && holderName.length === 0}
            />
          </Grid>
          <Grid item xs={12} sm={3}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="card-cvv"
              label="CVV"
              name="card-cvv"
              value={cvv}
              onChange={(e) => setCvv(e.target.value)}
              InputProps={{
                inputComponent: CvvFormatCustom,
              }}
              error={dirty && cvv.length !== 3}
            />
          </Grid>
        </Grid>
        <Button
          fullWidth
          variant="outlined"
          color="primary"
          className={classes.submit}
          onClick={handleDeposit}
        >
          Deposit
        </Button>
      </form>
    </Container>
  );
}

export default DepositMoneyDialog;
