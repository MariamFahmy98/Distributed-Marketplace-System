import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  id: null,
  name: null,
  image_url: null,
  balance: null,
  token: null,
  store_id: null,
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (state, action) => {
      state.id = action.payload ? action.payload.id : null;
      state.name = action.payload ? action.payload.name : null;
      state.image_url = action.payload ? action.payload.image_url : null;
      state.balance = action.payload ? action.payload.balance : null;
      state.token = action.payload ? action.payload.token : null;
      state.store_id = action.payload ? action.payload.store_id : null;
    },
    changePicture: (state, action) => {
      state.image_url = action.payload ? action.payload : null;
    },
    incrementBalance: (state, action) => {
      state.balance += action.payload;
    },
    decrementBalance: (state, action) => {
      state.balance -= action.payload;
    },
  },
});

export const { setUser, changePicture, incrementBalance, decrementBalance } =
  userSlice.actions;

export default userSlice.reducer;
