import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

export interface AuthState {
  token: string | null;
  userId: string | null;
}

const initialState: AuthState = {
  token: null,
  userId: null
}

export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setToken: (state, action: PayloadAction<string>) => {
      const token = action.payload
      try {
        // decode the JWT token to capture the user ID
        let tokenPayload = JSON.parse(atob(token.split('.')[1]))
        state.token = token
        state.userId = tokenPayload['sub']
      } catch {
        state.token = null
        state.userId = null
      }
    },
    logOut: (state, _action: PayloadAction) => {
      state.token = null
      state.userId = null
    }
  },
})

export const { setToken } = authSlice.actions
export default authSlice.reducer