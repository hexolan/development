import { useParams } from 'react-router-dom'
import { Container, Text } from '@mantine/core'

import FeedPost from '../components/FeedPost'
import LoadingBar from '../components/LoadingBar'
import { useGetPanelPostsQuery } from '../app/api/posts'

type PanelPageParams = {
  panelName: string;
}

function PanelPage() {
  const { panelName } = useParams<PanelPageParams>();
  if (panelName === undefined) {
    throw Error('panel name not provided')
  }

  // Load the panel posts.
  const { data } = useGetPanelPostsQuery({ panelName: panelName })
  if (data === undefined) {
    return <LoadingBar />
  }

  return (
    <Container mt='xl'>
      <Text>Panel - {panelName}</Text>
      <div>
        <Text>Posts:</Text>
        {
          Object.values(data.entities).map(post => {
            if (post === undefined) { return post }
            return <FeedPost key={post.id} post={post} />
          })
        }
      </div>
    </Container>
  )
}

export default PanelPage