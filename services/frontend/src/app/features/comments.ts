import { createSlice } from '@reduxjs/toolkit'

import type { Comment } from '../types'

type CommentsMap = {
  [postId: string]: Array<Comment>;
}

export interface CommentsState {
  items: CommentsMap
}

const initialState: CommentsState = {
  items: {},
}

export const commentsSlice = createSlice({
  name: 'comments',
  initialState,
  reducers: {

  },
})

export const { } = commentsSlice.actions
export default commentsSlice.reducer