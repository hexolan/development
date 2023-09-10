import { apiSlice } from '../api'

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: builder => ({
    signIn: builder.mutation({
      query: details => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...details }
      })
    }),
  })
})

export const { useSignInMutation } = authApiSlice