import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

export interface User {
  id: string;
  username: string;
  createdAt: Date;
  updatedAt?: Date;
}

export type Users = {
  [idKey: string]: User;
}

export interface UsersState {
  items: Users;
  loading: boolean;
}

const initialState: UsersState = {
  items: {},
  loading: false
}

export const usersSlice = createSlice({
  name: 'users',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<User>) => {
      const user = action.payload;
      state.items[user.id] = user
    }
  },
})

export const { setUser } = usersSlice.actions
export default usersSlice.reducer