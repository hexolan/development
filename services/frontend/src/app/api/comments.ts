import { apiSlice } from '../api'

import type {
  GetPostCommentsRequest,
  CreatePostCommentRequest,
  UpdatePostCommentRequest,
  DeletePostCommentRequest
} from '../types/comments'
import type { Comment } from '../types/common'

// todo: transforming all comment responses

export const commentsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPostComments: builder.query<Comment[], GetPostCommentsRequest>({
      query: data => ({
        url: `/v1/posts/${data.postId}/comments`
      })
    }),

    createPostComment: builder.mutation<Comment, CreatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments`,
        method: 'POST',
        body: { ...req.data }
      })
    }),

    updatePostComment: builder.mutation<Comment, UpdatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.commentId}`,
        method: 'PATCH',
        body: { ...req.data }
      })
    }),

    deletePostComment: builder.mutation<void, DeletePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.commentId}`,
        method: 'DELETE'
      })
    })
  })
})

export const { useGetPostCommentsQuery, useCreatePostCommentMutation, useUpdatePostCommentMutation, useDeletePostCommentMutation } = commentsApiSlice