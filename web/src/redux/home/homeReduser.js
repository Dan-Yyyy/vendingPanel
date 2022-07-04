import { createSlice } from "@reduxjs/toolkit";

const homeSlise = createSlice({
  name: "home",
  initialState: {
    status: 'fullfield',
    items: [
      {
        id: 1,
        adress: 'ул. Осктябрьская 24',
        mashine: 'Necta',
        isVisited: true,
      },
      {
        id: 2,
        adress: 'ул. Осктябрьская 222',
        mashine: 'Necta',
        isVisited: false,
      },
      {
        id: 3,
        adress: 'ул. Осктябрьская 22',
        mashine: 'Necta',
        isVisited: true,
      },
    ]
  },
  reducers: {
    visited(state, action) {
      state.items = state.items.map(item => {
        if(item.id === action.payload) {
          item.isVisited = !item.isVisited
        }
        return item;
      })
    },
  },
})

export default homeSlise.reducer;
export const { visited } = homeSlise.actions;