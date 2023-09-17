import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Comment } from '../types'

const commentsAdapter = createEntityAdapter<Comment>({
  selectId: (comment) => comment.id
})

export default commentsAdapter