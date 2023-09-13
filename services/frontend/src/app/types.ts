type AuthToken = {
  token_type: string;
  access_token: string;
  expires_in: number;
}

type Panel = {
  id: string;
  name: string;
  description: string;
  createdAt: string;
  updatedAt?: string;
}

type Post = {
  id: string;
  panelId: string;
  authorId: string;
  title: string;
  content: string;
  createdAt: string;
  updatedAt?: string;
}

type Comment = {
  id: string;
  postId: string;
  authorId: string;
  message: string;
  createdAt: string;
  updatedAt?: string;
}

type User = {
  id: string;
  username: string;
  createdAt?: string;
  updatedAt?: string;
}

export type { AuthToken, Panel, Post, Comment, User }