import { apiSlice } from '../api'

import type { CreateCommentData, UpdateCommentData } from '../types/comments'

export const commentsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPostComments: builder.query({
      query: (postId: string) => ({
        url: `/v1/posts/${postId}/comments`
      })
    }),

    createPostComment: builder.mutation({
      query: ({postId, data}: {postId: string, data: CreateCommentData}) => ({
        url: `/v1/posts/${postId}/comments`,
        method: 'POST',
        body: { ...data }
      })
    }),

    updatePostComment: builder.mutation({
      query: ({postId, commentId, data}: {postId: string, commentId: string, data: UpdateCommentData}) => ({
        url: `/v1/posts/${postId}/comments/${commentId}`,
        method: 'PATCH',
        body: { ...data }
      })
    }),

    deletePostComment: builder.mutation({
      query: ({ postId, commentId }: { postId: string, commentId: string }) => ({
        url: `/v1/posts/${postId}/comments/${commentId}`,
        method: 'DELETE'
      })
    })
  })
})

export const { useGetPostCommentsQuery, useCreatePostCommentMutation, useUpdatePostCommentMutation, useDeletePostCommentMutation } = commentsApiSlice