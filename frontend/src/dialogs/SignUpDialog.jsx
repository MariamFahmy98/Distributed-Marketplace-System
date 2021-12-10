import React, { useState } from "react";
import { useDispatch } from "react-redux";

import { makeStyles } from "@material-ui/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Link from "@material-ui/core/Link";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";

import { setUser } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { showSnackbar } from "../state/features/snackbar";
import { useSignUpMutation } from "../state/service";

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

function SignUpDialog() {
  const classes = useStyles();

  const dispatch = useDispatch();

  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [signUp] = useSignUpMutation();

  async function handleEmailSignUp() {
    try {
      const data = await signUp({
        name: `${firstName} ${lastName}`,
        email,
        password,
      }).unwrap();
      dispatch(setUser({ token: data.token, ...data.user }));
      localStorage.setItem(
        "user",
        JSON.stringify({ token: data.token, ...data.user })
      );
      dispatch(setDialog(null));
      dispatch(
        showSnackbar({ variant: "success", message: "Successfully signed up" })
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

  return (
    <Container maxWidth="xs">
      <form className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign-Up
        </Typography>
        <Grid className={classes.form} container spacing={2}>
          <Grid item xs={12} sm={6}>
            <TextField
              autoComplete="fname"
              name="firstName"
              variant="outlined"
              required
              fullWidth
              id="firstName"
              label="First Name"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="lastName"
              label="Last Name"
              name="lastName"
              autoComplete="lname"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              variant="outlined"
              required
              fullWidth
              id="email"
              label="Email"
              name="email"
              autoComplete="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              variant="outlined"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="new-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </Grid>
        </Grid>
        <Button
          fullWidth
          variant="outlined"
          color="primary"
          className={classes.submit}
          onClick={handleEmailSignUp}
        >
          Sign Up
        </Button>
        <Grid container justify="flex-end">
          <Grid item>
            <Link
              component="button"
              variant="body2"
              onClick={() => {
                dispatch(setDialog("sign-in"));
              }}
            >
              Already have an account? Sign In
            </Link>
          </Grid>
        </Grid>
      </form>
    </Container>
  );
}

export default SignUpDialog;
