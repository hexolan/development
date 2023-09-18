import { apiSlice } from '../api'
import { convertRawLoginData } from '../types/auth'

import type { RegisterUserRequest } from '../types/user'
import type { RawLoginResponse, LoginData } from '../types/auth'

export const usersApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUserByName: builder.query({
      query: (username: string) => ({
        url: `/v1/users/${username}`
      })
    }),

    getCurrentUser: builder.query({
      query: () => ({
        url: '/v1/users/me'
      })
    }),

    deleteCurrentUser: builder.mutation({
      query: () => ({
        url: '/v1/users/me',
        method: 'DELETE'
      })
    }),

    registerUser: builder.mutation<LoginData, RegisterUserRequest>({
      query: req => ({
        url: '/v1/users',
        method: 'POST',
        body: { ...req }
      }),
      transformResponse: (response: RawLoginResponse) => {
        if (response.data === undefined) { throw Error('invalid registration response') }

        return convertRawLoginData(response.data)
      }
    }),
  })
})

export const { useGetUserByNameQuery, useGetCurrentUserQuery, useDeleteCurrentUserMutation, useRegisterUserMutation } = usersApiSlice