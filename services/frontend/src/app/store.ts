import { configureStore } from '@reduxjs/toolkit'

import authReducer from './features/auth';
import panelsReducer from './features/panels';
import postsReducer from './features/posts';
import usersReducer from './features/users';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    panels: panelsReducer,
    posts: postsReducer,
    users: usersReducer,
  }
});

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch