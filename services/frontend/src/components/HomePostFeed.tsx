import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetFeedPostsQuery } from '../app/api/posts'

function HomePostFeed() {
  const { data, isLoading } = useGetFeedPostsQuery()
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
        return <FeedPost key={post.id} post={post} />
      })}
    </>
  )
}

export default HomePostFeed