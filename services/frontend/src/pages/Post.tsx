import { Text } from '@mantine/core'
import { useParams } from 'react-router-dom'

type PostPageParams = {
  panelName: string;
  postId: string;
}

function PostPage() {
  const { panelName, postId } = useParams<PostPageParams>();

  return (
    <Text>Post - {panelName} - {postId}</Text>
  )
}

export default PostPage