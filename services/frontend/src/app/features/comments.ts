import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Comment } from '../types/common'

export const commentsAdapter = createEntityAdapter<Comment>({
  selectId: (comment) => comment.id
})