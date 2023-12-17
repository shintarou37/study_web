import { createSlice } from "@reduxjs/toolkit";

export const counterSlice = createSlice({
  // sliceの名前
  name: "counter",
  // 初期値
  initialState: {
    value: 0,
  },
  // stateの状態を変更するための関数
  // reducersを作るとaction createも一緒に作られる
  reducers: {
    increment: (state) => {
      state.value++;
    },
    decrement: (state)=> {
      state.value--;
    },
    incrementByInput: (state, action)=> {
      state.value += action.payload
    }
  }
})

// viewで使用するためexportする必要がある
export const {increment, decrement, incrementByInput} = counterSlice.actions;
export default counterSlice.reducer;