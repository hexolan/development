import { apiSlice } from '../api'

export interface RegisterUserData {
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
      query: {
        url: '/v1/users/me'
      }
    }),

    deleteCurrentUser: builder.mutation({
      query: {
        url: '/v1/users/me',
        method: 'DELETE'
      }
    }),

    registerUser: builder.mutation({
      query: (data: RegisterUserData) => ({
        url: '/v1/users',
        method: 'POST',
        body: { ...data }
      })
    }),
  })
})

export const { useGetUserByNameQuery, useGetCurrentUserQuery, useDeleteCurrentUserMutation, useRegisterUserMutation } = usersApiSlice