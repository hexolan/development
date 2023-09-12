import { createSlice, createEntityAdapter } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import type { Post } from '../types'

const postsAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id
})

export const postsSlice = createSlice({
  name: 'posts',
  initialState: postsAdapter.getInitialState(),
  reducers: {
    postReceived(state, action: PayloadAction<Post>) {
      postsAdapter.upsertOne(state, action.payload)
    },
    postsReceived(state, action: PayloadAction<Post[]>) {
      postsAdapter.upsertMany(state, action.payload)
    }
  },
})

export const { } = postsSlice.actions
export default postsSlice.reducer