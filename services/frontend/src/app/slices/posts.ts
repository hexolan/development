import { createSlice } from '@reduxjs/toolkit'

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