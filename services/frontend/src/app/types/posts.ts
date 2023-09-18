import { convertRawTimestamp } from './api'

import type { Post } from './common'
import type { RawResponse, RawTimestamp } from './api'

export const convertRawPost = (rawPost: RawPost): Post => ({
  id: rawPost.id,
  panelId: rawPost.panel_id,
  authorId: rawPost.author_id,
  title: rawPost.title,
  content: rawPost.content,
  createdAt: convertRawTimestamp(rawPost.created_at),
  updatedAt: (rawPost.updated_at ? convertRawTimestamp(rawPost.updated_at) : undefined),
})

export type RawPost = {
  id: string;
  panel_id: string;
  author_id: string;
  title: string;
  content: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type GetPanelPostRequest = {
  panelName: string;
  postId: string;
}

export type RawGetPanelPostResponse = RawResponse & {
  data?: RawPost;
}

export type GetPanelPostsRequest = {
  panelName: string;
}

export type RawGetPanelPostsResponse = RawResponse & {
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

export type RawCreatePanelPostResponse = RawResponse & {
  data?: RawPost;
}

export type UpdatePostRequest = {
  postId: string;
  data: {
    title?: string;
    content?: string;
  };
}

export type RawUpdatePostResponse = RawResponse & {
  data?: RawPost;
}

export type DeletePostRequest = {
  postId: string;
}

export type RawDeletePostResponse = {
  status: string;
  msg: string;
}