import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { postsAdapter } from '../features/posts'
import { convertRawPost } from '../types/posts'

import type { Post } from '../types/common'
import type {
  RawPost, RawPostResponse, RawPostsResponse,
  GetPanelPostRequest, GetPanelPostsRequest,
  UpdatePostRequest,
  DeletePostRequest,
  CreatePostRequest
} from '../types/posts'

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPost: builder.query<EntityState<Post>, GetPanelPostRequest>({
      query: req => ({ url: `/v1/panels/${req.panelName}/posts/${req.postId}` }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),

    getPanelPosts: builder.query<EntityState<Post>, GetPanelPostsRequest>({
      query: req => `/v1/panels/${req.panelName}/posts`,
      transformResponse: (response: RawPostsResponse) => {
        if (response.data === undefined) { throw Error('invalid posts response') }

        const posts = response.data.posts.map<Post>((rawPost: RawPost) => convertRawPost(rawPost))
        return postsAdapter.setAll(postsAdapter.getInitialState(), posts)
      }
    }),

    updatePost: builder.mutation<EntityState<Post>, UpdatePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'PATCH',
        body: { ...req.data },
      }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),

    deletePost: builder.mutation<void, DeletePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'DELETE',
      })
    }),

    createPanelPost: builder.mutation<EntityState<Post>, CreatePostRequest>({
      query: req => ({
        url: `/v1/panels/${req.panelName}`,
        method: 'POST',
        body: { ...req.data },
      }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),
  })
})

export const { useGetPanelPostQuery, useGetPanelPostsQuery, useUpdatePostMutation, useDeletePostMutation, useCreatePanelPostMutation } = postsApiSlice