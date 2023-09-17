import { useParams } from 'react-router-dom'
import { Container, Text } from '@mantine/core'

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