import type { RawResponse } from './api';
import type { AuthToken, User } from './common'

export type LoginRequest = {
  username: string;
  password: string;
}

export type RawLoginResponse = RawResponse & {
  data?: LoginResponseData;
}

export type LoginResponseData = {
  token: AuthToken
  user: User
}