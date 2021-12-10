import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  order: null,
};

export const formsSlice = createSlice({
  name: "forms",
  initialState,
  reducers: {
    setForm: (state, action) => {
      state[action.payload.form] = action.payload.values;
    },
  },
});

export const { setForm } = formsSlice.actions;

export default formsSlice.reducer;
