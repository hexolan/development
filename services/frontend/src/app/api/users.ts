import { apiSlice } from '../api'

import type { RegisterUserRequest } from '../types/user'
import type { RawLoginResponse, LoginResponseData } from '../types/auth'

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

    registerUser: builder.mutation<LoginResponseData, RegisterUserRequest>({
      query: req => ({
        url: '/v1/users',
        method: 'POST',
        body: { ...req }
      }),
      transformResponse: (response: RawLoginResponse) => {
        return response.data as LoginResponseData
      }
    }),
  })
})

export const { useGetUserByNameQuery, useGetCurrentUserQuery, useDeleteCurrentUserMutation, useRegisterUserMutation } = usersApiSlice