import React from "react";
import { useParams } from "react-router-dom";

import { makeStyles } from "@material-ui/styles";
import Paper from "@material-ui/core/Paper";
import Typography from "@material-ui/core/Typography";

import Products from "../components/Products";

import { useGetStoreByIdQuery } from "../state/service";

const useStyles = makeStyles((theme) => ({
  root: {
    width: "100%",
    padding: theme.spacing(1),
    marginBottom: theme.spacing(2),
    overflowX: "auto",
  },
  subheader: {
    color: theme.palette.text.secondary,
    boxSizing: "border-box",
    fontFamily: theme.typography.fontFamily,
    fontWeight: 500,
    lineHeight: 1.5,
    textAlign: "center",
  },
}));

function Store() {
  const classes = useStyles();
  const params = useParams();

  const { data, error, isLoading } = useGetStoreByIdQuery(params.id);

  return (
    <React.Fragment>
      <Paper className={classes.root} elevation={25}>
        <Typography
          className={classes.subheader}
          component="h1"
          variant="header1"
        >
          {isLoading
            ? "Loading..."
            : error
            ? error.data && error.data.error
              ? error.data.error
              : error.status
            : data.data
            ? data.data.title
            : "Unkmown Owner"}
        </Typography>
      </Paper>
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>
          {error.data && error.data.error ? error.data.error : error.status}
        </div>
      ) : data &&
        data.data &&
        data.data.products &&
        data.data.products.length > 0 ? (
        <Products products={data.data.products} />
      ) : (
        <div>No data.</div>
      )}
    </React.Fragment>
  );
}

export default Store;
