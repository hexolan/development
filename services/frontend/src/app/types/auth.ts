import { convertRawUser } from './user'

import type { RawUser } from './user'
import type { RawResponse } from './api'
import type { AuthData, AuthToken } from './common'

// API Request Paramaters
export type LoginRequest = {
  username: string;
  password: string;
}

// API Responses
type RawLoginData = {
  token: AuthToken,
  user: RawUser
}

export type RawLoginResponse = RawResponse & {
  data?: RawLoginData;
}

// API Response Conversion
export const convertRawLoginData = (data: RawLoginData): AuthData => ({
  token: data.token,
  user: convertRawUser(data.user)
})