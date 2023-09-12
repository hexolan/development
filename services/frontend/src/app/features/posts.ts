import { createSlice } from '@reduxjs/toolkit'

import type { Post } from '../types'

export interface PostsState {

}

const initialState: PostsState = {

}

export const postsSlice = createSlice({
  name: 'posts',
  initialState,
  reducers: {

  },
})

export const { } = postsSlice.actions
export default postsSlice.reducer