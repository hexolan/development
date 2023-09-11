import { apiSlice } from '../api'

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
  })
})

export const { } = usersApiSlice