import { apiSlice } from '../api'

export interface SignInData {
  username: string;
  password: string;
}

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    signIn: builder.mutation({
      query: (data: SignInData) => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...data }
      })
    }),
  })
})

export const { useSignInMutation } = authApiSlice