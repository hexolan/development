import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { commentsAdapter } from '../features/comments'
import { convertRawComment } from '../types/comments'

import type { Comment } from '../types/common'
import type {
  RawComment,
  GetPostCommentsRequest, RawGetPostCommentsResponse,
  CreatePostCommentRequest, RawCreatePostCommentResponse,
  UpdatePostCommentRequest, RawUpdatePostCommentResponse,
  DeletePostCommentRequest, RawDeletePostCommentResponse
} from '../types/comments'

export const commentsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPostComments: builder.query<EntityState<Comment>, GetPostCommentsRequest>({
      query: data => ({ url: `/v1/posts/${data.postId}/comments` }),
      transformResponse: (response: RawGetPostCommentsResponse) => {
        if (response.data === undefined) { throw Error('invalid comments response') }

        const comments = response.data.comments.map<Comment>((rawComment: RawComment) => convertRawComment(rawComment))
        return commentsAdapter.setAll(commentsAdapter.getInitialState(), comments)
      }
    }),

    createPostComment: builder.mutation<EntityState<Comment>, CreatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments`,
        method: 'POST',
        body: { ...req.data }
      }),
      transformResponse: (response: RawCreatePostCommentResponse) => {
        if (response.data === undefined) { throw Error('invalid comment response') }

        return commentsAdapter.setOne(commentsAdapter.getInitialState(), convertRawComment(response.data))
      }
    }),

    updatePostComment: builder.mutation<EntityState<Comment>, UpdatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.commentId}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawUpdatePostCommentResponse) => {
        if (response.data === undefined) { throw Error('invalid comment response') }

        return commentsAdapter.setOne(commentsAdapter.getInitialState(), convertRawComment(response.data))
      }
    }),

    deletePostComment: builder.mutation<string, DeletePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.commentId}`,
        method: 'DELETE'
      }),
      transformResponse: (response: RawDeletePostCommentResponse, _meta, arg: DeletePostCommentRequest) => {
        // todo: invalidate comment
        console.log(arg)
        return response.msg
      }
    })
  })
})

export const { useGetPostCommentsQuery, useCreatePostCommentMutation, useUpdatePostCommentMutation, useDeletePostCommentMutation } = commentsApiSlice