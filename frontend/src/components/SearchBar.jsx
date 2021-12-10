import React from "react";
import PropTypes from "prop-types";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import match from "autosuggest-highlight/match";
import parse from "autosuggest-highlight/parse";

import { makeStyles } from "@material-ui/core/styles";
import Autocomplete from "@material-ui/lab/Autocomplete";
import CircularProgress from "@material-ui/core/CircularProgress";
import InputBase from "@material-ui/core/InputBase";
import Link from "@material-ui/core/Link";
import SearchIcon from "@material-ui/icons/SearchRounded";

import { useLazySearchQuery } from "../state/service";

function renderInputComponent(props) {
  const { params, classes, loading, ...other } = props;

  return (
    <div>
      <InputBase
        fullWidth
        className={classes.searchBar}
        startAdornment={
          <React.Fragment>
            {loading ? (
              <CircularProgress
                className={classes.inputAdornment}
                color="inherit"
                size={20}
              />
            ) : (
              <SearchIcon className={classes.inputAdornment} color="action" />
            )}
          </React.Fragment>
        }
        inputProps={{ ...params.inputProps, autoComplete: "off" }}
        classes={{
          input: classes.input,
        }}
        ref={params.InputProps.ref}
        {...other}
      />
    </div>
  );
}

renderInputComponent.propTypes = {
  params: PropTypes.object,
  classes: PropTypes.object,
  loading: PropTypes.bool,
};

const useStyles = makeStyles((theme) => ({
  autoComplete: {
    flexGrow: "1",
  },
  searchBar: {
    borderRadius: "100px",
    backgroundColor: theme.palette.grey["200"],
    "&:hover": {
      backgroundColor: theme.palette.grey["100"],
    },
    "&:focus-within": {
      backgroundColor: theme.palette.background.default,
      boxShadow: `0 0 0px 2px ${theme.palette.primary.main}`,
    },
  },
  inputAdornment: {
    marginLeft: theme.spacing(1),
  },
  input: {
    padding: theme.spacing(1.25, 0, 0.75, 1),
  },
}));

function SearchBar(props) {
  const classes = useStyles();
  const navigate = useNavigate();

  const [open, setOpen] = React.useState(false);
  const [query, setQuery] = React.useState("");
  const [options, setOptions] = React.useState([]);

  const [search, { data, error, isLoading: loading }] = useLazySearchQuery();

  React.useEffect(() => {
    if (query) search(query);
  }, [search, query]);

  React.useEffect(() => {
    if (data && !error) setOptions(data.data);
  }, [data, error]);

  React.useEffect(() => {
    if (!open) {
      setOptions([]);
    }
  }, [open]);

  const handleInputChange = (_event, value) => {
    setQuery(value);
  };

  return (
    <Autocomplete
      id={props.id}
      className={classes.autoComplete}
      open={open}
      loadingText="Loading..."
      noOptionsText="No products found for this search"
      autoComplete
      autoHighlight
      onOpen={() => {
        setOpen(true);
      }}
      onClose={() => {
        setOpen(false);
      }}
      onInputChange={handleInputChange}
      onChange={(event_, value, reason) => {
        if (reason === "select-option")
          navigate(`/product/${encodeURIComponent(value.id)}/`);
      }}
      getOptionLabel={(option) => {
        return option.title;
      }}
      options={options}
      loading={loading}
      inputValue={query}
      renderInput={(params) =>
        renderInputComponent({
          params,
          classes,
          loading, // TODO: must be passed through with params
          placeholder: props.placeholder,
          onFocus: props.onFocus,
          onBlur: props.onBlur,
        })
      }
      renderOption={(option, { inputValue }) => {
        const productSuggestion = option.title;
        const productMatches = match(productSuggestion, inputValue);
        const productParts = parse(productSuggestion, productMatches);

        return (
          <Link
            component={RouterLink}
            to={`/product/${encodeURIComponent(option.id)}/`}
          >
            <div>
              {productParts.map((part, index) => (
                <span
                  key={index}
                  style={{ fontWeight: part.highlight ? 700 : 400 }}
                >
                  {part.text}
                </span>
              ))}
            </div>
          </Link>
        );
      }}
      getOptionSelected={(option, value) => option.id === value.id}
    />
  );
}

SearchBar.propTypes = {
  id: PropTypes.string.isRequired,
  placeholder: PropTypes.string.isRequired,
  onFocus: PropTypes.func,
  onBlur: PropTypes.func,
};

export default SearchBar;
