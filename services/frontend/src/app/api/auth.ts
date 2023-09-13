import { apiSlice } from '../api'

import type { AuthToken, User } from '../types'

export type LoginRequest = {
  username: string;
  password: string;
}

export type RawLoginResponse = {
  status: string;
  msg?: string;
  data?: LoginResponseData;
}

export type LoginResponseData = {
  token: AuthToken
  user: User
}

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation<LoginResponseData, LoginRequest>({
      query: data => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...data }
      }),
      transformResponse: (response: RawLoginResponse) => {
        return response.data as LoginResponseData
      },
    }),
  })
})

export const { useLoginMutation } = authApiSlice