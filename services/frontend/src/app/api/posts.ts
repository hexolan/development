import { apiSlice } from '../api'
import { convertRawPost } from '../types/posts'

import type { Post } from '../types/common'
import type {
  RawPost, RawPostResponse, RawPostsResponse,
  GetPanelPostRequest, GetPanelPostsRequest,
  GetUserPostsRequest,
  UpdatePostRequest,
  DeletePostRequest,
  CreatePostRequest
} from '../types/posts'

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPost: builder.query<Post, GetPanelPostRequest>({
      query: req => ({ url: `/v1/panels/id/${req.panelId}/posts/${req.postId}` }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return convertRawPost(response.data)
      }
    }),

    getUserPosts: builder.query<Post[], GetUserPostsRequest>({
      query: req => `/v1/users/id/${req.userId}/posts`,
      transformResponse: (response: RawPostsResponse) => {
        if (response.data === undefined) {
          throw Error('invalid posts response')
        } else if (!response.data.posts) {
          return []
        }

        return response.data.posts.map<Post>((rawPost: RawPost) => convertRawPost(rawPost))
      }
    }),

    getPanelPosts: builder.query<Post[], GetPanelPostsRequest>({
      query: req => `/v1/panels/id/${req.panelId}/posts`,
      transformResponse: (response: RawPostsResponse) => {
        if (response.data === undefined) {
          throw Error('invalid posts response')
        } else if (!response.data.posts) {
          return []
        }

        return response.data.posts.map<Post>((rawPost: RawPost) => convertRawPost(rawPost))
      }
    }),

    updatePost: builder.mutation<Post, UpdatePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'PATCH',
        body: { ...req.data },
      }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return convertRawPost(response.data)
      }
    }),

    deletePost: builder.mutation<void, DeletePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'DELETE',
      })
    }),

    createPanelPost: builder.mutation<Post, CreatePostRequest>({
      query: req => ({
        url: `/v1/panels/name/${req.panelName}`,
        method: 'POST',
        body: { ...req.data },
      }),
      transformResponse: (response: RawPostResponse) => {
        if (response.data === undefined) { throw Error('invalid post response') }

        return convertRawPost(response.data)
      }
    }),
  })
})

export const { 
  useGetPanelPostQuery, useGetPanelPostsQuery,
  useUpdatePostMutation, 
  useDeletePostMutation, 
  useCreatePanelPostMutation
} = postsApiSlice