import { apiSlice } from '../api'

import type { AuthToken, User } from '../types';

export interface SignInData {
  username: string;
  password: string;
}

export interface RawAuthResponse {
  data: AuthResponseData;
  detail: string;
}

export interface AuthResponseData {
  token: AuthToken
  user: User
}

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    signIn: builder.mutation<AuthResponseData, SignInData>({
      query: (data: SignInData) => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...data }
      }),
      transformResponse: (response: RawAuthResponse) => {
        return response.data
      }
    }),
  })
})

export const { useSignInMutation } = authApiSlice