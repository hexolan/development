import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import { useGetPanelPostsQuery } from '../app/api/posts'
import type { Panel } from '../app/types/common'

function PanelPostFeed({ panel }: { panel: Panel }) {
  const { data, isLoading } = useGetPanelPostsQuery({ panelId: panel.id })
  if (isLoading) {
    return <Text align='center'>Loading Posts...</Text>
  } else if (!data) {
    return <Text align='center'>Failed to Load Posts</Text>
  }

  // Check that there are posts.
  if (!data.length) {
    return <Text align='center'>No Posts Found!</Text>
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