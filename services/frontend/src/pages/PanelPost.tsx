import { useParams, useOutletContext } from 'react-router-dom'
import { Container, Text } from '@mantine/core'

import LoadingBar from '../components/LoadingBar';
import { useGetPanelPostQuery } from '../app/api/posts';
import type { PanelContext } from '../components/PanelLayout';

type PanelPostPageParams = {
  panelName: string;
  postId: string;
}

function PanelPostPage() {
  const { postId } = useParams<PanelPostPageParams>();
  if (postId === undefined) { throw Error('post id not provided') }

  const { panel } = useOutletContext<PanelContext>()


  const { data, isLoading } = useGetPanelPostQuery({ panelId: panel.id, postId: postId })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    throw Error('Post not found!')  // todo: extract exact error msg (as it may not just be a 404)
  }

  return (
    <Container mt='xl'>
      <Text>Post - {data.title}</Text>
    </Container>
  )
}

export default PanelPostPage