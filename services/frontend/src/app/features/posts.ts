import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Post } from '../types/common'

export const postsAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id
})