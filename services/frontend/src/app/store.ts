import { configureStore } from '@reduxjs/toolkit'

import { apiSlice } from './api'
import authReducer from './features/auth'
import commentsReducer from './features/comments'
import panelsReducer from './features/panels'
import usersReducer from './features/users'

export const store = configureStore({
  reducer: {
    auth: authReducer,
    comments: commentsReducer,
    panels: panelsReducer,
    users: usersReducer,
    [apiSlice.reducerPath]: apiSlice.reducer,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware().concat(apiSlice.middleware)
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch