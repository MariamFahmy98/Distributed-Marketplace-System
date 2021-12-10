import React, { useState } from "react";
import { useDispatch } from "react-redux";

import { makeStyles } from "@material-ui/styles";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import MuiLink from "@material-ui/core/Link";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";

import { setUser } from "../state/features/user";
import { setDialog } from "../state/features/dialog";
import { showSnackbar } from "../state/features/snackbar";
import { useSignInMutation } from "../state/service";

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
  submit: {
    margin: theme.spacing(2, 0),
  },
}));

function SignInDialog() {
  const classes = useStyles();

  const dispatch = useDispatch();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [signIn] = useSignInMutation();

  async function handleEmailSignIn() {
    try {
      const data = await signIn({ email, password }).unwrap();
      dispatch(setUser({ token: data.token, ...data.user }));
      localStorage.setItem(
        "user",
        JSON.stringify({ token: data.token, ...data.user })
      );
      dispatch(setDialog(null));
      dispatch(
        showSnackbar({ variant: "success", message: "Successfully signed in" })
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
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign-In
        </Typography>
        <Grid className={classes.form} container spacing={2}>
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
              autoComplete="current-password"
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
          onClick={handleEmailSignIn}
        >
          Sign In
        </Button>
        <Grid container>
          <Grid item xs>
            <MuiLink component="button" variant="body2">
              Forgot Password?
            </MuiLink>
          </Grid>
          <Grid item>
            <MuiLink
              component="button"
              variant="body2"
              onClick={() => {
                dispatch(setDialog("sign-up"));
              }}
            >
              {"Don't have an account? Sign up"}
            </MuiLink>
          </Grid>
        </Grid>
      </div>
    </Container>
  );
}

export default SignInDialog;
