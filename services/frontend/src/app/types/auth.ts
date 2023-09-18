import { convertRawUser } from './user'

import type { RawUser } from './user'
import type { RawResponse } from './api'
import type { AuthToken, User } from './common'

export const convertRawLoginData = (data: RawLoginData): LoginData => ({
  token: data.token,
  user: convertRawUser(data.user)
})

export type LoginRequest = {
  username: string;
  password: string;
}

type RawLoginData = {
  token: AuthToken,
  user: RawUser
}

export type RawLoginResponse = RawResponse & {
  data?: RawLoginData;
}

export type LoginData = {
  token: AuthToken
  user: User
}