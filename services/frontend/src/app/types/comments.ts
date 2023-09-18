import { convertRawTimestamp } from './api';

import type { Comment } from './common';
import type { RawResponse, RawTimestamp } from './api';

export const convertRawComment = (rawComment: RawComment): Comment => ({
  id: rawComment.id,
  postId: rawComment.post_id,
  authorId: rawComment.author_id,
  message: rawComment.message,
  createdAt: convertRawTimestamp(rawComment.created_at),
  updatedAt: (rawComment.updated_at ? convertRawTimestamp(rawComment.updated_at) : undefined),
})

export type RawComment = {
  id: string;
  post_id: string;
  author_id: string;
  message: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type GetPostCommentsRequest = {
  postId: string;
}

export type RawGetPostCommentsResponse = RawResponse & {
  data?: {
    comments: RawComment[];  // todo: verify this is the key
    // todo: verify keys of RawComment
  };
}

export type CreateCommentData = {
  message: string;
}

export type CreatePostCommentRequest = {
  postId: string;
  data: CreateCommentData;
}

export type RawCreatePostCommentResponse = RawResponse & {
  data?: RawComment;
}

export type UpdateCommentData = {
  message?: string;
}

export type UpdateCommentRequest = {
  commentId: string;
  data: UpdateCommentData;
}

export type UpdatePostCommentRequest = UpdateCommentRequest & {
  postId: string;
}

export type RawUpdatePostCommentResponse = RawResponse & {
  data?: RawComment;
}

export type DeleteCommentRequest = {
  commentId: string;
}

export type DeletePostCommentRequest = DeleteCommentRequest & {
  postId: string;
}

export type RawDeletePostCommentResponse = {
  status: string;
  msg: string;
}