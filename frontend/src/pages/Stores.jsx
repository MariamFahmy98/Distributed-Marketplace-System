import React from "react";

import { makeStyles } from "@material-ui/core/styles";
import { red } from "@material-ui/core/colors";
import Avatar from "@material-ui/core/Avatar";
import Divider from "@material-ui/core/Divider";
import Link from "@material-ui/core/Link";
import { Link as RouterLink } from "react-router-dom";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemAvatar from "@material-ui/core/ListItemAvatar";
import ListItemText from "@material-ui/core/ListItemText";
import Paper from "@material-ui/core/Paper";
import Skeleton from "@material-ui/lab/Skeleton";

import { useGetStoresQuery } from "../state/service";

const useStyles = makeStyles((theme) => ({
  avatar: {
    backgroundColor: red[500],
  },
  fab: {
    position: "fixed",
    bottom: 0,
    right: 0,
    margin: theme.spacing(4),
  },
}));

function Stores() {
  const classes = useStyles();

  const { data, error, isLoading: loading } = useGetStoresQuery();

  if (loading)
    return (
      <Paper>
        <List>
          {[0, 1, 2, 3].map((id) => (
            <React.Fragment key={id}>
              <ListItem alignItems="flex-start">
                <ListItemAvatar>
                  <Skeleton variant="circle">
                    <Avatar />
                  </Skeleton>
                </ListItemAvatar>
                <ListItemText
                  primary={
                    <Skeleton
                      animation="wave"
                      height={10}
                      width="40%"
                      style={{ marginBottom: 6 }}
                    />
                  }
                  secondary={
                    <Skeleton animation="wave" height={10} width="20%" />
                  }
                />
              </ListItem>
              <Divider variant="inset" component="li" />
            </React.Fragment>
          ))}
        </List>
      </Paper>
    );
  if (error) return null;

  return (
    <Paper>
      <List>
        {data.data.map(({ id, title }) => (
          <React.Fragment key={id}>
            <ListItem alignItems="flex-start">
              <ListItemAvatar>
                <Avatar className={classes.avatar} alt={title} />
              </ListItemAvatar>
              <ListItemText
                primary={
                  <Link
                    component={RouterLink}
                    to={`/store/${encodeURIComponent(id)}`}
                  >
                    {title}
                  </Link>
                }
              />
            </ListItem>
            <Divider variant="inset" component="li" />
          </React.Fragment>
        ))}
      </List>
    </Paper>
  );
}

export default Stores;
