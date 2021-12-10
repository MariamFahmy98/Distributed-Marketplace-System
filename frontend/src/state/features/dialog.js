import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  value: null,
};

export const dialogSlice = createSlice({
  name: "dialog",
  initialState,
  reducers: {
    setDialog: (state, action) => {
      state.value = action.payload;
    },
  },
});

export const { setDialog } = dialogSlice.actions;

export default dialogSlice.reducer;
