import { apiSlice } from '../api'
import type { LoginResponseData } from './auth';

type RegisterUserRequest = {
  username: string;
  password: string;
}

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
      })
    }),
  })
})

export const { useGetUserByNameQuery, useGetCurrentUserQuery, useDeleteCurrentUserMutation, useRegisterUserMutation } = usersApiSlice