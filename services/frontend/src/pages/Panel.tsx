import { useParams } from 'react-router-dom'
import { Paper, Container, Stack, Text, rem } from '@mantine/core'

import FeedPost from '../components/FeedPost'
import LoadingBar from '../components/LoadingBar'
import { useGetPanelPostsQuery } from '../app/api/posts'

type PanelPageParams = {
  panelName: string;
}

// todo: get panel
// then call a get panel posts by panel id query

// todo: make posts nice by changing shadow effect on hover

function PanelPage() {
  const { panelName } = useParams<PanelPageParams>();
  if (panelName === undefined) {
    throw Error('panel name not provided')
  }

  // Load the panel posts.
  const { data, isLoading } = useGetPanelPostsQuery({ panelName: panelName })
  if (isLoading) {
    return <LoadingBar />
  } else if (!data) {
    // reminder to self: sometimes data error detail does not exist (e.g. connection refused to api / api offline)
    return <p>TODO: error page</p>
  }

  return (
    <>
      <Paper px="xl" py={rem(50)} shadow='md' sx={{ borderBottom: '1px' }}>
        <Text size='lg'>{panelName}</Text>
        <Text size='sm' color='dimmed'>This is the panel description</Text>

        {/* todo: new panel button here (on right side, opposite side, of two elems above */}
      </Paper>
      <Container mt='xl'>
        <Stack spacing='md'>
          {Object.values(data).map(post => {
            if (post === undefined) { return post }
            return <FeedPost key={post.id} post={post} />
          })}
        </Stack>
      </Container>
    </>
  )
}

export default PanelPage