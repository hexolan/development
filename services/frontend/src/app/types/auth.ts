import type { AuthToken, User } from './common'

export type LoginRequest = {
  username: string;
  password: string;
}

export type RawLoginResponse = {
  status: string;
  msg?: string;
  data?: LoginResponseData;
}

export type LoginResponseData = {
  token: AuthToken
  user: User
}