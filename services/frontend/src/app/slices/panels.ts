import { createSlice } from '@reduxjs/toolkit'

export interface PanelsState {

}

const initialState: PanelsState = {

}

export const panelsSlice = createSlice({
  name: 'panels',
  initialState,
  reducers: {

  },
})

export const { } = panelsSlice.actions
export default panelsSlice.reducer