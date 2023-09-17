import { Container, Text } from '@mantine/core'
import { useParams } from 'react-router-dom'

type PostPageParams = {
  panelName: string;
  postId: string;
}

function PostPage() {
  const { panelName, postId } = useParams<PostPageParams>();

  return (
    <Container mt='xl'>
      <Text>Post - {panelName} - {postId}</Text>
    </Container>
  )
}

export default PostPage