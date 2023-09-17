import type { RawResponse } from './api';

export type GetPostCommentsRequest = {
  postId: string;
}

export type CreateCommentData = {
  content: string;
}

export type CreatePostCommentRequest = {
  postId: string;
  data: CreateCommentData;
}

export type UpdateCommentData = {
  content?: string;
}

export type UpdatePostCommentRequest = {
  postId: string;
  commentId: string;
  data: UpdateCommentData;
}

export type DeletePostCommentRequest = {
  postId: string;
  commentId: string;
}