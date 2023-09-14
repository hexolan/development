import { createSlice, createEntityAdapter } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { postsApiSlice } from '../api/posts'
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
  extraReducers: (builder) => {
    builder.addMatcher(
      postsApiSlice.endpoints.getPanelPost.matchFulfilled,
      (state, { payload }) => {
        postsAdapter.upsertOne(state, payload)
      }
    ).addMatcher(
      postsApiSlice.endpoints.getPanelPosts.matchFulfilled,
      (state, { payload }) => {
        postsAdapter.upsertMany(state, payload)
      }
    ).addMatcher(
      postsApiSlice.endpoints.createPanelPost.matchFulfilled,
      (state, { payload }) => {
        postsAdapter.addOne(state, payload)
      }
    ).addMatcher(
      postsApiSlice.endpoints.updatePost.matchFulfilled,
      (state, { payload }) => {
        postsAdapter.upsertOne(state, payload)
      }
    ).addMatcher(
      postsApiSlice.endpoints.deletePost.matchFulfilled,
      (state) => {
        // todo: remove from posts adapter on deletion
        postsAdapter.removeOne(state, 'todo_get_post_id')
      }
    )
  },
})

export const { } = postsSlice.actions
export default postsSlice.reducer