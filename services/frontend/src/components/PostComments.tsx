import { Center, Text } from '@mantine/core'

import Comment from './Comment'
import { useGetPostCommentsQuery } from '../app/api/comments'
import type { Post } from '../app/types/common'

function PostComments({ post }: { post: Post }) {
  const { data, isLoading } = useGetPostCommentsQuery({ postId: post.id })
  if (isLoading) {
    return (
      <Center>
        <Text>Loading Comments...</Text>
      </Center>
    )
  } else if (!data) {
    return (
      <Center>
        <Text color='red'>Failed to Load Comments</Text>
      </Center>
    )
  }

  if (!data.length) {
    return (
      <Center>
        <Text>No Comments</Text>
      </Center>
    )
  }

  return (
    <>
      {Object.values(data).map(comment => {
        return <Comment key={comment.id} comment={comment} />
      })}
    </>
  )
}

export default PostComments