import { createEntityAdapter } from '@reduxjs/toolkit'

import type { User } from '../types/common'

const usersAdapter = createEntityAdapter<User>({
  selectId: (user) => user.id
})

export default usersAdapter