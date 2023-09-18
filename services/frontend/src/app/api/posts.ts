import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import postsAdapter from '../features/posts'
import { convertRawTimestamp } from '../types/api'
import type { Post } from '../types/common'
import type {
  RawPost,
  GetPanelPostRequest, RawGetPanelPostResponse,
  GetPanelPostsRequest, RawGetPanelPostsResponse,
  CreatePanelPostRequest, RawCreatePanelPostResponse,
  UpdatePostRequest, RawUpdatePostResponse,
  DeletePostRequest, RawDeletePostResponse
} from '../types/posts'

const convertRawPost = (rawPost: RawPost): Post => ({
  id: rawPost.id,
  panelId: rawPost.panel_id,
  authorId: rawPost.author_id,
  title: rawPost.title,
  content: rawPost.content,
  createdAt: convertRawTimestamp(rawPost.created_at),
  updatedAt: (rawPost.updated_at ? convertRawTimestamp(rawPost.updated_at) : undefined),
})

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPost: builder.query<EntityState<Post>, GetPanelPostRequest>({
      query: req => ({ url: `/v1/panels/${req.panelName}/posts/${req.postId}` }),
      transformResponse: (response: RawGetPanelPostResponse) => {
        if (response.data === undefined) { 
          return postsAdapter.getInitialState()
        }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),

    getPanelPosts: builder.query<EntityState<Post>, GetPanelPostsRequest>({
      query: req => `/v1/panels/${req.panelName}/posts`,
      transformResponse: (response: RawGetPanelPostsResponse) => {
        if (response.data === undefined) {
          return postsAdapter.getInitialState()
        }

        const posts = response.data.posts.map<Post>((rawPost: RawPost) => convertRawPost(rawPost))
        return postsAdapter.setAll(postsAdapter.getInitialState(), posts)
      }
    }),

    createPanelPost: builder.mutation<EntityState<Post>, CreatePanelPostRequest>({
      query: req => ({
        url: `/v1/panels/${req.panelName}`,
        method: 'POST',
        body: { ...req.data },
      }),
      transformResponse: (response: RawCreatePanelPostResponse) => {
        if (response.data === undefined) {
          return postsAdapter.getInitialState()
        }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),

    updatePost: builder.mutation<EntityState<Post>, UpdatePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'PATCH',
        body: { ...req.data },
      }),
      transformResponse: (response: RawUpdatePostResponse) => {
        if (response.data === undefined) {
          return postsAdapter.getInitialState()
        }

        return postsAdapter.setOne(postsAdapter.getInitialState(), convertRawPost(response.data))
      }
    }),

    deletePost: builder.mutation<string, DeletePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: RawDeletePostResponse, _meta, arg) => {
        // todo: invalidate post
        console.log(arg)
        return response.msg
      }
    }),
  })
})

export const { useGetPanelPostsQuery, useGetPanelPostQuery, useCreatePanelPostMutation, useUpdatePostMutation, useDeletePostMutation } = postsApiSlice