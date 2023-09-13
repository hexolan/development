import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import type { AuthToken, User } from '../types'

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
    setSignedIn: (state, action: PayloadAction<{ token: AuthToken, user: User }>) => {
      console.log(action.payload)
      const { token, user } = action.payload
      state.token = token.access_token
      state.currentUser = user
    },
    setSignedOut: (state, _action: PayloadAction) => {
      state.token = null
      state.currentUser = null
    }
  },
})

export const { setSignedIn, setSignedOut } = authSlice.actions
export default authSlice.reducer