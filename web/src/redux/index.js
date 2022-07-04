import { configureStore } from "@reduxjs/toolkit";
import { devToolsEnhancer } from 'redux-devtools-extension';

import homeReduser from "./home/homeReduser";

export const store = configureStore({
  reducer: {
    home: homeReduser,
  },
  devTools:[ devToolsEnhancer({ realtime: true }) ],
})