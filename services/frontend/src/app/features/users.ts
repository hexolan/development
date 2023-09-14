import { createSlice, createEntityAdapter } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { usersApiSlice } from '../api/users'
import type { User } from '../types'

const usersAdapter = createEntityAdapter<User>({
  selectId: (user) => user.id
})

export const usersSlice = createSlice({
  name: 'users',
  initialState: usersAdapter.getInitialState(),
  reducers: {
    setUser: (state, action: PayloadAction<User>) => {
      usersAdapter.upsertOne(state, action.payload)
    }
  },
  extraReducers: (builder) => {
    
  }
})

export const { setUser } = usersSlice.actions
export default usersSlice.reducer