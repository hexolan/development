import { Stack } from '@mantine/core'

import CommentComponent from './Comment'
import type { Comment } from '../app/types/common'

function CommentsFeed({ comments }: { comments: Comment[] }) {
  return (
    <Stack spacing='sm'>
      {Object.values(comments).map(comment => {
        return <CommentComponent key={comment.id} comment={comment} />
      })}
    </Stack>
  )
}

export default CommentsFeed