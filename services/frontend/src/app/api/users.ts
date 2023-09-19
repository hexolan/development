import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { usersAdapter } from '../features/users'
import { convertRawUser } from '../types/user'
import { convertRawAuthData } from '../types/auth'

import type { User, AuthData } from '../types/common'
import type { RawAuthResponse } from '../types/auth'
import type {
  RawUserResponse,
  GetUserByIdRequest, GetUserByNameRequest,
  DeleteUserByIdRequest, DeleteUserByNameRequest,
  RegisterUserRequest
} from '../types/user'

export const usersApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUserById: builder.query<EntityState<User>, GetUserByIdRequest>({
      query: req => ({ url: `/v1/users/id/${req.id}` }),
      transformResponse: (response: RawUserResponse) => {
        if (response.data === undefined) { throw Error('invalid user response') }

        return usersAdapter.setOne(usersAdapter.getInitialState(), convertRawUser(response.data))
      }
    }),

    getUserByName: builder.query<EntityState<User>, GetUserByNameRequest>({
      query: req => ({ url: `/v1/users/username/${req.username}` }),
      transformResponse: (response: RawUserResponse) => {
        if (response.data === undefined) { throw Error('invalid user response') }

        return usersAdapter.setOne(usersAdapter.getInitialState(), convertRawUser(response.data))
      }
    }),

    getCurrentUser: builder.query<EntityState<User>, void>({
      query: () => ({ url: '/v1/users/me' }),
      transformResponse: (response: RawUserResponse) => {
        if (response.data === undefined) { throw Error('invalid user response') }

        return usersAdapter.setOne(usersAdapter.getInitialState(), convertRawUser(response.data))
      }
    }),

    deleteUserById: builder.mutation<void, DeleteUserByIdRequest>({
      query: req => ({
        url: `/v1/users/id/${req.id}'`,
        method: 'DELETE'
      })
    }),

    deleteUserByName: builder.mutation<void, DeleteUserByNameRequest>({
      query: req => ({
        url: `/v1/users/username/${req.username}`,
        method: 'DELETE'
      })
    }),

    deleteCurrentUser: builder.mutation<void, void>({
      query: () => ({
        url: '/v1/users/me',
        method: 'DELETE'
      })
    }),

    registerUser: builder.mutation<AuthData, RegisterUserRequest>({
      query: req => ({
        url: '/v1/users',
        method: 'POST',
        body: { ...req }
      }),
      transformResponse: (response: RawAuthResponse) => {
        if (response.data === undefined) { throw Error('invalid registration response') }

        return convertRawAuthData(response.data)
      }
    }),
  })
})

// todo: redo these:
// export const { useGetUserByIdQuery, useGetUserByNameQuery, useGetCurrentUserQuery, useDeleteCurrentUserMutation, useRegisterUserMutation } = usersApiSlice