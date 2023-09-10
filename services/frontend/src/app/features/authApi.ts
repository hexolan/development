import { apiSlice } from '../api'

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: builder => ({
    signin: builder.mutation({
      query: details => ({
        url: '/v1/auth/',
        method: 'POST',
        body: { ...details }
      })
    }),
  })
})