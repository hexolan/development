import type { RawTimestamp } from './api';

export type RawUser = {
  id: string;
  username: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RegisterUserRequest = {
  username: string;
  password: string;
}