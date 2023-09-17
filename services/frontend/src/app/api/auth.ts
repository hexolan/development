import { apiSlice } from '../api'

import type { LoginRequest, RawLoginResponse, LoginResponseData } from '../types/auth'

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