import { useParams } from 'react-router-dom'
import { Paper, Container, Text, rem } from '@mantine/core'

import FeedPost from '../components/FeedPost'
import LoadingBar from '../components/LoadingBar'
import { useGetPanelPostsQuery } from '../app/api/posts'

type PanelPageParams = {
  panelName: string;
}

// todo: get panel
// then call a get panel posts by panel id query

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
    <>
      <Paper px="xl" py={rem(50)} shadow='md' sx={{ borderBottom: '1px' }}>
        <Text size='lg'>{panelName}</Text>
        <Text size='sm' color='dimmed'>This is the panel description</Text>
      </Paper>
      <Container mt='xl'>
        {Object.values(data).map(post => {
          if (post === undefined) { return post }
          return <FeedPost key={post.id} post={post} />
        })}
      </Container>
    </>
  )
}

export default PanelPage