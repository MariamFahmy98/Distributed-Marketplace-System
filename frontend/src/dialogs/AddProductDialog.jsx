import React, { useEffect, useRef } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";

import { makeStyles, useTheme } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import EditIcon from "@material-ui/icons/Edit";
import FormControl from "@material-ui/core/FormControl";
import Grid from "@material-ui/core/Grid";
import IconButton from "@material-ui/core/IconButton";
import InputLabel from "@material-ui/core/InputLabel";
import InputAdornment from "@material-ui/core/InputAdornment";
import OutlinedInput from "@material-ui/core/OutlinedInput";
import TextField from "@material-ui/core/TextField";
import useMediaQuery from "@material-ui/core/useMediaQuery";

import optimizeImage from "../utils/image";
import { setDialog } from "../state/features/dialog";
import { showSnackbar } from "../state/features/snackbar";
import { useAddProductMutation } from "../state/service";

const useStyles = makeStyles(({ breakpoints, shape, spacing }) => ({
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
  cardMedia: {
    [breakpoints.only("xs")]: { height: 240 },
    [breakpoints.only("sm")]: { height: 360 },
    [breakpoints.up("md")]: { height: 480 },
    borderRadius: shape.borderRadius,
  },
  mediaEditLabel: {
    position: "absolute",
    top: 0,
    right: 0,
    zIndex: 1,
  },
  mediaEditIcon: {
    color: "white",
  },
  cardContent: {
    boxShadow: "0 16px 40px -12.125px rgba(0,0,0,0.3)",
    borderRadius: spacing(0.5),
    backgroundColor: "#ffffff",
    position: "absolute",
    top: "60%",
    left: "2%",
    width: "96%",
    padding: spacing(3),
    textAlign: "left",
  },
  input: {
    display: "none",
  },
}));

function AddProductDialog() {
  const classes = useStyles();
  const theme = useTheme();
  const navigate = useNavigate();

  const dispatch = useDispatch();

  const cardContent = useRef(null);
  const [name, setName] = React.useState("");
  const [price, setPrice] = React.useState("0");
  const [media, setMedia] = React.useState(null);
  const [description, setDescription] = React.useState("");
  const [cardContentHeight, setCardContentHeight] = React.useState(0);

  const matchesSM = useMediaQuery(theme.breakpoints.down("sm"));
  const matchesXS = useMediaQuery(theme.breakpoints.down("xs"));

  let heightOffset = Math.round(0.4 * 480);
  if (matchesSM) heightOffset = Math.round(0.4 * 360);
  if (matchesXS) heightOffset = Math.round(0.4 * 240);

  const [addProduct] = useAddProductMutation();

  useEffect(() => {
    setCardContentHeight(cardContent.current.offsetHeight);
  }, [cardContent]);

  const handleAdd = async () => {
    const optimizedImage = media ? await optimizeImage(media) : null;

    try {
      await addProduct({
        title: name,
        price: parseFloat(price),
        content: description,
        image_url: optimizedImage,
      }).unwrap();
      dispatch(
        showSnackbar({
          variant: "success",
          message: "Product added successfully",
        })
      );
      dispatch(setDialog(null));
      navigate("/profile");
    } catch (err) {
      dispatch(
        showSnackbar({
          variant: "error",
          message: err.data.error && err.data ? err.data.error : err.status,
        })
      );
    }
  };

  const handleNameChange = (event) => {
    setName(event.target.value);
  };

  const handlePriceChange = (event) => {
    setPrice(event.target.value);
  };

  const handleMediaChange = (event) => {
    const files = event.target.files;
    if (files.length > 0) setMedia(files[0]);
  };

  const handleDescriptionChange = (event) => {
    setDescription(event.target.value);
  };

  return (
    <form className={classes.relative} noValidate autoComplete="off">
      <Card
        style={{
          marginBottom: cardContentHeight - heightOffset + theme.spacing(2),
        }}
        className={classes.card}
      >
        <input
          accept="image/*"
          className={classes.input}
          id="icon-button-file"
          type="file"
          onChange={handleMediaChange}
        />
        <label className={classes.mediaEditLabel} htmlFor="icon-button-file">
          <IconButton aria-label="edit-media" component="span">
            <EditIcon className={classes.mediaEditIcon} />
          </IconButton>
        </label>
        <CardMedia
          component="img"
          className={classes.cardMedia}
          onError={() => {
            if (media) setMedia(null);
          }}
          src={
            media
              ? URL.createObjectURL(media)
              : "https://images.unsplash.com/photo-1460925895917-afdab827c52f?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1404&q=80"
          }
          alt={media ? media.name : "Product Photo"}
        />
        <CardContent ref={cardContent} className={classes.cardContent}>
          <Grid container spacing={2}>
            <Grid item xs={9}>
              <TextField
                id="product-name"
                label="Product Name"
                fullWidth
                autoFocus
                value={name}
                onChange={handleNameChange}
                variant="outlined"
              />
            </Grid>
            <Grid item xs={3}>
              <FormControl fullWidth variant="outlined">
                <InputLabel htmlFor="product-price">Price</InputLabel>
                <OutlinedInput
                  id="product-price"
                  fullWidth
                  startAdornment={
                    <InputAdornment position="start">$</InputAdornment>
                  }
                  value={price}
                  onChange={handlePriceChange}
                  labelWidth={40}
                />
              </FormControl>
            </Grid>
          </Grid>
          <br />
          <br />
          <TextField
            id="product-description"
            label="Product Description"
            multiline
            fullWidth
            rows={6}
            placeholder="Write your description here..."
            value={description}
            onChange={handleDescriptionChange}
            variant="filled"
          />
          <br />
          <br />
          <Button
            fullWidth
            variant="outlined"
            color="primary"
            onClick={handleAdd}
          >
            Add Now
          </Button>
        </CardContent>
      </Card>
    </form>
  );
}

export default AddProductDialog;
