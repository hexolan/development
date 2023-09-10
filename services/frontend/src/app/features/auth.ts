import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import type { User } from './users';

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
    setLoggedIn: (state, action: PayloadAction<string>) => {
      const token = action.payload
      try {
        let tokenPayload = JSON.parse(atob(token.split('.')[1]))
        state.token = token
        state.currentUser = {
          'id': tokenPayload['sub'],
          'username': 'test',
          'createdAt': new Date()
        }
      } catch {
        state.token = null
        state.currentUser = null
      }
    },
    setLoggedOut: (state, _action: PayloadAction) => {
      state.token = null
      state.currentUser = null
    }
  },
})

export const { setLoggedIn, setLoggedOut } = authSlice.actions
export default authSlice.reducer