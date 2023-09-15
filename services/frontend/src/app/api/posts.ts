import { createEntityAdapter } from '@reduxjs/toolkit'
import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'

import type { Post } from '../types'

interface GetPanelPostRequest {
  panelName: string;
  postId: string;
}

export type RawGetPanelPostResponse = {
  status: string;
  msg?: string;
  data?: Post;
}

export type GetPanelPostsRequest = {
  panelName: string;
}

export type RawGetPanelPostsResponse = {
  status: string;
  msg?: string;
  data?: Post[];
}

export type CreatePanelPostRequest = {
  panelName: string;
  data: CreatePostData;
}

export type CreatePostData = {
  title: string;
  content: string;
}

export type RawCreatePanelPostResponse = {
  status: string;
  msg?: string;
  data?: Post;
}

export type RawCreatePostResponse = {
  status: string;
  msg?: string;
  data?: Post;
}

export type UpdatePostRequest = {
  postId: string;
  data: UpdatePostData;
}

export type UpdatePostData = {
  title?: string;
  content?: string;
}

export type RawUpdatePostResponse = {
  status: string;
  msg?: string;
  data?: Post;
}

export type DeletePostRequest = {
  postId: string;
}

export type RawDeletePostResponse = {
  status: string;
  msg: string;
}

const postsAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id
})

export const postsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelPost: builder.query<EntityState<Post>, GetPanelPostRequest>({
      query: req => ({ url: `/v1/panels/${req.panelName}/posts/${req.postId}` }),
      transformResponse: (response: RawGetPanelPostResponse) => {
        // todo: map response post onto post object (convert times to ISO strings)
        return postsAdapter.setOne(postsAdapter.getInitialState(), response.data as Post)
      }
    }),

    getPanelPosts: builder.query<EntityState<Post>, GetPanelPostsRequest>({
      query: req => `/v1/panels/${req.panelName}/posts`,
      transformResponse: (response: RawGetPanelPostsResponse) => {
        // todo: map all response posts onto post objects
        return postsAdapter.setAll(postsAdapter.getInitialState(), response.data as Post[])
      }
    }),

    createPanelPost: builder.mutation<EntityState<Post>, CreatePanelPostRequest>({
      query: req => ({
        url: `/v1/panels/${req.panelName}`,
        method: 'POST',
        body: { ...req.data },
      }),
      transformResponse: (response: RawCreatePanelPostResponse) => {
        // todo: map response post onto post object (convert times to ISO strings)
        return postsAdapter.setOne(postsAdapter.getInitialState(), response.data as Post)
      }
    }),

    updatePost: builder.mutation<EntityState<Post>, UpdatePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'PATCH',
        body: { ...req.data },
      }),
      transformResponse: (response: RawUpdatePostResponse) => {
        // todo: map response post onto post object (convert times to ISO strings)
        return postsAdapter.setOne(postsAdapter.getInitialState(), response.data as Post)
      }
    }),

    deletePost: builder.mutation<string, DeletePostRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}`,
        method: 'DELETE',
      }),
      transformResponse: (response: RawDeletePostResponse, _meta, arg) => { // todo: try with { arg }: { arg: string }
        // todo: invalidate post
        console.log(arg)
        return response.msg
      }
    }),
  })
})

export const { useGetPanelPostsQuery, useGetPanelPostQuery, useCreatePanelPostMutation, useUpdatePostMutation, useDeletePostMutation } = postsApiSlice