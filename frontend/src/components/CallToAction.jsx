import React from "react";
import PropTypes from "prop-types";

import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardHeader from "@material-ui/core/CardHeader";
import Typography from "@material-ui/core/Typography";

import Button from "./Button";

const useStyles = makeStyles((theme) => {
  const cardShadow = "0px 14px 80px rgba(34, 35, 58, 0.2)";
  const headerShadow = "4px 4px 20px 1px rgba(33, 203, 243, .3)";
  return {
    card: {
      borderRadius: theme.shape.borderRadius,
      transition: "0.3s",
      boxShadow: cardShadow,
      position: "relative",
      overflow: "initial",
      background: "#ffffff",
      marginTop: theme.spacing(4),
      padding: theme.spacing(2, 0, 0, 0),
      [theme.breakpoints.only("xs")]: {
        marginTop: theme.spacing(6),
      },
    },
    header: {
      flexShrink: 0,
      position: "absolute",
      right: 20,
      left: 20,
      borderRadius: theme.spacing(2),
      background: `linear-gradient(45deg, ${theme.palette.primary.dark} 30%, ${theme.palette.primary.light} 90%)`,
      overflow: "hidden",
      boxShadow: headerShadow,
      textAlign: "center",
      top: theme.spacing(-4),
      [theme.breakpoints.only("xs")]: {
        top: theme.spacing(-6),
      },
    },
    title: {
      color: theme.palette.getContrastText(theme.palette.primary.main),
      fontWeight: 700,
      letterSpacing: 1,
    },
    subheader: {
      color: theme.palette.getContrastText(theme.palette.primary.main),
      opacity: 0.87,
      fontWeight: 400,
      letterSpacing: 0.4,
    },
    contentContainer: {
      textAlign: "center",
    },
    content: {
      paddingTop: "38px",
      overflowX: "auto",
    },
  };
});

export default function CallToAction(props) {
  const classes = useStyles();

  return (
    <Card className={classes.card}>
      <CardHeader
        className={classes.header}
        title={
          <Typography
            variant="subtitle1"
            component="h1"
            className={classes.title}
          >
            {props.headerTitle}
          </Typography>
        }
        subheader={
          <Typography
            variant="subtitle2"
            component="h1"
            className={classes.subheader}
          >
            {props.subheaderTitle}
          </Typography>
        }
      />
      <CardContent className={classes.contentContainer}>
        <div className={classes.content}>
          <Typography variant="subtitle1" component="h3">
            {props.subtitle}
          </Typography>
          <Button color="red" onClick={props.handlePrimaryAction}>
            {props.primaryActionText}
          </Button>
          {props.showSecondaryAction && (
            <Button color="red" onClick={props.handleSecondaryAction}>
              {props.secondaryActionText}
            </Button>
          )}
        </div>
      </CardContent>
    </Card>
  );
}

CallToAction.propTypes = {
  headerTitle: PropTypes.string.isRequired,
  subheaderTitle: PropTypes.string.isRequired,
  subtitle: PropTypes.string.isRequired,
  primaryActionText: PropTypes.string.isRequired,
  handlePrimaryAction: PropTypes.func.isRequired,
  showSecondaryAction: PropTypes.bool,
  secondaryActionText: PropTypes.string,
  handleSecondaryAction: PropTypes.func,
};
