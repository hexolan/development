export type AuthToken = {
  token_type: string;
  access_token: string;
  expires_in: number;
}

export type Panel = {
  id: string;
  name: string;
  description: string;
  createdAt: string;
  updatedAt?: string;
}

export type Post = {
  id: string;
  panelId: string;
  authorId: string;
  title: string;
  content: string;
  createdAt: string;
  updatedAt?: string;
}

export type Comment = {
  id: string;
  postId: string;
  authorId: string;
  message: string;
  createdAt: string;
  updatedAt?: string;
}

export type User = {
  id: string;
  username: string;
  createdAt?: string;
  updatedAt?: string;
}