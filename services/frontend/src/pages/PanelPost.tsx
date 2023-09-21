import { useParams } from 'react-router-dom'
import { Container, Text } from '@mantine/core'

type PanelPostPageParams = {
  panelName: string;
  postId: string;
}

function PanelPostPage() {
  const { postId } = useParams<PanelPostPageParams>();

  return (
    <Container mt='xl'>
      <Text>Post - - {postId}</Text>
    </Container>
  )
}

export default PanelPostPage