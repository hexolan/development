import { Stack, Center, Loader, Text } from '@mantine/core'

import Comment from './Comment'
import { useGetPostCommentsQuery } from '../app/api/comments'
import type { Post } from '../app/types/common'

function PostComments({ post }: { post: Post }) {
  const { data, isLoading } = useGetPostCommentsQuery({ postId: post.id })
  if (isLoading) {
    return (
      <Center>
        <Loader color='dark' size='sm' />
      </Center>
    )
  } else if (!data) {
    return <Text color='red' align='center'>Failed to Load Comments</Text>
  }

  if (!data.length) {
    return null
  }

  return (
    <Stack mb='lg' spacing='sm'>
      {Object.values(data).map(comment => {
        return <Comment key={comment.id} comment={comment} />
      })}
    </Stack>
  )
}

export default PostComments