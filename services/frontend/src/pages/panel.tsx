import { Text } from '@mantine/core'
import { useParams } from 'react-router-dom'

import { useGetPanelPostsQuery } from '../app/api/posts'
import type { GetPanelPostsRequest } from '../app/api/posts'
import FeedPost from '../components/FeedPost'

type PanelPageParams = {
  panelName: string;
}

function PanelPage() {
  const { panelName } = useParams<PanelPageParams>();

  // Load the panel posts.
  const { data } = useGetPanelPostsQuery({ panelName: panelName } as GetPanelPostsRequest)
  if (data === undefined) {
    return <Text>Loading...</Text>
  }

  return (
    <>
      <Text>Panel - {panelName}</Text>
      <div>
        <Text>Posts:</Text>
        {
          Object.values(data.entities).map((post, _index) => {
            if (post === undefined) { return post }
            return <FeedPost key={post.id} post={post} />
          })
        }
      </div>
    </>
  )
}

export default PanelPage