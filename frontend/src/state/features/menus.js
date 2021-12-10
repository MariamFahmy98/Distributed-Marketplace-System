import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  account: null,
};

export const menusSlice = createSlice({
  name: "menus",
  initialState,
  reducers: {
    setMenuAnchor: (state, action) => {
      state[action.payload.menu] = action.payload.anchor;
    },
  },
});

export const { setMenuAnchor } = menusSlice.actions;

export default menusSlice.reducer;
