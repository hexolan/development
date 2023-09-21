import { Center, Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetPanelPostsQuery } from '../app/api/posts'

function PanelPostFeed({ panelId }: { panelId: string }) {
  const { data, isLoading } = useGetPanelPostsQuery({ panelId: panelId })
  
  // todo: improve
  if (isLoading) {
    return (
      <Center>
        <Text>Loading Posts...</Text>
      </Center>
    )
  } else if (!data) {
    // todo: error message
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
        return <FeedPost key={post.id} post={post} hidePanel={true} />
      })}
    </>
  )
}

export default PanelPostFeed