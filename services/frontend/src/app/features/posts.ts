import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Post } from '../types/common'

const postsAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id
})

export default postsAdapter