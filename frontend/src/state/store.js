import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { marketplaceApi } from "./service";

import userReducer from "./features/user";
import dialogReducer from "./features/dialog";
import menusReducer from "./features/menus";
import formsReducer from "./features/forms";
import snackbarReducer from "./features/snackbar";

export const store = configureStore({
  reducer: {
    user: userReducer,
    dialog: dialogReducer,
    menus: menusReducer,
    forms: formsReducer,
    snackbar: snackbarReducer,
    [marketplaceApi.reducerPath]: marketplaceApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(marketplaceApi.middleware),
});

setupListeners(store.dispatch);
