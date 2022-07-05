import { createAsyncThunk } from "@reduxjs/toolkit";
import { auth } from "../../Api/authApi";

export const authThunk = createAsyncThunk(
  'auth/singIn', 
  async (userData) => {
    try{
      const result = await auth.signIn(userData)
      .then(response => {
        return response;
      });
      return result;
    }
    catch(err) {
      return { 
        message: err.response.data.message,
        status: err.response.status
      };
    }
  }
);