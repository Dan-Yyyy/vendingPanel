import { configureStore } from "@reduxjs/toolkit";
import { devToolsEnhancer } from 'redux-devtools-extension';

import homeReduser from "./home/homeReduser";
import authReduser from "./auth/authReduser";

export const store = configureStore({
  reducer: {
    home: homeReduser,
    auth: authReduser,
  },
  devTools:[ devToolsEnhancer({ realtime: true }) ],
})