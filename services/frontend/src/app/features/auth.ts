import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { useAppDispatch } from '../hooks';
import { SignInData, useSignInMutation } from './authApi'
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
    setSignedIn: (state, action: PayloadAction<{ token: string, currentUser: User }>) => {
      const { token, currentUser } = action.payload
      state.token = token
      state.currentUser = currentUser
    },
    setSignedOut: (state, _action: PayloadAction) => {
      state.token = null
      state.currentUser = null
    }
  },
})

export const { setSignedIn, setSignedOut } = authSlice.actions
export default authSlice.reducer

/*
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
*/