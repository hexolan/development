import { createSlice } from '@reduxjs/toolkit'

// todo: expand upon
type Comment = {
  id: string;
  content: string;
}

type LoadedCommentsMap = {
  [postId: string]: Array<Comment>;
}

export interface CommentsState {
  items: LoadedCommentsMap
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