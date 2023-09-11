import { createSlice } from '@reduxjs/toolkit'

export interface CommentsState {

}

const initialState: CommentsState = {

}

export const commentsSlice = createSlice({
  name: 'comments',
  initialState,
  reducers: {

  },
})

export const { } = commentsSlice.actions
export default commentsSlice.reducer