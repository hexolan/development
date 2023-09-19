import { apiSlice } from '../api'
import { convertRawLoginData } from '../types/auth'

import type { LoginRequest, RawLoginResponse, LoginData } from '../types/auth'

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation<LoginData, LoginRequest>({
      query: data => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...data }
      }),
      transformResponse: (response: RawLoginResponse) => {
        if (response.data === undefined) { throw Error('invalid login response') }

        return convertRawLoginData(response.data)
      },
    }),
  })
})

export const { useLoginMutation } = authApiSlice