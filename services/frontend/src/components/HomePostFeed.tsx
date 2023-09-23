import { Center, Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetFeedPostsQuery } from '../app/api/posts'

function HomePostFeed() {
  const { data, isLoading } = useGetFeedPostsQuery()
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
        return <FeedPost key={post.id} post={post} />
      })}
    </>
  )
}

export default HomePostFeed