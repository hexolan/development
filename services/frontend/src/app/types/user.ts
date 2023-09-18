import { convertRawTimestamp } from './api';

import type { User } from './common';
import type { RawResponse, RawTimestamp } from './api';

export const convertRawUser = (rawUser: RawUser): User => ({
  id: rawUser.id,
  username: rawUser.username,
  createdAt: convertRawTimestamp(rawUser.created_at),
  updatedAt: (rawUser.updated_at ? convertRawTimestamp(rawUser.updated_at) : undefined),
})

export type RawUser = {
  id: string;
  username: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RawUserResponse = RawResponse & {
  data?: RawUser;
}

export type GetUserByNameRequest = {
  username: string;
}

export type RegisterUserRequest = {
  username: string;
  password: string;
}