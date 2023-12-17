import { configureStore } from "@reduxjs/toolkit";
import counterSlice from './counterSlice';
// configureStoreはstoreを簡単に作成できる関数
export const store = configureStore({
  // storeにreducer（stateの状態を変更するための関数）を追加する
  reducer: {
    counter: counterSlice,
  },
})