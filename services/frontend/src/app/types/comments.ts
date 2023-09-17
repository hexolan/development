import type { RawResponse, RawTimestamp } from './api';

export type RawComment = {
  id: string;
  post_id: string;
  author_id: string;
  content: string;
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
  content: string;
}

export type CreatePostCommentRequest = {
  postId: string;
  data: CreateCommentData;
}

export type RawCreatePostCommentResponse = RawResponse & {
  data?: RawComment;
}

export type UpdateCommentData = {
  content?: string;
}

export type UpdatePostCommentRequest = {
  postId: string;
  commentId: string;
  data: UpdateCommentData;
}

export type RawUpdatePostCommentResponse = RawResponse & {
  data?: RawComment;
}

export type DeletePostCommentRequest = {
  postId: string;
  commentId: string;
}