import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import commentsAdapter from '../features/comments'
import type {
  RawComment,
  GetPostCommentsRequest, RawGetPostCommentsResponse,
  CreatePostCommentRequest, RawCreatePostCommentResponse,
  UpdatePostCommentRequest, RawUpdatePostCommentResponse,
  DeletePostCommentRequest
} from '../types/comments'
import { convertRawTimestamp } from '../types/api'
import type { Comment } from '../types/common'

// todo: transforming all comment responses

const convertRawComment = (rawComment: RawComment): Comment => ({
  id: rawComment.id,
  postId: rawComment.post_id,
  authorId: rawComment.author_id,
  message: rawComment.message,
  createdAt: convertRawTimestamp(rawComment.created_at),
  updatedAt: (rawComment.updated_at ? convertRawTimestamp(rawComment.updated_at) : undefined),
})

export const commentsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPostComments: builder.query<EntityState<Comment>, GetPostCommentsRequest>({
      query: data => ({ url: `/v1/posts/${data.postId}/comments` }),
      transformResponse: (response: RawGetPostCommentsResponse) => {
        if (response.data === undefined) {
          return commentsAdapter.getInitialState()
        }

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
        if (response.data === undefined) {
          return commentsAdapter.getInitialState()
        }

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
        if (response.data === undefined) {
          return commentsAdapter.getInitialState()
        }

        return commentsAdapter.setOne(commentsAdapter.getInitialState(), convertRawComment(response.data))
      }
    }),

    deletePostComment: builder.mutation<void, DeletePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.commentId}`,
        method: 'DELETE'
      })
    })
    // todo: transforming response of deleting comment
  })
})

export const { useGetPostCommentsQuery, useCreatePostCommentMutation, useUpdatePostCommentMutation, useDeletePostCommentMutation } = commentsApiSlice