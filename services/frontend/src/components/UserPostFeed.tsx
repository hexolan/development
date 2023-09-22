import { Center, Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetUserPostsQuery } from '../app/api/posts'
import { User } from '../app/types/common'

function UserPostFeed({ user }: { user: User }) {
  if (!user) { return null }

  const { data, isLoading } = useGetUserPostsQuery({ userId: user.id })
  if (isLoading) {
    return (
      <Center>
        <Text>Loading Posts...</Text>
      </Center>
    )
  } else if (!data) {
    return (
      <Center>
        <Text color='red'>Failed to Load Posts</Text>
      </Center>
    )
  }

  // Check that there are posts.
  if (!data.length) {
    return (
      <Center>
        <Text>No Posts Found!</Text>
      </Center>
    )
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