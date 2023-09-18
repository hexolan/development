import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { usersAdapter } from '../features/users'
import { convertRawUser } from '../types/user'
import { convertRawLoginData } from '../types/auth'

import type { User } from '../types/common'
import type { RawLoginResponse, LoginData } from '../types/auth'
import type { GetUserByNameRequest, RawUserResponse, RegisterUserRequest } from '../types/user'

export const usersApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUserByName: builder.query<EntityState<User>, GetUserByNameRequest>({
      query: req => ({ url: `/v1/users/${req.username}` }),
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

    deleteCurrentUser: builder.mutation<void, void>({
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