import { createEntityAdapter } from '@reduxjs/toolkit'

import type { User } from '../types/common'

export const usersAdapter = createEntityAdapter<User>({
  selectId: (user) => user.id
})