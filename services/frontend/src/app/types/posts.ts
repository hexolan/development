import { RawTimestamp } from './api'

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
  data: {
    title?: string;
    content?: string;
  };
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