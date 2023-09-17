import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'

import postsAdapter from '../features/posts';
import type { Post } from '../types'

type RawPost = {
  id: string;
  panel_id: string;
  author_id: string;
  title: string;
  content: string;
  created_at: {
    seconds: number;
    nanos?: number;
  };
  updated_at?: {
    seconds: number;
    nanos?: number;
  };
}

interface GetPanelPostRequest {
  panelName: string;
  postId: string;
}

export type RawGetPanelPostResponse = {
  status: string;
  msg?: string;
  data?: RawPost;
}

export type GetPanelPostsRequest = {
  panelName: string;
}

export type RawGetPanelPostsResponse = {
  status: string;
  msg?: string;
  data?: {
    posts: RawPost[];
  };
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
  data?: RawPost;
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
  data?: RawPost;
}

export type DeletePostRequest = {
  postId: string;
}

export type RawDeletePostResponse = {
  status: string;
  msg: string;
}

const convertRawPost = (rawPost: RawPost): Post => {
  let updatedAt = undefined
  if (rawPost.updated_at) {
    updatedAt = new Date(rawPost.updated_at.seconds * 1000).toISOString()
  }

  return {
    id: rawPost.id,
    panelId: rawPost.panel_id,
    authorId: rawPost.author_id,
    title: rawPost.title,
    content: rawPost.content,
    createdAt: new Date(rawPost.created_at.seconds * 1000).toISOString(),
    updatedAt: updatedAt,
  }
}

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
      transformResponse: (response: RawDeletePostResponse, _meta, arg) => { // todo: try with { arg }: { arg: string }
        // todo: invalidate post
        console.log(arg)
        return response.msg
      }
    }),
  })
})

export const { useGetPanelPostsQuery, useGetPanelPostQuery, useCreatePanelPostMutation, useUpdatePostMutation, useDeletePostMutation } = postsApiSlice