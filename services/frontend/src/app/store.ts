import { configureStore } from '@reduxjs/toolkit'

import { apiSlice } from './api'
import authReducer from './features/auth'
import panelsReducer from './features/panels'
import postsReducer from './features/posts'
import usersReducer from './features/users'

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