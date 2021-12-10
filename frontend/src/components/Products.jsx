import React, { useState } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

import { emphasize, makeStyles, withStyles } from "@material-ui/core/styles";
import { red } from "@material-ui/core/colors";
import Avatar from "@material-ui/core/Avatar";
import Breadcrumbs from "@material-ui/core/Breadcrumbs";
import Button from "@material-ui/core/Button";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import CardHeader from "@material-ui/core/CardHeader";
import CardMedia from "@material-ui/core/CardMedia";
import Chip from "@material-ui/core/Chip";
import EditIcon from "@material-ui/icons/Edit";
import Grid from "@material-ui/core/Grid";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import StorefrontIcon from "@material-ui/icons/Storefront";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import PlaylistAddIcon from "@material-ui/icons/PlaylistAdd";
import ShoppingCartIcon from "@material-ui/icons/ShoppingCart";
import MoreVertIcon from "@material-ui/icons/MoreVert";

import optimizeImage from "../utils/image";
import { setDialog } from "../state/features/dialog";
import { setForm } from "../state/features/forms";
import { showSnackbar } from "../state/features/snackbar";
import {
  useAddProductToStoreMutation,
  useDeleteProductMutation,
  useEditProductMutation,
} from "../state/service";

const userPropTypes = {
  id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  image_url: PropTypes.string,
};

const StyledBreadcrumb = withStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.grey[100],
    height: theme.spacing(3),
    color: theme.palette.grey[800],
    fontWeight: theme.typography.fontWeightRegular,
    "&:hover, &:focus": {
      backgroundColor: theme.palette.grey[300],
    },
    "&:active": {
      boxShadow: theme.shadows[1],
      backgroundColor: emphasize(theme.palette.grey[300], 0.12),
    },
  },
}))(Chip);

const useProductStyles = makeStyles(({ typography }) => ({
  hide: { display: "none" },
  cardMedia: {
    height: 240,
    objectFit: "cover",
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
  avatar: {
    backgroundColor: red[500],
  },
  nameInput: {
    fontFamily: typography.subtitle1.fontFamily,
    fontWeight: typography.subtitle1.fontWeight,
    fontSize: typography.subtitle1.fontSize,
    lineHeight: typography.subtitle1.lineHeight,
    letterSpacing: typography.subtitle1.letterSpacing,
    padding: 0,
    blockSize: "fit-content",
  },
  priceInput: {
    fontFamily: typography.caption.fontFamily,
    fontWeight: typography.caption.fontWeight,
    fontSize: typography.caption.fontSize,
    lineHeight: typography.caption.lineHeight,
    letterSpacing: typography.caption.letterSpacing,
    padding: 0,
    blockSize: "fit-content",
  },
  atEnd: {
    marginLeft: "auto",
  },
}));

export function Product(props) {
  const classes = useProductStyles();
  const navigate = useNavigate();

  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);

  const { id, title, user, price, image_url, content, handleProductDelete } =
    props;

  const [anchorEl, setAnchorEl] = React.useState(null);
  const [isEditing, setIsEditing] = React.useState(false);
  const [nameState, setNameState] = React.useState(title);
  const [priceState, setPriceState] = React.useState(price);
  const [media, setMedia] = React.useState(null);
  const [descriptionState, setDescriptionState] = React.useState(content);

  const [addProductToStore] = useAddProductToStoreMutation();
  const [editProduct] = useEditProductMutation();
  const [deleteProduct] = useDeleteProductMutation();

  const handleMenuClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleAddProductToStore = async () => {
    try {
      const data = await addProductToStore(id).unwrap();
      if (!data.sccess) {
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
            message: "Product added to your store",
          })
        );
      }
    } catch (err) {
      dispatch(
        showSnackbar({
          variant: "error",
          message: err.data.error && err.data ? err.data.error : err.status,
        })
      );
    }
  };

  const handleBuyProduct = () => {
    dispatch(
      setForm({
        form: "order",
        values: {
          id,
          name: nameState,
          user,
          price: priceState,
          image_url: media ? media : image_url ? image_url : null,
          description: descriptionState,
        },
      })
    );
    dispatch(setDialog("order-product"));
  };

  const handleDeleteProduct = async () => {
    try {
      const data = await deleteProduct(id).unwrap();
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
            message: "Product deleted",
          })
        );
      }
    } catch (err) {
      dispatch(
        showSnackbar({
          variant: "error",
          message: err.data.error && err.data ? err.data.error : err.status,
        })
      );
    }
    handleProductDelete(id);
  };

  const handleEditProduct = async () => {
    if (isEditing) {
      const optimizedImage = media ? await optimizeImage(media) : null;

      try {
        await editProduct({
          id,
          product: {
            title: nameState,
            price: parseFloat(priceState),
            content: descriptionState,
            image_url: optimizedImage,
          },
        }).unwrap();
        dispatch(
          showSnackbar({
            variant: "success",
            message: "Product edited successfully",
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

    setIsEditing((old_value) => {
      return !old_value;
    });
  };

  const handleOnNameChanged = (event) => {
    setNameState(event.target.value);
  };
  const handleOnPriceChanged = (event) => {
    setPriceState(event.target.value);
  };
  const handleOnMediaChanged = async (event) => {
    const files = event.target.files;
    if (files.length > 0) {
      setMedia(files[0]);
    }
  };
  const handleOnDescriptionChanged = (event) => {
    setDescriptionState(event.target.value);
  };

  function handleStoreClick(event) {
    event.preventDefault();
    navigate(`/store/${encodeURIComponent(user.id)}`);
  }

  return (
    <Card elevation={25} style={{ position: "relative" }}>
      <CardHeader
        avatar={
          user.image_url ? (
            <Avatar src={user.image_url} aria-label={user.name} />
          ) : (
            <Avatar aria-label="avatar" className={classes.avatar}>
              {user.name.length > 0 ? user.name[0] : "A"}
            </Avatar>
          )
        }
        action={
          current_user.id === user.id && (
            <React.Fragment>
              <IconButton aria-label="more" onClick={handleMenuClick}>
                <MoreVertIcon />
              </IconButton>
              <Menu
                id="more-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleMenuClose}
              >
                <MenuItem onClick={handleDeleteProduct}>Delete</MenuItem>
              </Menu>
            </React.Fragment>
          )
        }
        title={
          isEditing ? (
            <InputBase
              id="name-input"
              name="name-input"
              classes={{ input: classes.nameInput }}
              autoFocus
              fullWidth
              value={nameState}
              onChange={handleOnNameChanged}
              inputProps={{ "aria-label": "naked" }}
            />
          ) : (
            <Breadcrumbs aria-label="breadcrumb">
              <Typography variant="subtitle1">{nameState}</Typography>
              <StyledBreadcrumb
                component="a"
                href="#"
                label="Store"
                icon={<StorefrontIcon fontSize="small" />}
                onClick={handleStoreClick}
              />
            </Breadcrumbs>
          )
        }
        subheader={
          isEditing ? (
            <InputBase
              id="price-input"
              name="price-input"
              classes={{ input: classes.priceInput }}
              fullWidth
              value={priceState}
              onChange={handleOnPriceChanged}
              inputProps={{ "aria-label": "naked" }}
            />
          ) : (
            <Typography variant="caption">{`$${priceState}`}</Typography>
          )
        }
      />
      <div className={classes.cardMedia} style={{ position: "relative" }}>
        {isEditing && (
          <React.Fragment>
            <input
              id="media-input"
              name="media-input"
              className={classes.hide}
              accept="image/*"
              type="file"
              onChange={handleOnMediaChanged}
            />
            <label className={classes.mediaEditLabel} htmlFor="media-input">
              <IconButton aria-label="edit-media" component="span">
                <EditIcon className={classes.mediaEditIcon} />
              </IconButton>
            </label>
          </React.Fragment>
        )}
        <CardMedia
          component="img"
          className={classes.cardMedia}
          onError={() => {
            if (media) setMedia(null);
          }}
          src={
            media
              ? URL.createObjectURL(media)
              : image_url
              ? image_url
              : "https://images.unsplash.com/photo-1460925895917-afdab827c52f?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1404&q=80"
          }
          alt={media ? media.name : nameState}
        />
      </div>
      <CardContent>
        {isEditing ? (
          <TextField
            id="description-input"
            name="description-input"
            label="Description"
            multiline
            fullWidth
            rows={3}
            value={descriptionState}
            onChange={handleOnDescriptionChanged}
            variant="filled"
          />
        ) : (
          <Typography variant="body2" color="textSecondary" component="p">
            {descriptionState}
          </Typography>
        )}
      </CardContent>
      <CardActions disableSpacing>
        <IconButton aria-label="add" onClick={handleAddProductToStore}>
          <PlaylistAddIcon />
        </IconButton>
        <IconButton aria-label="buy" onClick={handleBuyProduct}>
          <ShoppingCartIcon />
        </IconButton>
        {current_user.id === user.id &&
          (isEditing ? (
            <Button
              className={classes.atEnd}
              variant="outlined"
              color="primary"
              onClick={handleEditProduct}
            >
              Save
            </Button>
          ) : (
            <IconButton
              className={classes.atEnd}
              aria-label="edit"
              onClick={handleEditProduct}
            >
              <EditIcon />
            </IconButton>
          ))}
      </CardActions>
    </Card>
  );
}

const productPropTypes = {
  id: PropTypes.number.isRequired,
  title: PropTypes.string.isRequired,
  user: PropTypes.exact(userPropTypes),
  price: PropTypes.number.isRequired,
  image_url: PropTypes.string,
  content: PropTypes.string.isRequired,
  handleProductDelete: PropTypes.func,
};

Product.propTypes = productPropTypes;

function Products(props) {
  const [products, setProducts] = useState(props.products);

  const handleProductDeletion = (id) => {
    setProducts((products) => products.filter((el) => el.id !== id));
  };

  return (
    <Grid container spacing={4}>
      {products.map((product) => (
        <Grid key={product.id} item xs={12} sm={6}>
          <Product {...product} handleProductDelete={handleProductDeletion} />
        </Grid>
      ))}
    </Grid>
  );
}

const productsPropTypes = {
  products: PropTypes.arrayOf(PropTypes.exact(productPropTypes)).isRequired,
};

Products.propTypes = productsPropTypes;

export default Products;
