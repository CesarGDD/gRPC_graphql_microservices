import { configureStore } from '@reduxjs/toolkit';
import shoppingReducer from './shoppingcart/shoppingSlice';
import userReducer from './users/userSlice';

export const store = configureStore({
  reducer: {
    basket: shoppingReducer,
    user: userReducer
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
