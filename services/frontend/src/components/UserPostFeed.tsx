import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetUserPostsQuery } from '../app/api/posts'
import type { User } from '../app/types/common'

function UserPostFeed({ user }: { user: User }) {
  const { data, isLoading } = useGetUserPostsQuery({ userId: user.id })
  if (isLoading) {
    return <Text align='center'>Loading Posts...</Text>
  } else if (!data) {
    return <Text align='center' color='red'>Failed to Load Posts</Text>
  }

  // Check that there are posts.
  if (!data.length) {
    return <Text align='center'>No Posts Found!</Text>
  }

  return (
    <>
      {Object.values(data).map(post => {
        return <FeedPost key={post.id} post={post} authorInfo={user} />
      })}
    </>
  )
}

export default UserPostFeed