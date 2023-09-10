import { configureStore } from '@reduxjs/toolkit'

import { apiSlice } from './api'
import authReducer from './slices/auth'
import panelsReducer from './slices/panels'
import postsReducer from './slices/posts'
import usersReducer from './slices/users'

export const store = configureStore({
  reducer: {
    [apiSlice.reducerPath]: apiSlice.reducer,
    auth: authReducer,
    panels: panelsReducer,
    posts: postsReducer,
    users: usersReducer,
  }, 
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch