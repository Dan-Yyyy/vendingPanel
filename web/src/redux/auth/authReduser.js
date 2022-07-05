import { createSlice } from "@reduxjs/toolkit";
import { authThunk } from "./authThunk";

const authReduser = createSlice({
  name: "auth",
  initialState: {
    isAuth: false,
    message: '',
    token: '', 
    status: ''
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
    .addCase(authThunk.pending, (state) => {
      state.status = "pending";
    })
    .addCase(authThunk.fulfilled, (state, {payload}) => {
      if(payload.status === 200) {
        state.isAuth = true;
        state.token = payload.data.token;
      }
      state.message = payload.message;
      state.status = "fulfilled";
    })
  }
})

export default authReduser.reducer;