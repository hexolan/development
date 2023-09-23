import { Center, Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetPanelPostsQuery } from '../app/api/posts'
import type { Panel } from '../app/types/common'

function PanelPostFeed({ panel }: { panel: Panel }) {
  const { data, isLoading } = useGetPanelPostsQuery({ panelId: panel.id })
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
        return <FeedPost key={post.id} post={post} panelInfo={panel} />
      })}
    </>
  )
}

export default PanelPostFeed