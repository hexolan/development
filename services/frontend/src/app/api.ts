import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

import type { RootState } from './store'

const baseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_API_URL,
  // credentials: 'include',
  prepareHeaders: (headers, { getState }) => {
    const state = getState() as RootState

    const token = state.auth.token
    if (token !== null) {
      headers.set('Authorization', `Bearer ${token}`)
    }

    return headers
  }
})

export const apiSlice = createApi({
  baseQuery: baseQuery,
  endpoints: _builder => ({}),
})