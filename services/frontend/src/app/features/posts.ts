import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Post } from '../types'

const postsAdapter = createEntityAdapter<Post>({
  selectId: (post) => post.id
})

export default postsAdapter