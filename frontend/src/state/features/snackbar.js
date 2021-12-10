import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  queue: [],
  open: false,
  messageInfo: undefined,
};

export const snackbarSlice = createSlice({
  name: "snackbar",
  initialState,
  reducers: {
    showSnackbar: (state, action) => {
      state.queue.push({
        key: new Date().getTime(),
        variant: action.payload.variant,
        message: action.payload.message,
        actionLabel: action.payload.actionLabel,
        action: action.payload.action,
      });

      if (!state.open) {
        if (state.queue.length > 0) {
          state.open = true;
          state.messageInfo = state.queue.shift();
        }
      }
    },
    setSnackbar: (state, action) => {
      state.open = action.payload.open;
      if (action.payload.messageInfo)
        state.messageInfo = action.payload.messageInfo;
    },
  },
});

export const { setSnackbar, showSnackbar } = snackbarSlice.actions;

export default snackbarSlice.reducer;
