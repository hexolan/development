import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { authApiSlice } from '../api/auth'
import type { LoginResponseData } from '../api/auth'
import type { User } from '../types'

export interface AuthState {
  token: string | null;
  currentUser: User | null;
}

const initialState: AuthState = {
  token: null,
  currentUser: null
}

export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setAuthed: (state, action: PayloadAction<LoginResponseData>) => {
      state.token = action.payload.token.access_token
      state.currentUser = action.payload.user
    },
    setUnauthed: (state, _action: PayloadAction<void>) => {
      state.token = null
      state.currentUser = null
    }
  },
  extraReducers: (builder) => {
    builder.addMatcher(
      authApiSlice.endpoints.login.matchFulfilled,
      (state, { payload }) => {
        state.token = payload.token.access_token
        state.currentUser = payload.user
      }
    )
  },
})

export const { setAuthed, setUnauthed } = authSlice.actions
export default authSlice.reducer